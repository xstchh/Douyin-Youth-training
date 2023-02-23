package socialize

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

const ValidComment = 0
const InvalidComment = 1

type Comment struct {
	Id           int64
	User_id      int64
	Publish_id   int64
	Comment_text string
	Create_date  time.Time
	Cancel       int32
}

func (Comment) TableName() string {
	return "comment"
}
func GetCommentList(videoId int64, db *gorm.DB) ([]Comment, error) {
	var commentList []Comment
	res := db.Model(Comment{}).Where(map[string]interface{}{"video_id": videoId, "cancel": ValidComment}).Order("create_date desc").Find(&commentList)
	if res.RowsAffected == 0 {
		log.Println("return no comments")
		return nil, nil
	}
	if res.Error != nil {
		log.Println(res.Error.Error())
		log.Println("running failed")
		return commentList, errors.New("Failed")
	}
	log.Println("success!")
	return commentList, nil

}
