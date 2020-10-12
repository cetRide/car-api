package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/cetRide/car-api/models"
	"github.com/cetRide/car-api/sessions"
	u "github.com/cetRide/car-api/utils"
	"github.com/gorilla/mux"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user := &models.UserInput{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	response := models.LoginUser(user.Username, user.Password, w, r)
	u.Respond(w, response)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	user := &models.UserInput{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	response := models.RegisterUser(user.Username, user.Email, user.Password)
	u.Respond(w, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {

}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	follower := sessions.GetLoggedInUser(r)
	following := params["follow"]
	response := models.Follow(follower, following)
	u.Respond(w, response)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	follower := sessions.GetLoggedInUser(r)
	following := params["follow"]
	response := models.UnFollow(follower, following)
	u.Respond(w, response)
}
