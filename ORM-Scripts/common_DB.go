package ormscripts

//gorm默认名为 `ID` 的字段会作为primary key

//1. User DB
type User_DB struct {
	ID              int64
	Name            string
	FollowCount     int64
	FollowerCount   int64
	IsFollow        bool
	Avatar          string
	BackgroundImage string
	Signature       string
	TotalFavorited  string
	WorkCount       int
	FavoritedCount  int
}

func (*User_DB) TableName() string {
	return "user_table"
}

//2. Video DB
type Video_DB struct {
	ID            int64
	Author        User_DB
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	IsFavorite    bool
	Title         string
}

func (*Video_DB) TableName() string {
	return "video_table"
}

//3. Comment DB
type Comment_DB struct {
	ID         int64
	User       User_DB
	Content    string
	CreateDate string
}

func (*Comment_DB) TableName() string {
	return "comment_table"
}

//4. Fans-Follower DB
type Follower_DB struct {
	ID     int64
	UserId int64
	FansId int64
}

func (*Follower_DB) TableName() string {
	return "follower_table"
}

//5. Likes DB
type Likes_DB struct {
	ID      int64
	UserId  int64
	VideoId int64
}

func (*Likes_DB) TableName() string {
	return "likes_table"
}
