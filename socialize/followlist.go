package socialize

import (
	"context"
	"log"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


type Follow struct {
	Id          int64
	user_id     int64
	follower_id int64
	Cancel      int8
}

func (Follow) TableName() string {
	return "follows"
}

type FollowUser struct {
}

var (
	followuser *FollowUser
	followOnce sync.Once
)

func UniFollow() *FollowUser {
	followOnce.Do(
		func() {
			followuser = &FollowUser{}
		})
	return followuser
}
func (*FollowUser) GetFollowList(ctx context.Context, c *app.RequestContext) {
	userId := c.Param("user_id")
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Println("failed to link")
	}
	var ids []int64
	if err := db.Model(Follow{}).Where("follower.id = ?", userId).Pluck("user_id", &ids).Error; nil != err {
		if "record not found" == err.Error() {
			log.Println(err.Error())
		}
		log.Println(err.Error())
	}
	c.JSON(200, utils.H{
		"UserList": ids,
	})
	log.Println("success!")
}
