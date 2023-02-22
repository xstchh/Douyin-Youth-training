package ormscripts

//User Json
type User_JSON struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  string `json:"total_favorited"`
	WorkCount       int    `json:"work_count"`
	FavoritedCount  int    `json:"favorited_count"`
	FollowerList    []User_JSON
	FriendList      []User_JSON
}

//Video Json
type Video_JSON struct {
	ID            int64     `json:"id,omitempty"`
	Author        User_JSON `json:"author"`
	PlayUrl       string    `json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount int64     `json:"favorite_count,omitempty"`
	CommentCount  int64     `json:"comment_count,omitempty"`
	IsFavorite    bool      `json:"is_favorite,omitempty"`
	Title         string    `json:"title,omitempty"`
}

//Comment Json
type Comment_JSON struct {
	ID         int64     `json:"id,omitempty"`
	User       User_JSON `json:"user"`
	Content    string    `json:"content,omitempty"`
	CreateDate string    `json:"create_date,omitempty"`
}
