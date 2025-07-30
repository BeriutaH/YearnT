package config

import (
	"Yearn-go/service"
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type mysqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type Config struct {
	MySQL mysqlConfig `yaml:"mysql"`
}

var (
	DB  *gorm.DB
	Cfg Config
)

func InitConfig() {
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("读取yaml配置文件失败: %v", err)
	}

	if err := yaml.Unmarshal(data, &Cfg); err != nil {
		log.Fatalf("yaml解组错误: %v", err)
	}
}

func InitDB() {
	InitConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Cfg.MySQL.User,
		Cfg.MySQL.Password,
		Cfg.MySQL.Host,
		Cfg.MySQL.Port,
		Cfg.MySQL.DBName,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	// 自动建表
	if err := DB.AutoMigrate(service.AutoMigrateAll()...); err != nil {
		log.Fatal("数据库迁移失败:", err)
	}
}
