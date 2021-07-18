package model
type Book struct {
	Id int
	Name string `gorm:"unique_index;type:varchar(100)"`
	Price int `gorm:"unique;not null"`
	PageSize string `gorm:"index:pagesize"` //创建索引
	Sname string `gorm:"-"` //忽略字段
}

//迁移的时候重命名表明
func (Book) TableName() string  {
	return "test_book"
}