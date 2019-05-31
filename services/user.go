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
	Uid           int64  `form:"uid" json:"user" valid:"Required;Min(1)"`
	Username      string `form:"username" valid:"MaxSize(32)"`
	Idcard        string `form:"idcard" valid:"MaxSize(64)"`
	Mobile        string `form:"mobile" valid:"MaxSize(64)"`
}

func (us UserService) NewUser (form RegisterUserForm) (*model.User, *e.ErrorCode) {
	//合法性判断
	if err := app.Valid(form); err != nil {
		return nil, err
	}

	params := map[string]interface{} {
		"uid": form.Uid,
		"username": form.Username,
		"idcard": form.Idcard,
		"mobile": form.Mobile,
	}
	u := &model.User{}
	//判断是否已存在
	err := model.FindOne(params, &u)
	if err != nil {
		return u, e.Wrap(err)
	}
	if u.Uid != 0 {
		return u, nil
	}

	//新建
	if err = model.Save(params); err != nil {
		return nil, e.Wrap(err)
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

func (us UserService)QueryUser(ucond QueryUserForm) (*model.User, *e.ErrorCode)  {
	if valid := ucond.ValidData(); valid != e.SUCCESS {
		return nil, e.Wrap(fmt.Errorf("参数错误，%s", e.GetMsg(valid)))
	}

	var u model.User
	u.ID = ucond.ID
	u.Uid = ucond.Uid
	u.Username = ucond.Username
	u.Idcard = ucond.Idcard
	u.Mobile = ucond.Mobile

	fmt.Println("service/user.go QueryUser", u)
	err := model.FindOne(u, &u)
	if err != nil {
		return nil, e.Wrap(err)
	}
	return &u, nil
}
