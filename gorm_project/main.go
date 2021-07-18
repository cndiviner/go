package main

import (
	"fmt"
	"gindemo/gorm_project/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main()  {
	db, err := gorm.Open("mysql", "root:OF9pt0qHkh@(127.0.0.1:3306)/hr?charset=utf8mb4&parseTime=True&loc=Local")
	if err!= nil{
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&model.User{},&model.Book{},&model.Project{},&model.Profile{})
	//创建表
	//db.CreateTable(&model.User{})
	//添加数据
	/*user := model.User{Name: "Jinzhu", Age: 18, Addr: "肖云路"}
	db.Create(&user)*/
	//查询
	// 获取第一条记录，按主键排序
	var user model.User
	db.First(&user,3) //默认ID查询
	fmt.Println(user)

	var user2 model.User
	db.First(&user2,"name=?","Jinzhu") //指定字段查询
	fmt.Println(user2)
	//更新单个
	var user3 model.User
	db.First(&user3,5)
	db.Model(&user3).Update("age",22)
	//更新多个
	var user4 model.User
	db.First(&user4,3)
	db.Model(&user4).Update(map[string]interface{}{"age":22,"addr":"三里屯"})
	//whereAll:=db.Where("name = ?", "jinzhu").Find(&user)
	//fmt.Println(whereAll)
	//删除
	var user5 model.User
	db.First(&user5,2) //默认ID查询
	db.Delete(&user5)
}
