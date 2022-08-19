package cache

import (
	"fmt"
	"github.com/CanghaiLi/LearnPros/gin-pro/pkg/utils"
	"github.com/go-redis/redis"
	"gopkg.in/ini.v1"
	"strconv"
)

// RedisClient Redis缓存客户端单例
var (
	RedisClient *redis.Client
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

// Redis 初始化redis链接
func Redis() {

	file, err := ini.Load("conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadRedis(file)

	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPw,
		DB:       int(db),
	})
	pong, err := client.Ping().Result()
	if err != nil {
		utils.LoggersObj.Info(err)
		panic(err)
	}
	RedisClient = client
	fmt.Println("redis", RedisClient, pong)
}

func LoadRedis(file *ini.File) {
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}
