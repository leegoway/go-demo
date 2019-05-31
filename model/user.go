package model

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

func (u *User) Save () error {
	if err := db.Create(u).Error; err != nil {
		return err
	}
	var i int32
	err := db.Exec("SELECT LAST_INSERT_ID()", i).Error
	if err != nil {
		return err
	}
	u.ID = i
	return nil
}

