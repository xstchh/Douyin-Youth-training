package base

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"sync/atomic"
)

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

// Register 函数用于实现用户注册功能，会接受用户名和密码作为参数，并检查是否已经存在该用户，
// 如果不存在则新建用户，并返回注册成功信息。
func Register(ctx *context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password
	// 打开数据库
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 自动创建表
	err = db.AutoMigrate(&User{})
	if err != nil {
		return
	}
	// 查找用户是否存在
	var user User
	result := db.First(&user, "Name = ?", username)
	// 查询出记录不是0条，说明用户存在
	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		// 创建用户
		db.Create(&User{
			Id:       userIdSequence,
			Name:     username,
			PassWord: password,
			Token:    token,
		})
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}

// Login 函数用于实现用户登录功能，会接受用户名和密码作为参数，并检查是否已经存在该用户，
// 如果存在则返回登录成功信息。
func Login(ctx *context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password
	// 打开数据库
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 查找用户是否存在
	var user User
	result := db.First(&user, "Name = ?", username)
	// 如果用户存在
	if result.RowsAffected != 0 {
		// 返回登录信息
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

// UserInfo 函数用于查询用户的信息，会接受token作为参数，并检查是否存在该用户，
// 如果存在则返回用户信息。
func UserInfo(ctx *context.Context, c *app.RequestContext) {
	token := c.Query("token")
	// 打开数据库
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 查找用户是否存在
	var user User
	result := db.First(&user, "Token = ?", token)
	// 如果用户存在
	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
