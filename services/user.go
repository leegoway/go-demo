package services

import (
	"demo/models"
	"fmt"
	"demo/pkg/e"
	"demo/pkg/app"
)

type UserService struct {

}

type RegisterUserForm struct {
	FormValidator
	Uid           int64  `form:"uid" valid:"Required;Min(1)"`
	Username      string `form:"username" valid:"MaxSize(32)"`
	Idcard        string `form:"idcard" valid:"MaxSize(64)"`
	Mobile        string `form:"mobile" valid:"MaxSize(64)"`
}

func (us UserService) NewUser (form RegisterUserForm) (u *models.User, err error) {
	//合法性判断
	if valid := app.Valid(form); valid != e.SUCCESS {
		return nil, fmt.Errorf("参数错误，%s", e.GetMsg(valid))
	}

	u = &models.User{}
	u.Uid = form.Uid
	u.Username = form.Username
	u.Idcard = form.Idcard
	u.Mobile = form.Mobile

	//判断是否已存在
	models.FindOne(u, &u)
	if u != nil {
		return u, nil
	}

	//新建
	if err = u.Save(); err != nil {
		return nil, err
	}
	return u, nil
}

type QueryUserForm struct {
	FormValidator
	ID            int32  `form:"id"`
	Uid           int64  `form:"uid"`
	Username      string `form:"username" valid:"MaxSize(32)"`
	Idcard        string `form:"idcard" valid:"MaxSize(64)"`
	Mobile        string `form:"mobile" valid:"MaxSize(64)"`
}

func (us UserService)QueryUser(ucond QueryUserForm) (*models.User, error)  {
	if valid := ucond.ValidData(); valid != e.SUCCESS {
		return nil, fmt.Errorf("参数错误，%s", e.GetMsg(valid))
	}

	var u models.User
	u.ID = ucond.ID
	u.Uid = ucond.Uid
	u.Username = ucond.Username
	u.Idcard = ucond.Idcard
	u.Mobile = ucond.Mobile

	fmt.Println("service/user.go QueryUser", u)
	err := models.FindOne(u, &u)
	return &u, err
}
