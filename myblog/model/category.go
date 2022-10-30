package model

//文章分类
type Category struct {
	ID   int    `swaggerignore:"true" gorm:"primary_key;auto_increment" json:"id" example:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name" example:"name"`
	Num  int    `swaggerignore:"true" gorm:"type:int;default:0" json:"num" example:"num"`
}
