package services

import (
	"demo/models"
	"fmt"
	"demo/pkg/e"
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
	if valid := form.ValidData(); valid != e.SUCCESS {
		return nil, fmt.Errorf("参数错误，%s", e.GetMsg(valid))
	}

	u = &models.User{}

	conds := map[string]interface{} {
		"Uid": form.Uid,
		"Username": form.Username,
		"Idcard": form.Idcard,
		"Mobile": form.Mobile,
	}
	//判断是否已存在
	models.FindOne(conds, u)
	if u != nil {
		return u, nil
	}

	//新建
	u.Uid = form.Uid
	u.Username = form.Username
	u.Idcard = form.Idcard
	u.Mobile = form.Mobile
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
	var u models.User
	conds := map[string]interface{}{
		"id": ucond.ID,
		"uid": ucond.Uid,
		"username": ucond.Username,
		"idcard": ucond.Idcard,
		"mobile": ucond.Mobile,
	}
	fmt.Println("service/user.go QueryUser", conds)
	models.FindOne(conds, &u)
	fmt.Println("from db", u)
	return &u, nil
}
