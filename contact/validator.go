package contact

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
	lib "github.com/maxiloEmmmm/go-tool"
	"reflect"
	"strings"
)

type Trans struct {
	Tran ut.Translator
}

var Tranintance Trans

func init() {
	var zh2 = zh.New()
	var ValidateTrans, found = ut.New(zh2, zh2).GetTranslator("zh")

	if !found {
		fmt.Println("validator init fail, i18n package not found")
		return
	}

	Tranintance = Trans{Tran: ValidateTrans}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := zh_trans.RegisterDefaultTranslations(v, ValidateTrans); err != nil {
			fmt.Println(err)
		}

		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			j := fld.Tag.Get("json")

			if len(j) > 0 {
				// get json alias
				jTmp := strings.Split(j, ",")
				j = jTmp[0]
			}
			return lib.StringJoin(fld.Tag.Get("comment"), lib.StringJoin("(", lib.AssetsReturn(len(j) > 0, j, fld.Name).(string), ")"))
		})
	} else {
		fmt.Println("validator init fail, transform fail")
	}
}
