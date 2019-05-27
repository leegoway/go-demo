package services

import (
	"github.com/astaxie/beego/validation"
	"demo/pkg/e"
)

type FormValidator struct {

}

func (f *FormValidator) ValidData() int {
	valid := validation.Validation{}
	check, err := valid.Valid(f)
	if err != nil {
		return e.ERROR
	}
	if !check {
		return e.INVALID_PARAMS
	}
	return e.SUCCESS
}