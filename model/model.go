package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/leegoway/go-demo/config"
	"time"
)

var db *gorm.DB

//初始化全局链接
func InitDB() error {
	var err error
	dbConfig := config.Cfg.Database
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.DbName))
	if err != nil {
		return err
	}
	db.LogMode(true)
	fmt.Println("已经连接数据库DB")
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(100) // 最大空闲连接数
	db.DB().SetMaxOpenConns(200) // 最大打开连接数
	db.DB().SetConnMaxLifetime(28000*time.Second) //否则会出现一次ErrInvalidConn
	return nil
}

func CloseDB() error {
	return db.Close()
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

func Save(m interface{}) error {
	err := db.Create(m).Error
	if err != nil {
		return err
	}
	//var i int32
	//err = db.Exec("SELECT LAST_INSERT_ID()", i).Error
	//if err != nil {
	//	return err
	//}
	//m.ID = i
	return nil
}

/*
* 查找一个
*/
func FindOne(conditions interface{}, out interface{}) error {
	err := db.Where(conditions).First(out).Error
	return err
}

/*
* 查找符合条件的所有
*/
func FindAll(conditions map[string]interface{}, out *interface{}) error {
	err := db.Where(conditions).Find(out).Error
	return err
}

/*
* 查找符合条件的分页数据
*/
func FindPage() {

}