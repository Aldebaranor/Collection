package global

import (
	"emptyProject/global/viper"
	"log"
)

// 和yml对应
type ServerSettings struct {
	RunMode  string
	HttpPort string
}

type DataSourceSettings struct {
	Source string
}

type MysqlDbSettings struct {
	DriverName string
	Host       string
	Port       string
	User       string
	Password   string
	Dbname     string
}

type PostgresDbSettings struct {
	DriverName string
	Host       string
	Port       string
	User       string
	Password   string
	Dbname     string
	Sslmode    string
}

type MqttSettings struct {
	Address  string
	UserName string
	Password string
	Topic    string
	QoS      byte
}

// 定义全局变量
var (
	ServerSetting     *ServerSettings
	DataSourceSetting *DataSourceSettings
	MysqlDbSetting    *MysqlDbSettings
	PostgresDbSetting *PostgresDbSettings
	MqttSetting       *MqttSettings
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
	ds, err := viper.NewSetting()
	if err != nil {
		return err
	}
	err = ds.ReadSection("DataSource", &DataSourceSetting)
	if err != nil {
		return err
	}
	mydb, err := viper.NewSetting()
	if err != nil {
		return err
	}
	err = mydb.ReadSection("Mysql", &MysqlDbSetting)
	if err != nil {
		return err
	}
	pgdb, err := viper.NewSetting()
	if err != nil {
		return err
	}
	err = pgdb.ReadSection("Postgres", &PostgresDbSetting)
	if err != nil {
		return err
	}
	mq, err := viper.NewSetting()
	if err != nil {
		return err
	}
	err = mq.ReadSection("Mqtt", &MqttSetting)
	if err != nil {
		return err
	}
	log.Printf("DataSource:")
	log.Printf("%+v", DataSourceSetting)
	log.Printf("MysqlSetting:")
	log.Printf("%+v", MysqlDbSetting)
	log.Printf("PgsqlSetting:")
	log.Printf("%+v", PostgresDbSetting)
	log.Printf("MqttSetting:")
	log.Printf("%+v", MqttSetting)
	return nil
}
