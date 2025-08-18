package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var AppConfig Config

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	DataBase DataBaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
}
type ServerConfig struct {
	Port int `yaml:"port"`
}
type DataBaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}
type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func InitConfig(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln("读取配置文件失败", err)
	}
	if err := yaml.Unmarshal(data, &AppConfig); err != nil {
		log.Fatalln("yaml反序列化配置失败", err)
	}
	log.Println("加载配置文件成功")

}
