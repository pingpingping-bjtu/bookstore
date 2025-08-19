package global

import (
	"bookstore-manager/config"
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBClient *gorm.DB
var RedisClient *redis.Client

// InitMysql 数据库初始化，通过gorm
func InitMysql() {
	mysqlConfig := config.AppConfig.DataBase
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Name)
	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("连接数据库失败", err)
	}
	DBClient = client
	log.Println("连接数据库成功")
}

// InitRedis 初始化redis，使用redis库
func InitRedis() {
	redisConfig := config.AppConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	RedisClient = client
	str, err := client.Ping(context.TODO()).Result()
	if err != nil {
		log.Fatalln("redis连接失败", err)
	}
	log.Println("str:", str)

	log.Println("redis连接成功")

}

func GetDB() *gorm.DB {
	return DBClient
}
func CloseDB() {
	if DBClient != nil {
		sqlDB, err := DBClient.DB()
		if err != nil {
			err := sqlDB.Close()
			if err != nil {
				return
			}
		}

	}
}
