package Validator

import (
	"ginblog/utils/ErrorInfo"
	"github.com/go-playground/locales/zh_Hans_CN"
	uniTran "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTran "github.com/go-playground/validator/v10/translations/zh"
	"log"
	"reflect"
)

// 数据验证
func Validate(data interface{}) (string, int) {
	validate := validator.New()
	uni := uniTran.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")

	err := zhTran.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		log.Fatalln("err:", err.Error())
	}

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), ErrorInfo.Error
		}
	}
	return "", ErrorInfo.SucCse
}
