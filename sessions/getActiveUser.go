package sessions

import "net/http"

func GetLoggedInUser(r *http.Request) string {
	userID := GetSession(r)
	return userID
}
