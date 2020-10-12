package sessions

import (
	"net/http"
	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func GetSession(request *http.Request) (userID string) {
	if cookie, err := request.Cookie("userID"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("userID", cookie.Value, &cookieValue); err == nil {
			userID = cookieValue["userID"]
		}
	}
	return userID
}

func DestroySession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "userID",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}
func CreateSession(userID string, w http.ResponseWriter, r *http.Request) {
	if userID != "" {
		value := map[string]string{
			"userID": userID,
		}
		if encoded, err := cookieHandler.Encode("userID", value); err == nil {
			cookie := &http.Cookie{
				Name:   "userID",
				Value:  encoded,
				Path:   "/",
				MaxAge: 3600,
			}
			http.SetCookie(w, cookie)
		}
	} else {
		DestroySession(w)
	}
}
