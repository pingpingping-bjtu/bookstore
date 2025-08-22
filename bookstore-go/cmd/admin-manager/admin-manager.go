package main

import (
	"bookstore-manager/config"
	"bookstore-manager/global"
	"bookstore-manager/web/router"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	//初始化，MySQL redis、配置文件
	config.InitConfig("conf/config.yaml")
	global.InitMysql()
	global.InitRedis()
	r := router.InitAdminRouter()
	addr := fmt.Sprintf("%s:%d", "localhost", config.AppConfig.Server.AdminPort)
	server := http.Server{
		Addr:    addr,
		Handler: r,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("服务启动失败")
		os.Exit(-1)
	}
	err = server.Shutdown(context.TODO())
	if err != nil {
		log.Println("服务器错误退出:", err)
		return
	} else {
		log.Println("服务器正常退出")
	}
	cleanResources()
}
func cleanResources() {
	if global.RedisClient != nil {
		log.Println("redis退出，资源清理")
		global.CloseRedis()
	}
	if global.DBClient != nil {
		log.Println("mysql退出，资源清理")
		global.CloseDB()
	}
	time.Sleep(1 * time.Second)
	log.Println("所有资源清理完毕")

}
