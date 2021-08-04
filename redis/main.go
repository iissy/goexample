package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// zpop pops a value from the ZSET key using WATCH/MULTI/EXEC commands.
func zpop(c redis.Conn, key string) (result string, err error) {

	defer func() {
		// Return connection to normal state on error.
		if err != nil {
			c.Do("DISCARD")
		}
	}()

	// Loop until transaction is successful.
	for {
		if _, err := c.Do("WATCH", key); err != nil {
			return "", err
		}

		members, err := redis.Strings(c.Do("ZRANGE", key, 0, 0))
		if err != nil {
			return "", err
		}
		if len(members) != 1 {
			return "", redis.ErrNil
		}

		c.Send("MULTI")
		c.Send("ZREM", key, members[0])
		queued, err := c.Do("EXEC")
		if err != nil {
			return "", err
		}

		if queued != nil {
			result = members[0]
			break
		}
	}

	return result, nil
}

// zpopScript pops a value from a ZSET.
var zpopScript = redis.NewScript(1, `
    local r = redis.call('ZRANGE', KEYS[1], 0, 0)
    if r ~= nil then
        r = r[1]
        redis.call('ZREM', KEYS[1], r)
    end
    return r
`)

// This example implements ZPOP as described at
// http://redis.io/topics/transactions using WATCH/MULTI/EXEC and scripting.
func main() {
	c, err := redis.Dial("tcp", "192.168.1.234:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	// Add test data using a pipeline.

	for i, member := range []string{"red", "blue", "green"} {
		c.Send("ZADD", "zset", i, member)
	}
	if _, err := c.Do(""); err != nil {
		fmt.Println(err)
		return
	}

	// Pop using WATCH/MULTI/EXEC

	v, err := zpop(c, "zset")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)

	// Pop using a script.

	v, err = redis.String(zpopScript.Do(c, "zset"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)

}
