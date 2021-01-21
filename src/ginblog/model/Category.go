package model

import (
	"ginblog/utils/ErrorInfo"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Category struct {
	Name string `gorm:"column:Name;NOT NULL" json:"Name"` // 文章类别名称
	Model
}

// 添加分类
func AddCategory(cgy Category) int {
	cgy.Model.CreationTime = time.Now()
	cgy.Model.Id = uuid.NewV1()
	err := db.Debug().Table("Category").Select("ID", "Name", "CreationTime").Create(&cgy).Error
	if err != nil {
		return ErrorInfo.Error
	}
	return ErrorInfo.SucCse
}

// 查询分类列表
func CategoryList() (c []SaveList, code int) {

	err = db.Debug().Table("Category").Where("DeleteTime is null").Find(&c).Error
	if err != nil {
		code = ErrorInfo.Error
		return
	}
	code = ErrorInfo.SucCse
	return
}
