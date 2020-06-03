package contact

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
)

type Trans struct {
	Tran ut.Translator
}

var Tranintance Trans

func init() {
	var zh2 = zh.New()
	var ValidateTrans, _ = ut.New(zh2, zh2).GetTranslator("zh")

	Tranintance = Trans{Tran: ValidateTrans}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zh_trans.RegisterDefaultTranslations(v, Tranintance.Tran)
	}
}
