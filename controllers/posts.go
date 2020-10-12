package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/cetRide/car-api/models"
	"github.com/cetRide/car-api/sessions"
	u "github.com/cetRide/car-api/utils"
	"github.com/gorilla/mux"
	"github.com/h2non/filetype"
)

//1MB
const MAX_MEMORY = 1 * 1024 * 1024

func PostFile(w http.ResponseWriter, r *http.Request) {
	/*Check the length of the video file and audio file*/
	/*Check the size of the file....set the maximum size to be uploaded.*/
	user := sessions.GetLoggedInUser(r)
	var filename string
	fileFormat := [...]string{
		"video/mp4",
		"video/x-m4v",
		"video/x-matroska",
		"video/webm",
		"video/quicktime",
		"video/x-msvideo",
		"video/x-ms-wmv",
		"video/mpeg",
		"video/x-flv",
		"video/3gpp",
		"video/wvm",
		"video/avi",

		"audio/midi",
		"audio/mpeg",
		"audio/m4a",
		"audio/ogg",
		"audio/x-flac",
		"audio/x-wav",
		"audio/amr",
		"audio/aac"}

	path := "userFiles" + string(os.PathSeparator) + user + string(os.PathSeparator) + "posts"
		if err := r.ParseMultipartForm(MAX_MEMORY); err != nil {
		u.Respond(w, u.Message(false, "File size is too large"))
		return
	}
	caption := r.FormValue("caption")
	file, _, err := r.FormFile("file")
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	kind, _ := filetype.Match(fileBytes)
	if kind == filetype.Unknown {
		u.Respond(w, u.Message(false, "Unknown file type"))
		return
	}
	formatOfFile := http.DetectContentType(fileBytes)
	fmt.Println("the file format is", formatOfFile)

	for i, _ := range fileFormat {
		if fileFormat[i] == formatOfFile {
			tempFile, err := ioutil.TempFile(path, user+"_dailypods_*."+kind.Extension)
			if err != nil {
				u.Respond(w, u.Message(false, "Internal server error"))
				return
			}
			defer tempFile.Close()
			tempFile.Write(fileBytes)
			filename = tempFile.Name()
			response := models.PostApost(user, filename, caption)
			u.Respond(w, response)
			return
		}
	}
	u.Respond(w, u.Message(false, "Invalid file format"))
	return

}

func Repost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := sessions.GetLoggedInUser(r)
	postId := params["postId"]
	response := models.Reposting(user, postId)
	u.Respond(w, response)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := sessions.GetLoggedInUser(r)
	postId := params["postId"]
	status := params["status"]
	response := models.LikingPost(status, user, postId)
	u.Respond(w, response)
}

func CommentPost(w http.ResponseWriter, r *http.Request) {
	user := &models.UserInput{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	params := mux.Vars(r)
	userId := sessions.GetLoggedInUser(r)
	postId := params["postId"]
	response := models.CommentingPost(user.Comment, userId, postId)
	u.Respond(w, response)
}
func ReplyToComments(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	comentId := params["commentId"]
	userID := sessions.GetLoggedInUser(r)
	text := &models.UserInput{}
	err := json.NewDecoder(r.Body).Decode(text)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := models.ReplyingToComments(comentId, userID, text.Comment)
	u.Respond(w, resp)
}

// func DeletePost(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	user := sessions.GetLoggedInUser(r)
// 	postId := params["postId"]
// 	response := models.DeletingPost(user, postId)
// 	u.Respond(w, response)
// }

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	commentId := params["commentId"]
	response := models.DeletingComment(commentId)
	u.Respond(w, response)
}
func DeleteReplies(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	replyId := params["replyId"]
	response := models.DeletingReplies(replyId)
	u.Respond(w, response)
}
