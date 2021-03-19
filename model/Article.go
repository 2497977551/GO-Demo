package model

import (
	"ginblog/utils/ErrorInfo"
	uuid "github.com/satori/go.uuid"
)

type Article struct {
	Model
	Title    string `gorm:"column:Title;NOT NULL" json:"Title"`       // 文章标题
	CID      string `gorm:"column:CID;NOT NULL" json:"CID"`           // 文章分类
	Describe string `gorm:"column:Describe;NOT NULL" json:"Describe"` // 文章详情
	Content  string `gorm:"column:Content;NOT NULL" json:"Content"`   // 文章主体内容

}

// 添加文章
func AddArticle(title, describe, content, cid string) int {
	var a Article
	a.Id = uuid.NewV1()
	a.Title = title
	a.Describe = describe
	a.Content = content
	a.CID = cid
	err = db.Debug().Table("Article").Create(&a).Error
	if err != nil {
		return ErrorInfo.Error
	}
	return ErrorInfo.SucCse
}

// 查询文章
func QueryArticle(title string) (a []ArticleList, code int) {
	err = db.Table("Article").Where(" Title like ? ", "%"+title+"%").Find(&a).Error
	if err != nil {
		code = ErrorInfo.Error
		return
	}
	code = ErrorInfo.SucCse
	return
}

// 查询文章列表
func QueryAllArticle(pageSize, pageNum int) (a []ArticleList, code int, count int64) {

	err = db.Debug().Table("Article").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&a).Count(&count).Error
	if err != nil {
		code = ErrorInfo.Error
		return
	}
	code = ErrorInfo.SucCse
	return
}

// 编辑文章
func EditArticle(a EditArticles) int {
	err = db.Debug().Table("Article").Select("Title", "Describe", "Content", "CID", "UpdateTime").Updates(&a).Error
	if err != nil {
		return ErrorInfo.Error
	}
	return ErrorInfo.SucCse
}

// 删除文章
func DeleteArticle(id uuid.UUID) int {
	err = db.Debug().Table("Article").Where("id = ?", id).Delete(&Article{}).Error
	if err != nil {
		return ErrorInfo.Error
	}
	return ErrorInfo.SucCse
}
