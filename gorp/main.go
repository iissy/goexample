package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v2"
	"log"
	"time"
)

func main() {
	// 初始化 DbMap
	dbmap := initDb()
	defer dbmap.Db.Close()

	// 删除所有数据
	err := dbmap.TruncateTables()
	checkErr(err, "TruncateTables failed")

	// create two posts
	p1 := newPost("明天去爬山有人一起吗？", "莲花山走起，明天天气会很好，加入吗？")
	p2 := newPost("官宣，我们结婚了！", "取得白富美为妻子，快来祝福我吧！")

	// 插入上面两条数据，Id 是自增，无需赋值
	err = dbmap.Insert(&p1, &p2)
	checkErr(err, "Insert failed")

	// SelectInt 获取整型单值
	count, err := dbmap.SelectInt("select count(*) from posts")
	checkErr(err, "select count(*) failed")
	log.Println("Rows after inserting:", count)

	// 更新值
	p2.Title = "Go 1.2 is better than ever"
	count, err = dbmap.Update(&p2)
	checkErr(err, "Update failed")
	log.Println("Rows updated:", count)

	// 获取一行，注意 “post_id” 代替 “Id”，它是一个别名
	err = dbmap.SelectOne(&p2, "select * from posts where post_id=?", p2.Id)
	checkErr(err, "SelectOne failed")
	log.Println("p2 row:", p2)

	// 获取所有行
	var posts []Post
	_, err = dbmap.Select(&posts, "select * from posts order by post_id")
	checkErr(err, "Select failed")
	log.Println("All rows:")
	for x, p := range posts {
		log.Printf("    %d: %v\n", x, p)
	}

	// 根据主键删除数据
	count, err = dbmap.Delete(&p1)
	checkErr(err, "Delete failed")
	log.Println("Rows deleted:", count)

	// 删除行通过 exec
	_, err = dbmap.Exec("delete from posts where post_id=?", p2.Id)
	checkErr(err, "Exec failed")

	// 确定表中已无数据
	count, err = dbmap.SelectInt("select count(*) from posts")
	checkErr(err, "select count(*) failed")
	log.Println("Row count - should be zero:", count)

	log.Println("Done!")
}

type Post struct {
	Id      int64 `db:"post_id"`
	Created int64
	Title   string `db:",size:50"`
	Body    string `db:"article_body,size:1024"`
}

func newPost(title, body string) Post {
	return Post{
		Created: time.Now().UnixNano(),
		Title:   title,
		Body:    body,
	}
}

func initDb() *gorp.DbMap {
	// 使用标准的 Go database/sql API 建立连接
	db, err := sql.Open("mysql", "root:pwd123456@tcp(192.168.111.151:3306)/hrefs?charset=utf8")
	checkErr(err, "sql.Open failed")

	// 创建一个结构
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

	// 添加一个表，设置表名“posts”，指定一个主键为“Id”，true 表示自增
	dbmap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")

	// 如果表不存在，创建表
	//err = dbmap.CreateTablesIfNotExists()
	//checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
