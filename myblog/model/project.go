package model

import "gorm.io/gorm"

//项目
type Project struct {
	gorm.Model     `swaggerignore:"true"`   //用于设置swagger文档，将此属性忽视，因为swagger不识别gorm.Model这个类型
	Name       string `gorm:"type:varchar(20);not null" json:"name" example:"name"`
	Desc       string `gorm:"type:varchar(50);not null" json:"desc" example:"desc"`
	Url        string `gorm:"type:varchar(50);not null" json:"url" example:"url"`
	Img        string `gorm:"type:varchar(100);not null" json:"img" example:"img"`
}
