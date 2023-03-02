package datavalidate

import (
	"errors"
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func ValidateData(any interface{}) error {
	// 中文翻译器
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")

	// 校验器

	validate := validator.New()

	// 注册翻译器到校验器
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
	}
	err = validate.Struct(any)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Translate(trans))
		}
	}
	return nil
}
