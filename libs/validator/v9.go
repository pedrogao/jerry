package validator

import (
	"github.com/gin-gonic/gin/binding"
	cn "github.com/go-playground/locales/zh_Hans_CN"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
	"strings"
	"sync"
)

type AwesomeValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &AwesomeValidator{}

var (
	Trans ut.Translator
)

func New() *AwesomeValidator {
	awesomeValidator := new(AwesomeValidator)
	awesomeValidator.lazyinit()
	translator := cn.New()
	uni := ut.New(translator)
	Trans, _ = uni.GetTranslator("zh_Hans_CN")
	zh_translations.RegisterDefaultTranslations(awesomeValidator.validate, Trans)
	return awesomeValidator
}

func (v *AwesomeValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		v.lazyinit()
		if err := v.validate.Struct(obj); err != nil {
			return error(err)
		}
	}
	return nil
}

func (v *AwesomeValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *AwesomeValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")
		// add any custom validations etc. here
		v.validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			switch name {
			case "password":
				return "密码"
			case "nickname":
				return "昵称"
			case "confirm_password":
				return "确认密码"
			default:
				return name
			}
		})
	})
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
