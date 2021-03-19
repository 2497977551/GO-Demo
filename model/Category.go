package model

import (
	"ginblog/utils/ErrorInfo"
	uuid "github.com/satori/go.uuid"
)

type Category struct {
	Name string `gorm:"column:Name;NOT NULL" json:"Name"` // 文章类别名称
	Model
}

// 添加分类
func AddCategory(cgy Category) int {
	//cgy.Model.CreationTime = time.Now()
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

// 查询单个分类下的所有文章
func QueryCateBlog(id string) (a []ArticleList, code int) {
	err = db.Debug().Table("Article").Where("cid = ?", id).Order("CreationTime desc").Find(&a).Error
	if err != nil {
		code = ErrorInfo.Error
		return
	}
	code = ErrorInfo.SucCse
	return
}

// 编辑分类
func EditCate(cgy Category) int {
	var err error
	//cgy.Model.UpdateTime = time.Now()

	err = db.Debug().Table("category").Select("Name", "UpdateTime").Updates(&cgy).Error
	if err != nil {
		return ErrorInfo.Error
	}
	return ErrorInfo.SucCse
}

// 删除分类
func DeleteCate(id uuid.UUID) int {
	err := db.Debug().Table("category").Where("id = ?", id).Delete(&Category{}).Error
	if err != nil {
		return ErrorInfo.Error
	}
	return ErrorInfo.SucCse
}
