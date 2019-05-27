package models

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
)

var db *gorm.DB

//初始化全局链接
func init() {
	var err error
	db, err = gorm.Open("mysql", "root:root123!@#@/ucenter?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("connect db error: ", err)
	}
	db.LogMode(true)
	fmt.Println("connected to db")
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

//通用Model结构体方法
type Model struct {
	ID int32 `json:"id" gorm:"AUTO_INCREMENT;primary_key;"`
}

/*
 * 保存到数据库
 */
func (m *Model) Save () error {
	if err := db.Create(m).Error; err != nil {
		return err
	}
	var i int32
	err := db.Exec("SELECT LAST_INSERT_ID()", i).Error
	if err != nil {
		return err
	}
	m.ID = i
	return nil
}

/*
* 从数据库删除
*/
func (m *Model) Delete () {

}

/*
* 更新数据库
*/
func (m *Model) Update() {

}

/*
* 查找一个
*/
func FindOne(conditions map[string]interface{}, out interface{}) {
	db.Where(conditions).First(out)
}

/*
* 查找符合条件的所有
*/
func FindAll(conditions map[string]interface{}, out *interface{}) {
	db.Where(conditions).Find(out)
}

/*
* 查找符合条件的分页数据
*/
func FindPage() {

}