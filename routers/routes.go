package routers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cetRide/car-api/controllers"
	"github.com/cetRide/car-api/jwt_auth"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(jwt_auth.JwtAuthentication)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

type Routes []Route

var routes = Routes{

	Route{
		"Login",
		"POST",
		"/account/login",
		controllers.Login,
	},
	Route{
		"Signup",
		"POST",
		"/account/signup",
		controllers.Signup,
	},
	Route{
		"Logout",
		"POST",
		"/account/logout",
		controllers.Logout,
	},
	Route{
		"FollowUser",
		"POST",
		"/follow/{follow}",
		controllers.FollowUser,
	},
	Route{
		"UnFollowUser",
		"DELETE",
		"/unfollow/{follow}",
		controllers.UnfollowUser,
	},
	Route{
		"PostFile",
		"POST",
		"/post",
		controllers.PostFile,
	},
	Route{
		"LikePost",
		"POST",
		"/likepost/{postId}/status",
		controllers.LikePost,
	},
	Route{
		"CommentPost",
		"POST",
		"/commentpost/{postId}",
		controllers.CommentPost,
	},
	Route{
		"ReplyToComments",
		"POST",
		"/commentpost/{commentId}",
		controllers.ReplyToComments,
	},
	Route{
		"Repost",
		"POST",
		"/repost/{postId}",
		controllers.Repost,
	},
	// Route{
	// 	"DeletePost",
	// 	"DELETE",
	// 	"/deletepost/{postId}",
	// 	controllers.DeletePost,
	// },
	Route{
		"DeleteComment",
		"DELETE",
		"/deletecomment/{commentId}",
		controllers.DeleteComment,
	},
	Route{
		"DeleteReplies",
		"DELETE",
		"/deletecomment/{replyId}",
		controllers.DeleteReplies,
	},
}
