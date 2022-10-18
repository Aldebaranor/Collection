package global

import (
	"emptyProject/global/viper"
	"log"
)

// 和yml对应
type ServerSettingS struct {
	RunMode  string
	HttpPort string
}

type MysqlDbSettings struct {
	DriverName string
	Host       string
	Port       string
	User       string
	Password   string
	Dbname     string
}

// 定义全局变量
var (
	ServerSetting  *ServerSettingS
	MysqlDbSetting *MysqlDbSettings
)

// 读取配置到全局便量
func SetupSetting() error {
	s, err := viper.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &ServerSetting)
	if err != nil {
		return err
	}
	pgdb, err := viper.NewSetting()
	if err != nil {
		return err
	}
	err = pgdb.ReadSection("Mysql", &MysqlDbSetting)
	if err != nil {
		return err
	}
	log.Printf("pgSetting:")
	log.Printf("%+v", MysqlDbSetting)
	return nil
}
