package config

import (
	"bluebell/cmd"
	"fmt"
	"log"
)

var (
	SnowFlakeConfig = SnowFlake{StartTime: "2022-11-23", MachineID: 1}
)

// 启动配置项
type AppConfig struct {
	Port         int    `json:"port" mapstructure:"port"`
	Name         string `json:"name" mapstructure:"name"`
	Mode         string `json:"mode" mapstructure:"mode"`
	Version      string `json:"version" mapstructure:"version"`
	*LogConfig   `json:"log" mapstructure:"log"`
	*MySQLConfig `json:"mysql" mapstructure:"mysql"`
	*RedisConfig `json:"redis" mapstructure:"redis"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `json:"level" mapstructure:"level"`
	Filename   string `json:"filename" mapstructure:"filename"`
	MaxSize    int    `json:"max_size" mapstructure:"max_size"`
	MaxAge     int    `json:"max_age" mapstructure:"max_age"`
	MaxBackups int    `json:"max_backups" mapstructure:"max_backups"`
}

// DBConfig DB配置项
type MySQLConfig struct {
	Host         string `json:"host" mapstructure:"host"`
	User         string `json:"user" mapstructure:"user"`
	Password     string `json:"password" mapstructure:"password"`
	DBName       string `json:"dbname" mapstructure:"dbname"`
	Port         int    `json:"port" mapstructure:"port"`
	MaxOpenConns int    `json:"max_open_conns" mapstructure:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns" mapstructure:"max_idle_conns"`
}

// RedisConfig redis配置项
type RedisConfig struct {
	Host     string `json:"host" mapstructure:"host"`
	Password string `json:"password" mapstructure:"password"`
	Port     int    `json:"port" mapstructure:"port"`
	DB       int    `json:"db" mapstructure:"db"`
	PoolSize int    `json:"pool_size" mapstructure:"pool_size"`
}

type SnowFlake struct {
	StartTime string
	MachineID int64
}

// Conf 全局配置变量
var Conf = new(AppConfig)

func init() {
	filePath := cmd.GetConfigFileName()
	if filePath == ""{
		panic("not found the config file")
	}
	log.Println("config file path is:", filePath)
	if err := InitViper(filePath); err != nil {
		panic(fmt.Sprintf("init config failed, err:%v\n", err))
	}
}
