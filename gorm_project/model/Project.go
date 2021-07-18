package model

import "github.com/jinzhu/gorm"

type Project struct {
	gorm.Model
	Name string
	Address string
	Status int
}

//迁移的时候重命名表明
func (Project) TableName() string  {
	return "cj_project"
}