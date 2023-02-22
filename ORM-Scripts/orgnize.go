package ormscripts

func Orgnize() {
	db := db_init("douyin-youth-training")
	defer db.Close()

	//进行数据库的迁移（也就是新建五个表）
	db.AutoMigrate(&User_DB{})
	db.AutoMigrate(&Video_DB{})
	db.AutoMigrate(&Comment_DB{})
	db.AutoMigrate(&Likes_DB{})
	db.AutoMigrate(&Follower_DB{})
}
