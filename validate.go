package govalidate

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

// TranslateError 翻译错误信息为注册的语言
func TranslateError(err error) error {
	if err == nil {
		return nil
	}
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}
	em := errs.Translate(trans)
	for _, e := range em {
		return errors.New(e)
	}
	return err
}

// ValidateStructWithRules 根据传入的规则验证结构体数据
// data 结构体值
// rules 规则
func ValidateStructWithRules(data interface{}, rules *Rules) error {
	for _, item := range rules.Items {
		Validator.RegisterStructValidationMapRules(item.Rules, item.Type)
	}
	return TranslateError(Validator.Struct(data))
}

// ValidateStruct 执行自定义验证验证结构体数据
// data 结构体值
// call 自定义验证函数，示例：
// Example:
//
//	type BagItem struct {
//		ID  int `json:"id"`
//		Num int `json:"num"`
//	}
//
//	type Foo struct {
//		Name string   `json:"name" `
//		Age  int      `json:"age" `
//		Bag  *BagItem `json:"bag" `
//	}
//
//	err := Init(zh.New(), func(valid *validator.Validate, trans ut.Translator) error {
//		return zhTrans.RegisterDefaultTranslations(valid, trans)
//	})
//	if err != nil {
//		panic(err)
//	}
//
//	foo := Foo{
//		Name: "John1111111111111111111111111111111111111111111111111111111111",
//		Age:  20,
//		Bag: &BagItem{
//			ID:  12,
//			Num: 13,
//		},
//	}
//
//	reqAddr := func(data interface{}) error {
//		rules := &Rules{
//			Items: []*StructRules{
//				{
//					Type: Foo{},
//					Rules: map[string]string{
//						"ID":    "required",
//						"Title": "required",
//					},
//				},
//			},
//		}
//
//		return ValidateStructWithRules(data, rules)
//	}
//
//	err = ValidateStruct(foo, reqAddr)
func ValidateStruct(data interface{}, call ...ValidatorFunc) error {
	for _, handler := range call {
		if err := handler(data); err != nil {
			return err
		}
	}
	return nil
}
