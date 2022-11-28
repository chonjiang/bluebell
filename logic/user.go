package logic

import (
	"bluebell/config"
	"bluebell/models"
	"bluebell/tools/encrypt"
	"bluebell/tools/snowflake"
	"database/sql"
	"log"
)

// Register
func Register(user *models.User) (err error) {
	res, err := models.Find("username=?", user.Username)
	if err != nil {
		return
	}
	if len(res) > 0 {
		// 用户已存在
		return config.ErrorUserExist
	}
	// 生成user_id
	user.UserId = snowflake.GenIDInt64()
	log.Println(user.UserId)
	// 生成加密密码
	user.Password = encrypt.GetEncrypt([]byte(user.Password))
	// 把用户插入数据库
	_, err = user.Add()
	return err
}

// Login
func Login(loginUser *models.User) (user *models.User, err error) {
	originPassword := loginUser.Password // 记录一下原始密码
	res, err := models.Find("username=?", loginUser.Username)
	if err != nil {
		return nil, err
	}
	if err == sql.ErrNoRows {
		// 用户不存在
		return nil, config.ErrorUserNotExist
	}

	loginUser = res[0]

	// 生成加密密码与查询到的密码比较
	password := encrypt.GetEncrypt([]byte(originPassword))
	if loginUser.Password != password {
		return nil, config.ErrorPasswordWrong
	}
	return loginUser, nil
}
