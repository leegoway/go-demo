package models

import (
	"time"
)

type User struct {
	Model
	Uid int64 `json:"uid"`
	Username string `json:"username"`
	Idcard string `json:"idcard" gorm:"size:64"`
	Mobile string `json:"mobile"`
	CreateTime *time.Time `json:"create_time"`
	UpdateTime *time.Time `json:"update_time"`
}

func (User) TableName() string {
	return "user"
}

//func QueryUser(form services.RegisterUserForm) (*User, error) {
//	var u User
//	fmt.Println("check db", db)
//	if err := db.Where(form).First(&u).Error; err != nil {
//		return nil, err
//	}
//	return &u, nil
//}