package socialize

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"gorm.io/driver/sqlite"
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
func GetCommentList(ctx context.Context, c *app.RequestContext) {
	publishId := c.Param("user_id")
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Println("failed to link")
	}
	var commentList []Comment
	res := db.Model(Comment{}).Where(map[string]interface{}{"publish_id": publishId, "cancel": ValidComment}).Order("create_date desc").Find(&commentList)
	if res.RowsAffected == 0 {
		log.Println("return no comments")
	}
	if res.Error != nil {
		log.Println(res.Error.Error())
		log.Println("running failed")
	}
	c.JSON(400, utils.H{
		"UserList": commentList,
	})
	log.Println("success!")
}
