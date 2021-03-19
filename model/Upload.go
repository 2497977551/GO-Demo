package model

import (
	"context"
	"fmt"
	"ginblog/utils"
	"ginblog/utils/ErrorInfo"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

var accessKey = utils.AccessKey
var secretKey = utils.SecretKey
var bucKet = utils.BucKet
var imgURL = utils.QiniuSever

func UploadFile(file multipart.File, fileSize int64) (string, int) {
	// 客户端上传凭证
	putPolicy := storage.PutPolicy{
		Scope: bucKet,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	// 空间配置
	cfg := storage.Config{
		// 空间对应的机房
		Zone: &storage.ZoneHuanan,
		// 是否使用https
		UseHTTPS: false,
		// 是否使用加速
		UseCdnDomains: false,
	}
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{}
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		fmt.Println(err)
		return err.Error(), ErrorInfo.Error
	}
	url := imgURL + ret.Key
	fmt.Println(ret.Key)
	fmt.Println(ret.Hash)
	fmt.Println(ret.PersistentID)
	return url, ErrorInfo.SucCse
}
