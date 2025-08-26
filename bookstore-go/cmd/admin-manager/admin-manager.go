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
	"os/signal"
	"syscall"
	"time"
)

func main() {

	//初始化，MySQL redis、配置文件
	config.InitConfig("conf/config.yaml")
	cfg := config.AppConfig
	global.InitMysql()
	global.InitRedis()
	//设置优雅关闭

	//启动web服务
	r := router.InitRouter()
	addr := fmt.Sprintf("%s:%d", "localhost", cfg.Server.AdminPort)
	server := http.Server{
		Addr:    addr,
		Handler: r,
		//添加超时配置
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	//信号监听
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	//在goroutine中启动服务器
	go func() {
		log.Printf("Server running at http://localhost%s", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()
	sig := <-quit
	log.Printf("Received signal:%v", sig)
	switch sig {
	case syscall.SIGINT:
		log.Println("Received SIGINT (Ctrl+C),shutting down server...")
	case syscall.SIGTERM:
		log.Println("Received SIGTERM,shutting down server...")
	case syscall.SIGHUP:
		log.Println("Received SIGHUP,shutting down server...")
	case syscall.SIGQUIT:
		log.Println("Received SIGQUIT,shutting down server...")
	default:
		log.Printf("Received %s,shutting down server...", sig)
	}
	//创建5秒超时的context用于优雅关闭
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//关闭HTTP服务器
	log.Println("Shutting down HTTP server...")

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	} else {
		log.Println("HTTP server gracefully stopped")
	}
	//清理资源
	log.Println("Cleaning up resources...")
	cleanResources()
	// 等待一小段时间确保所有资源都被正确释放
	time.Sleep(100 * time.Millisecond)
	log.Println("Server exited successfully")

	// 强制退出程序
	os.Exit(0)
}
func cleanResources() {
	if global.RedisClient != nil {
		log.Println("Closing Redis connection...")
		global.CloseRedis()
	}
	if global.DBClient != nil {
		log.Println("Closing database connection...")
		global.CloseDB()
	}
	time.Sleep(100 * time.Millisecond)
	log.Println("All resources cleaned up")
}
