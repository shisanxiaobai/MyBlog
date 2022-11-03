package dao

import (
	"fmt"
	"log"
	"myblog/conf"
	"myblog/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 用于外部调用
var DB *gorm.DB

// mysql的初始化操作
func init() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Config.User,
		conf.Config.DbPassword,
		conf.Config.Host,
		conf.Config.Port,
		conf.Config.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		//创建表时的命名策略
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}
	//迁移数据表，若没有此表则自定创建
	db.AutoMigrate(&model.Project{}, &model.Article{}, &model.Category{}, &model.Board{}, &model.ArtComment{}, &model.LifeComment{}, &model.Life{}, &model.Project{})

	// 获取通用数据库对象 sql.DB，然后使用其提供的功能
	sqlDB, _ := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	//连接暴露的接口
	DB = db
}
