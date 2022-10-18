package mapper

import (
	"emptyProject/global"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

var SqlSession *gorm.DB

// 连接数据库
func ConnectDB() (err error) {
	DriverName := global.MysqlDbSetting.DriverName
	Host := global.MysqlDbSetting.Host
	Port := global.MysqlDbSetting.Port
	User := global.MysqlDbSetting.User
	Password := global.MysqlDbSetting.Password
	Dbname := global.MysqlDbSetting.Dbname
	psqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", User, Password, Host, Port, Dbname)
	SqlSession, err = gorm.Open(DriverName, psqlInfo)
	if err != nil {
		log.Println("connectDBError", err)
	}
	return SqlSession.DB().Ping()
}

func Close() {
	SqlSession.Close()

}
