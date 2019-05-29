package services

import (
	"fmt"
	"github.com/leegoway/go-demo/model"
	"github.com/leegoway/go-demo/pkg/e"
	"github.com/leegoway/go-demo/pkg/app"
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

func (us UserService) NewUser (form RegisterUserForm) (u *model.User, err error) {
	//合法性判断
	if valid := app.Valid(form); valid != e.SUCCESS {
		return nil, fmt.Errorf("参数错误，%s", e.GetMsg(valid))
	}

	u = &model.User{}
	u.Uid = form.Uid
	u.Username = form.Username
	u.Idcard = form.Idcard
	u.Mobile = form.Mobile

	//判断是否已存在
	err = model.FindOne(u, &u)
	if u != nil {
		return u, err
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

func (us UserService)QueryUser(ucond QueryUserForm) (*model.User, error)  {
	if valid := ucond.ValidData(); valid != e.SUCCESS {
		return nil, fmt.Errorf("参数错误，%s", e.GetMsg(valid))
	}

	var u model.User
	u.ID = ucond.ID
	u.Uid = ucond.Uid
	u.Username = ucond.Username
	u.Idcard = ucond.Idcard
	u.Mobile = ucond.Mobile

	fmt.Println("service/user.go QueryUser", u)
	err := model.FindOne(u, &u)
	return &u, err
}
