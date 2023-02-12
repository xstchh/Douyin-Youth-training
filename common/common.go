package common

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

type User struct {
	Id            int64 `json:"user_id"`
	Name          string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
	FollowerList  []User // TODO:新添加粉丝列表属性，是否最优做法待商榷
	FriendList    []User // TODO:同上
}

type Video struct {
	Id            int    `json:"id"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int    `json:"favorite_count"`
	CommentCount  int    `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite" default:"true"`
	Title         string `json:"title"`
}
