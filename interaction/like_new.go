package main

import (

	"Douyin-Youth-training/controller"
	"errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 喜欢列表对应结构体
type Favorite struct {
	Id         int64     `json:"id" gorm:"column:favorite;primaryKey"`
	UserId     int64     `josn:"user_id" gorm:"column:favorite;"`
	VideoId    int64     `josn:"video_id" gorm:"column:favorite;"`
	CreateTime time.Time `josn:"create_time" gorm:"column:favorite;"`
	DeleteTime gorm.DeletedAt
}

// 查询用户

func QueryUserByID(id int64) (*User, error) {
	var res *User
	dsn := "abc" //这里需要用户相关数据库的dsn
	db, err0 := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err0 != nil {
		panic("failed to connect database, error=" + err0.Error())
	}
	err := db.Where("Id = ?", id).First(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

// 判断视频是否存在
func QueryVideoByID(id int64) (*Video, error) {
	var video *Video
	dsn := "abc" //这里需要视频相关数据库的dsn
	db, err0 := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err0 != nil {
		panic("failed to connect database, error=" + err0.Error())
	}
	err := db.Where("Id = ?", id).First(&video).Error
	if err != nil {
		return nil, err
	}
	return video, nil
}
func RunFavoriteAction(video Video, actionType int, user User) error {
	//判断用户是否存在
	_, err := QueryUserByID(user.Id)
	if err != nil {
		return errors.New("the user doesn't exist")
	}
	//判断视频是否存在
	videos, err := QueryVideoByID(video.Id)
	if err != nil || videos == nil {
		return errors.New("the video doesn't exist")
	}
	//这里的dsn
	db, err := gorm.Open(mysql.Open("dsn"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, error=" + err.Error())
	}
	//判断操作状态
	switch actionType {
	case 1:
		if video.IsFavorite == true {
			return errors.New("you have liked this video before")
		} else {
			video.FavoriteCount++
			video.IsFavorite = true
			//这里的db
			like := &Favorite{UserId: user.Id, VideoId: video.Id, CreateTime: time.Now()}
			db.Create(like)
			return nil
		}
	case 2:
		if video.IsFavorite == false {
			return errors.New("you have not like this video before")
		} else {
			video.FavoriteCount--
			video.IsFavorite = false
			//注意这里的db
			db.Delete(&Favorite{UserId: user.Id, VideoId: video.Id})
			return nil
		}
	default:
		return errors.New("undefined action_type")
	}
	return nil
}

func ShowFavoriteList(user_id int64) {
	dsn := "table_name: favorite"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, error=" + err.Error())
	}
	videos := make([]*Video, 0)
	db.Where("UserId = ?", user_id).Find(&videos)
	//怎么呈现出来
}
