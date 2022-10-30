package service

import (
	"fmt"
	"myblog/dao"
	"myblog/errmsg"
	"myblog/model"

	"gorm.io/gorm"
)

//<----------文章模块逻辑处理------------>

// 后台创建文章
func CreateArticle(data *model.Article) int {
	var category model.Category
	err := dao.DB.Create(&data).Error
	err = dao.DB.Model(&category).Where("id=?", data.Cid).UpdateColumn("num", gorm.Expr("num+1")).Error
	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 后台删除文章
func DeleteArticle(id int) int {
	var art model.Article
	err := dao.DB.Where("id=?", id).Delete(&art).Error
	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 后台更新文章
func UpdateArticle(id int, data *model.Article) int {
	var article model.Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err := dao.DB.Model(&article).Where("id=?", id).Updates(&maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 前台查看文章
func GetArticleInfo(id int) (model.Article, int) {
	var article model.Article
	err := dao.DB.Preload("Category").Where("id=?", id).First(&article).Error
	err = dao.DB.Model(&article).Where("id=?", id).UpdateColumn("read_count", gorm.Expr("read_count+1")).Error
	if err != nil {
		return article, errmsg.ERROR_ARI_NOT_EXIST
	}
	return article, errmsg.SUCCESS
}

// 查看文章列表
func GetArticleList() ([]model.Article, int) {
	var articleList []model.Article

	err := dao.DB.Select("article.id", "title", "img", "updated_at", "created_at", "desc", "read_count", "Category.name").Order("article.created_at DESC").Joins("Category").Find(&articleList).Error
	if err != nil {
		fmt.Println(err)
		return nil, errmsg.ERROR
	}
	return articleList, errmsg.SUCCESS
}

// 根据标题检索文章
func SearchArticle(title string) ([]model.Article, int) {
	var articleList []model.Article
	err := dao.DB.Select("title", "img", "created_at", `desc`, "read_count", "Category.name").Order("Created_At").Joins("Category").Where("title LIKE ?", title+"%").Find(&articleList).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return articleList, errmsg.SUCCESS
}

// 查看分类文章
func GetCateArt(name string, pageSize int, pageNum int) ([]model.Article, int) {
	var cateArtList []model.Article

	err := dao.DB.Joins("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("Category.name=?", name).Find(&cateArtList).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST
	}
	return cateArtList, errmsg.SUCCESS
}
