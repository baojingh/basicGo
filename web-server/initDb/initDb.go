package initDb

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

func init() {
	var mydb_dsn = "root:Hadoop.1qaz!QAZ@tcp(121.5.73.196:3306)/devops"
	Db, _ = gorm.Open(mysql.Open(mydb_dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})
}

// import (
// 	"log"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var Db *gorm.DB

// func init() {
// 	var err error
// 	log.Println("start connect*************")
// 	Db, err = gorm.Open("mysql", "root:Hadoop.1qaz!QAZ@tcp(121.5.73.196:3306)/devops")
// 	if err != nil {
// 		log.Panicln("err to connect:", err.Error())
// 	}
// 	log.Println("finish connect*************")
// }

// func init() {
// 	var err error
// 	Db, err := sql.Open("mysql", "root:Hadoop.1qaz!QAZ@tcp(121.5.73.196:3306)/devops")
// 	if err != nil {
// 		log.Panicln("fail to connect database, ", err)
// 	}
// 	Db.SetMaxOpenConns(10)
// 	Db.SetMaxIdleConns(60)
// }
