package interaction

import (
	"Douyin-Youth-training/base"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"time"
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

func QueryUserByID(id int64) (*base.User, error) {
	var res *base.User
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
func QueryVideoByID(id int64) (*base.Video, error) {
	var video *base.Video
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
func RunFavoriteAction(video base.Video, actionType int, user base.User) error {
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

func ShowFavoriteList(ctx context.Context, c *app.RequestContext) {
	userId := c.Query("user_id")
	db, err := gorm.Open(mysql.Open("dsn"), &gorm.Config{}) //dsn没搞
	if err != nil {
		panic("failed to connect database, error=" + err.Error())
	}
	var videos_list []base.Video
	db.Where("UserId = ?", userId).Find(&videos_list)
	//返回请求序列化为josn格式
	c.JSON(http.StatusOK, base.VideoListResponse{
		Response: base.Response{
			StatusCode: 0, StatusMsg: "success"},
		VideoList: videos_list,
	})
}
