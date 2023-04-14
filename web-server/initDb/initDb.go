package initDb

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	log.Println("start connect*************")
	Db, err := sql.Open("mysql", "root:Hadoop.1qaz!QAZ@tcp(121.5.73.196:3306)/devops")
	if err != nil {
		log.Panicln("err to connect:", err.Error())
	}
	Db.SetMaxOpenConns(20)
	Db.SetMaxIdleConns(20)
	log.Println("finish connect*************")
}

// func init() {
// 	var err error
// 	Db, err := sql.Open("mysql", "root:Hadoop.1qaz!QAZ@tcp(121.5.73.196:3306)/devops")
// 	if err != nil {
// 		log.Panicln("fail to connect database, ", err)
// 	}
// 	Db.SetMaxOpenConns(10)
// 	Db.SetMaxIdleConns(60)
// }
