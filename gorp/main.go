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

	// 删除所有数据，这样可以使得自增的数字从新从 1 开始
	err := dbmap.TruncateTables()
	checkErr(err, "TruncateTables failed")

	// 创建两个话题，无需给 Id 赋值，因为 Id 是自增列
	p1 := newPost("go-gorp 是一个 orm 框架", "我们使用 gorp 在我们的 web 项目中")
	p2 := newPost("gorp 与 mysql 的使用样例", "gorp 是一款流行的 go orm 框架")

	// 插入上面两条数据，Id 是自增，无需赋值
	err = dbmap.Insert(&p1, &p2)
	checkErr(err, "Insert failed")

	// SelectInt 获取整型单值
	count, err := dbmap.SelectInt("select count(*) from posts")
	checkErr(err, "select count(*) failed")
	log.Println("Rows after inserting:", count)

	// 更新值
	p2.Title = "go-gorp 更新值，按照 id 指定的更新了"
	count, err = dbmap.Update(&p2)
	checkErr(err, "Update failed")
	log.Println("Rows updated:", count)

	// 获取一行，注意 “post_id” 代替 “Id”，它是一个别名
	err = dbmap.SelectOne(&p2, "select * from posts where post_id=?", p2.Id)
	checkErr(err, "SelectOne failed")
	log.Println("p2 row:", p2)

	// 获取所有行，这里最体现 orm 的特点，它会自动对应结构体的字段，
	// 注意结构体的字段必须不能少于数据库返回的字段，否则会报错
	// 返回的每个字段均不能为空，否则会报错
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

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
