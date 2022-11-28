package database

import (
	"bluebell/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// 定义一个全局对象db
var db *sqlx.DB

func init() {
	cfg := config.Conf.MySQLConfig
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", cfg.User, cfg.Password, "tcp", cfg.Host, cfg.Port, cfg.DBName)
	db = sqlx.MustConnect("mysql", conn)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
}

// Close 因为全局变量db没有对外暴露，所以要封装一个对外暴露的方法Close供主函数使用,这是一个技巧
func Close() {
	err := db.Close()
	if err != nil {
		zap.L().Info("mysql close failed", zap.String("error", err.Error()))
	}
}

// GetDBClient 因为db对象不会对全局暴露，所以要通过方法对外暴露
func GetDBClient() *sqlx.DB{
	return db
}