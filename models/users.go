package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UsersInfo struct {
	Model
	Uid      string `json:"uid" gorm:"column:uid"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	Avatar   string `json:"avatar" gorm:"column:avatar"`
}

type MUser struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

func AddUser(u UsersInfo) (err error) {
	cryptPwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(cryptPwd)
	return db.Create(&u).Error
}

func GetUserInfo(uid string) (user UsersInfo,  err error) {
	if err = db.Where("uid = ?", uid).First(&user).Error; err != nil && err != gorm.ErrRecordNotFound{
		return
	}
	err = nil
	return
}
