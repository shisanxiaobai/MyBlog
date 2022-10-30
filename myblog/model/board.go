package model

import "gorm.io/gorm"

type Board struct {
	gorm.Model    `swaggerignore:"true"`    //用于设置swagger文档，将此属性忽视，因为swagger不识别gorm.Model这个类型
	Name       string `gorm:"varchar(20);not null;" json:"name"`
	Content    string `gorm:"varchar(500);not null;" json:"content"`
}
