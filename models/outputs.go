package models

import (
	"time"

	u "github.com/cetRide/car-api/utils"
)

func DisplayPosts() {

	type PostDesc struct {
		Post_Id       string
		File          string
		Comment_Count int64
		Likes_Count   int64
		Views_Count   int64
		Reposts_Count int64
	}
	type TimelinePosts struct {
		User_Id   string
		Username  string
		Profile   string
		PostsDesc []PostDesc
	}
}

func DisplayNewsFeedsPosts(user string) map[string]interface{} {

	type PostDesc struct {
		Post_Id       string
		Filename      string
		Caption       string
		Created_At    time.Time
		Comment_Count int64
		Likes_Count   int64
		Views_Count   int64
		Reposts_Count int64
	}
	type TimelinePosts struct {
		User_Id   string
		Username  string
		Profile   string
		PostsDesc []PostDesc
	}
	if err := GetDB().Table("posts").Select("accounts.user_id, accounts.username, accounts.profile, posts.post_id, posts.filename, posts.caption, posts.created_at").
		Joins("JOIN accounts ON accounts.user_id = posts.user_id").
		Where("").
		Order("posts.created_at").Scan(&TimelinePosts{}).Error; err != nil {
		return u.Message(false, "You currently do not have posts")
	}
	response := u.Message(true, "Display posts")
	response["posts"] = &TimelinePosts{}
	return response
}
