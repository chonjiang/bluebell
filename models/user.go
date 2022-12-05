package models

import (
	"bluebell/database"
	"time"
)

var db = database.GetDBClient()

// User 针对数据库表 user 的结构体定义
type User struct {
	Id         int64     `json:"id" db:"id"`          // Id 类型: int64 主健字段（Primary Key） 自增长字段
	UserId     int64     `json:"user_id,string" db:"user_id"`     // UserId 类型: int64
	Username   string    `json:"username" db:"username"`    // Username 类型: string
	Password   string    `json:"password" db:"password"`    // Password 类型: string
	Email      string    `json:"email" db:"email"`       // Email 类型: string
	Gender     int8      `json:"gender" db:"gender"`      // Gender 类型: int8 默认值: 0
	CreateTime time.Time `json:"create_time" db:"create_time"` // CreateTime 类型: time.Time 默认值: CURRENT_TIMESTAMP
	UpdateTime time.Time `json:"update_time" db:"update_time"` // UpdateTime 类型: time.Time 默认值: CURRENT_TIMESTAMP
}

// Add 插入User
func (model *User) Add() (int64, error) {
	sql := "INSERT INTO `user` (`user_id`, `username`, `password`, `email`, `gender`) VALUES (?,?,?,?,?)"
	result, err := db.Exec(sql, model.UserId, model.Username, model.Password, model.Email, model.Gender)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// Update 更新
func (model *User) Update() (int64, error) {
	sqlStr := "UPDATE `user` SET `username` = ?, `password` = ?, `email` = ?, `gender` = ? WHERE `id` = ?"
	result, err := db.Exec(sqlStr, model.Username, model.Password, model.Email, model.Gender, model.Id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// Find 查询
func Find(condition string, args ...interface{}) ([]*User, error) {
	sqlStr := "SELECT `user_id`, `username`, `password`, `email`, `gender` FROM `user`"
	if len(condition) > 0 {
		sqlStr = sqlStr + " WHERE " + condition
	}

	results := make([]*User, 0)
	err := db.Select(&results, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	return results, nil
}
