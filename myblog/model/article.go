package model

import "gorm.io/gorm"

//文章
type Article struct {
	Category   Category               `swaggerignore:"true" gorm:"foreignkey:Cid" ` //关联Category，并将Cid重写为外键
	gorm.Model `swaggerignore:"true"` //用于设置swagger文档，将此属性忽视，因为swagger不识别gorm.Model这个类型
	Title      string                 `gorm:"type:varchar(100);not null" json:"title" example:"title"`
	Cid        int                    `gorm:"type:int;not null" json:"cid" `
	Desc       string                 `gorm:"type:varchar(200);not null" json:"desc" example:"desc"`
	Content    string                 `gorm:"type:longtext;not null" json:"content" example:"content"`
	Img        string                 `gorm:"type:varchar(100);not null" json:"img" example:"img"`
	ReadCount  int                    `gorm:"type:int;not null;default:0" json:"read_ount" `
}

//文章评论
type ArtComment struct {
	Article    Article                `swaggerignore:"true" gorm:"foreignkey:Aid"` //关联Article，并将Aid重写为外键
	gorm.Model `swaggerignore:"true"` //用于设置swagger文档，将此属性忽视，因为swagger不识别gorm.Model这个类型
	Name       string                 `gorm:"type:varchar(20);not null" json:"name"`
	Content    string                 `gorm:"type:varchar(500);not null" json:"content"`
	Aid        int                    `gorm:"type:int;not null" json:"aid" `
}
