package model

import "gorm.io/gorm"

//动态、生活
type Life struct {
	gorm.Model     `swaggerignore:"true"`   //用于设置swagger文档，将此属性忽视，因为swagger不识别gorm.Model这个类型
	Title      string `gorm:"type:varchar(20);not null" json:"title"`
	Desc       string `gorm:"type:varchar(20);not null" json:"desc"`
	Content    string `gorm:"type:longtext;" json:"content"`
	Img        string `gorm:"type:varchar(100);" json:"img"`
	ReadCount  int    `gorm:"type:varchar(20);default:0" json:"read_count"`
}

//生活评论
type LifeComment struct {
	Life       Life                   `gorm:"foreignkey:Lid"` //关联Life，并将Lid重写为外键
	gorm.Model `swaggerignore:"true"` //用于设置swagger文档，将此属性忽视，因为swagger不识别gorm.Model这个类型
	Name       string                 `gorm:"type:varchar(20);not null" json:"name"`
	Content    string                 `gorm:"type:longtext;not null" json:"content"`
	Lid        int                    `gorm:"type:uint;not null" json:"lid"`
}
