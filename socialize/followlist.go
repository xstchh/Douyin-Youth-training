package socialize

import (
	"log"
	"sync"
)

type Follow struct {
	Id         int64
	UserId     int64
	FollowerId int64
	Cancel     int8
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
func (*FollowUser,db *gorm) GetFollowingList(userId int64) ([]int64, error) {
	var ids []int64
	if err := db.Model(Follow{}).Where("follower.id = ?", userId).Pluck("user_id", &ids).Error; nil != err {
		if "record not found" == err.Error() {
			return nil, nil
		}
		log.Println(err.Error())
		return nil, err
	}
	return ids, nil
}
