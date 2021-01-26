package db
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func NewMysql() (db *gorm.DB){
	db, err := gorm.Open("mysql", "root:123456@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		log.Panicln("数据库连接失败,err"+err.Error())
	}
	return
}
