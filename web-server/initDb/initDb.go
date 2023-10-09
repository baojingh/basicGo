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
