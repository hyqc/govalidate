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
