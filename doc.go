//go:build ignore

// 快速翻译github.com/go-playground/validator/v10的验证结果为指定的语言
// Example:
// package main
// import (
//
//		"fmt"
//		"github.com/go-playground/locales"
//		ut "github.com/go-playground/universal-translator"
//		"github.com/go-playground/validator/v10"
//		"reflect"
//	)
//
//	type BagItem struct {
//		ID  int `json:"id" validate:"required,min=10" label:"背包ID"`
//		Num int `json:"num" validate:"required,min=10" label:"背包数量"`
//	}
//
//	type Foo struct {
//		Name string   `json:"name" validate:"required,max=32" label:"用户名"`
//		Age  int      `json:"age" validate:"required,min=18" label:"年龄"`
//		Bag  *BagItem `json:"bag" validate:"required" label:"背包"`
//	}
//
//	func init() {
//		err := Init(zh.New(), func(valid *validator.Validator, trans ut.Translator) error {
//			return zhTrans.RegisterDefaultTranslations(valid, trans)
//		})
//		if err != nil {
//			panic(err)
//		}
//	}
//
//	 func main() {
//		foo := Foo{
//			Name: "John1111111111111111111111111111111111111111111111111111111111",
//			Age:  -1,
//			Bag: &BagItem{
//				ID:  -1,
//				Num: -1,
//			},
//		}
//
//		err = TranslateError(Validator.Struct(foo))
//			fmt.Println(err)
//		}
//		// Output: 背包ID最小只能为10
package govalidate
