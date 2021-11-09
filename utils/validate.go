package utils

import (
	"log"
	"reflect"

	"github.com/Peterliang233/techtrainingcamp-AppUpgrade/errmsg"

	"github.com/go-playground/locales/zh_Hans_CN"
	uniTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	TransZh "github.com/go-playground/validator/v10/translations/zh"
)

// Validate 验证器，验证参数格式是否合法
func Validate(data interface{}) (string, int) {
	validate := validator.New()
	uni := uniTrans.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN") // 指定翻译的语言
	// 验证器注册翻译器
	err := TransZh.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		log.Fatalf("err %s\n", err)
	}

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	err = validate.Struct(data)

	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.Error
		}
	}

	return "验证成功", errmsg.Success
}
