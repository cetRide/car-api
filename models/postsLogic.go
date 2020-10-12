package models

import (
	u "github.com/cetRide/car-api/utils"
)

func PostApost(user, file, caption string) map[string]interface{} {
	if err := GetDB().Create(
		&Posts{
			Post_Id:    u.GenerateUID(),
			User_Id:    user,
			Filename:   file,
			Caption:    caption,
			Created_At: u.GenerateTimeNow()}).Error; err != nil {
		return u.Message(false, "Unable to post the file.")
	}
	return u.Message(true, "Successfully posted.")
}

func DeleteApost(user, post_id string) map[string]interface{} {
	if err := GetDB().Table("posts").Where("user_id = ? AND post_id = ?", user, post_id).Delete(&Posts{}).Error; err != nil {
		return u.Message(false, "Unable to delete the post.")
	}
	return u.Message(true, "Post is successfully deleted.")
}

func Reposting(user, postId string) map[string]interface{} {
	if err := GetDB().Create(&Reposts{
		Repost_Id:  u.GenerateUID(),
		User_Id:    user,
		Post_Id:    postId,
		Created_At: u.GenerateTimeNow()}).Error; err != nil {
		return u.Message(false, "Reposting error")
	}
	return u.Message(true, "Reposting successfull")
}

func LikingPost(status, user, postId string) map[string]interface{} {
	if status == "1" {
		if err := GetDB().Create(&Likes{
			Like_Id:    u.GenerateUID(),
			User_Id:    user,
			Post_Id:    postId,
			Created_At: u.GenerateTimeNow()}).Error; err != nil {
			return u.Message(false, "Error in liking the post")
		}
		return u.Message(true, "Post liked")
	} else {
		if err := GetDB().Table("likes").Where("user_id = ? AND post_id = ?", user, postId).Delete(&Likes{}).Error; err != nil {
			return u.Message(false, "Error in handling likes")
		}
		return u.Message(true, "Unliked the post")
	}
}

func CommentingPost(comment, user, postId string) map[string]interface{} {
	if err := GetDB().Create(&Comments{Post_Id: postId, Comment_Id: u.GenerateUID(), User_Id: user, Created_At: u.GenerateTimeNow(), Comment: comment}).Error; err != nil {
		return u.Message(false, "Comment not be send.")
	}
	return u.Message(true, "Comment send.")
}

func DeletingComment(comment_id string) map[string]interface{} {
	if err := GetDB().Table("comments").Where("comment_id = ?", comment_id).Delete(&Comments{}).Error; err != nil {
		return u.Message(false, "Comment not be deleted.")
	}
	return u.Message(true, "Comment deleted.")
}
func ReplyingToComments(commentId, userID, reply string) map[string]interface{} {
	if err := GetDB().Create(&Replies{Comment_Id: commentId, Reply_Id: u.GenerateUID(), User_Id: userID, Reply: reply, Created_At: u.GenerateTimeNow()}).Error; err != nil {
		return u.Message(false, "Reply not send.")
	}
	return u.Message(true, "Reply send.")
}
func DeletingReplies(reply_id string) map[string]interface{} {
	if err := GetDB().Table("replies").Where("reply_id = ?", reply_id).Delete(&Replies{}).Error; err != nil {
		return u.Message(false, "Reply not be deleted.")
	}
	return u.Message(true, "Reply deleted.")
}
