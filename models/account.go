package models

import (
	"net/http"
	"os"

	"github.com/cetRide/car-api/sessions"
	u "github.com/cetRide/car-api/utils"
	validation "github.com/cetRide/car-api/validation"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(username string, password string, w http.ResponseWriter, r *http.Request) map[string]interface{} {
	account := &Accounts{}
	err := GetDB().Table("accounts").Where("username = ?", username).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Username not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Incorrect password. Please try again")
	}

	//Create JWT token
	tk := &Token{Username: username, Password: password}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	generatedToken := GeneratedToken{Token: tokenString}
	sessions.CreateSession(account.Account_Id, w, r)
	response := u.Message(true, "Successful log in")
	response["generatedToken"] = generatedToken
	return response
}

func RegisterUser(username, email, password string) map[string]interface{} {
	if validation.ValidateEmail(email) == false {
		return u.Message(false, "Invalid email address")
	}

	if email == "" {
		return u.Message(false, "Email address field is empty")
	}
	if validation.ValidatePassword(password) == false {
		return u.Message(false, "Password should have atleast 8 characters including small and capital letters")
	}
	if username == "" {
		return u.Message(false, "Username field is empty")
	}
	temp := &Accounts{}

	err := GetDB().Table("accounts").Where("email = ?", email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry")
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user.")
	}

	err = GetDB().Table("accounts").Where("username = ?", username).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry")
	}
	if temp.Username != "" {
		return u.Message(false, "Username already in use by another user.")
	}

	// err = GetDB().Table("accounts").Where("phone = ?", phone).First(temp).Error
	// if err != nil && err != gorm.ErrRecordNotFound {
	// 	return u.Message(false, "Connection error. Please retry")
	// }
	// if temp.Phone != "" {
	// 	return u.Message(false, "Phone number already in use by another user.")
	// }

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	saltedPassword := string(hashedPassword)
	userID := u.GenerateUID()
	if err := GetDB().Create(&Accounts{
		Account_Id: userID,
		Username:   username,
		Email:      email,
		// Phone:      phone,
		Password:   saltedPassword,
		Created_At: u.GenerateTimeNow()}).
		Error; err != nil {
		return u.Message(false, "Unable to create user account")
	}

	//creating posts folder
	directoryPath := "userFiles/" + userID + "/posts"
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		err = os.MkdirAll(directoryPath, 0777)
		if err != nil {
			return u.Message(false, "Error in created user folder.")
		}
	}

	//creating profile folder
	directory := "userFiles/" + userID + "/profile"
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err = os.MkdirAll(directory, 0777)
		if err != nil {
			return u.Message(false, "Error in created user folder.")
		}
	}
	tk := &Token{Username: username, Password: password}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	generatedToken := GeneratedToken{Token: tokenString}
	response := u.Message(true, "Your account is successfully created.")
	response["generatedToken"] = generatedToken
	return response
}

func Follow(follower, following string) map[string]interface{} {
	err := GetDB().Table("").Where("follower_id = ? AND following_id = ?", follower, following).First(&Followings{}).Error
		if err == gorm.ErrRecordNotFound {
		if err = GetDB().Create(&Followings{
			Id:           u.GenerateUID(),
			Following_Id: following,
			Follower_Id:  follower,
			Created_At:   u.GenerateTimeNow(),
		}).Error; err != nil {
			return u.Message(false, "You are unable to follow")
		}
		return u.Message(true, "You are now following")
	}
	return u.Message(true, "You are already following")
}
func UnFollow(follower, following string) map[string]interface{} {
	err := GetDB().Where(&Followings{Follower_Id: follower, Following_Id: following}).Delete(&Followings{}).Error
	if err != nil {
		return u.Message(true, "You are unable to unfollow now.Try again")
	}
	return u.Message(true, "You have unfollowed")
}
func CheckIffollowing(follower, following string) map[string]interface{} {

	type FollowStatus struct {
		Status bool
	}
	err := GetDB().Table("followings").Where("following_id = ? AND follower_id = ?", following, follower).First(&Followings{}).Error
	if err == gorm.ErrRecordNotFound {
		resp := u.Message(true, "Your are not following")
		resp["follow_status"] = &FollowStatus{Status: false}
		return resp
	} else {
		resp := u.Message(true, "Your are following")
		resp["follow_status"] = &FollowStatus{Status: true}
		return resp
	}

}
