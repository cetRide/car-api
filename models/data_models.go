package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserInput struct {
	Username string `json : "username,omitempty"`
	Email    string `json : "email,omitempty"`
	Phone    string `json : "phone,omitempty"`
	Password string `json : "password,omitempty"`
	Comment  string `json : "comment,omitempty"`
	Status   string `json : "status,omitempty"`
	Caption  string `json : "caption,omitempty"`
}

type Token struct {
	Username string
	Password string
	jwt.StandardClaims
}

type GeneratedToken struct {
	Token string `json: "token"`
}

type Accounts struct {
	Account_Id string    `gorm:"primary_key;type:varchar(256)"`
	Username   string    `gorm:"not null;type:varchar(256)"`
	Email      string    `gorm:"not null;type:varchar(256)"`
	Phone      string    `gorm:"type:varchar(256)"`
	Password   string    `gorm:"not null;type:varchar(256)"`
	Profile    string    `gorm:"type:varchar(256)"`
	Status     string    `gorm:"type:varchar(256)"`
	Created_At time.Time `gorm:"not null;type:timestamp"`
}

type Followings struct {
	Id           string    `gorm:"primary_key;type:varchar(256)"`
	Following_Id string    `gorm:"not null;type:varchar(256)"`
	Follower_Id  string    `gorm:"not null;type:varchar(256)"`
	Created_At   time.Time `gorm:"not null;type:timestamp"`
}

type Posts struct {
	Post_Id    string    `gorm:"primary_key;type:varchar(256)"`
	User_Id    string    `gorm:"not null;type:varchar(256)"`
	Filename   string    `gorm:"not null;type:varchar(256)"`
	Caption    string    `gorm:"type:varchar(256)"`
	Created_At time.Time `gorm:"not null;type:timestamp"`
}

type Likes struct {
	Like_Id    string    `gorm:"primary_key;type:varchar(256)"`
	User_Id    string    `gorm:"not null;type:varchar(256)"`
	Post_Id    string    `gorm:"not null;type:varchar(256)"`
	Created_At time.Time `gorm:"not null;type:timestamp"`
}

type Comments struct {
	Comment_Id string    `gorm:"primary_key;type:varchar(256)"`
	User_Id    string    `gorm:"not null;type:varchar(256)"`
	Post_Id    string    `gorm:"not null;type:varchar(256)"`
	Comment    string    `gorm:"not null;type:varchar(256)"`
	Created_At time.Time `gorm:"not null;type:timestamp"`
}

type Replies struct{
	Reply_Id string `gorm:"primary_key;type:varchar(256)"`
	Reply string `gorm:"not null;type:varchar(256)"`
	User_Id string `gorm:"not null;type:varchar(256)"`
	Comment_Id string `gorm:"not null;type:varchar(256)"`
	Created_At time.Time `gorm:"not null;type:timestamp"`
}

type Reposts struct {
	Repost_Id  string    `gorm:"primary_key;type:varchar(256)"`
	User_Id    string    `gorm:"not null;type:varchar(256)"`
	Post_Id    string    `gorm:"not null;type:varchar(256)"`
	Created_At time.Time `gorm:"not null;type:timestamp"`
}
