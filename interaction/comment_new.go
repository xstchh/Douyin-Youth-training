package main

import (
	//"Douyin-Youth-training/controller"
	"errors"
	"time"

	//"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Video      Video  `json:"video"`
	Content    string `json:"content,omitempty" gorm:"column:comment;"`
	CreateDate time.Time
	DeleteDate gorm.DeletedAt
}

func QueryCommentById(id int64) (*Comment, error) {
	var comment *Comment
	db_comment, err := gorm.Open(mysql.Open("dsn"), //dsn
		&gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := db_comment.Where("Id = ?", id).Find(&comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

// 我参考的代码上有ctx context.Context这个参数，但我还不太会，就先不要了
func RunCommentAction(commentId int64, video Video, actionType int, user User, commentText string) error {
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
	switch actionType {
	case 1:
		//添加评论
		db.Create(&Comment{
			User:       user,
			Video:      video,
			Content:    commentText,
			CreateDate: time.Now()})
		video.CommentCount++
	case 2:
		//通过id查找评论，没有：评论不存在，有：删除评论,count-1
		res_comment, err := QueryCommentById(commentId)
		if res_comment == nil || err != nil {
			return errors.New("not find the comment")
		}
		video.CommentCount--
		db.Delete(&Comment{
			User:    user,
			Video:   video,
			Content: commentText})
	default:
		return errors.New("undefined action_type")
	}
	return nil
}
