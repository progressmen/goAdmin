package handle

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//const Dsn = "work:npWS1Iu5MCmYmG9U@tcp(104.160.43.85:3306)/shahejiuhuo?charset=utf8mb4&parseTime=True&loc=Local"
const Dsn = "work:npWS1Iu5MCmYmG9U@tcp(127.0.0.1:3306)/shahejiuhuo?charset=utf8mb4&parseTime=True&loc=Local"

const UplodPath = "https://t.me33.cn/uploads/"

var DbCon *gorm.DB

func InitDb() {
	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}
	DbCon = db
}
