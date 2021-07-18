package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string
	Age int
	Addr string
}
// `Profile` 属于 `User`， `UserID` 是外键
type Profile struct {
	gorm.Model
	UserID int
	User   User
	Name   string
}