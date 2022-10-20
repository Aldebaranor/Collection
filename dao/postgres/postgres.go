package postgres

import (
	"Collection/global"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

var PgSqlSession *gorm.DB

// 连接数据库
func ConnectPGDB() (err error) {
	DriverName := global.PostgresDbSetting.DriverName
	Host := global.PostgresDbSetting.Host
	Port := global.PostgresDbSetting.Port
	User := global.PostgresDbSetting.User
	Password := global.PostgresDbSetting.Password
	Dbname := global.PostgresDbSetting.Dbname
	Sslmode := global.PostgresDbSetting.Sslmode
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", Host, Port, User, Password, Dbname, Sslmode)
	PgSqlSession, err = gorm.Open(DriverName, psqlInfo)
	if err != nil {
		log.Println("connectDBError", err)
	}
	return PgSqlSession.DB().Ping()
}

func ClosePG() {
	PgSqlSession.Close()

}
