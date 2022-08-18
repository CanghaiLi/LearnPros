package conf

import (
	"fmt"
	"github.com/CanghaiLi/LearnPros/gin-pro/model"
	"gopkg.in/ini.v1"
	"log"
	"strings"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		log.Fatalln("配置文件读取错误！", err)
	}
	// 将配置文件读入变量中
	LoadIniConfig(file)

	// 链接数据库，设置空闲连接池中(保留)的最大连接数
	Dsn := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.DBConnect(Dsn)
}

func LoadIniConfig(file *ini.File) {
	// service config
	service := file.Section("service")
	AppMode = service.Key("AppMode").String()
	HttpPort = service.Key("HttpPort").String()
	fmt.Println("----- AppMode -----  ==> ", AppMode)

	// mysql config
	mysql := file.Section("mysql")
	Db = mysql.Key("Db").String()
	DbHost = mysql.Key("DbHost").String()
	DbPort = mysql.Key("DbPort").String()
	DbUser = mysql.Key("DbUser").String()
	DbPassword = mysql.Key("DbPassword").String()
	DbName = mysql.Key("DbName").String()
}
