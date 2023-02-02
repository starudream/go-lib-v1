package validator

import (
	"github.com/go-playground/validator/v10"

	ut "github.com/go-playground/universal-translator"

	enL "github.com/go-playground/locales/en"
	zhL "github.com/go-playground/locales/zh"

	enT "github.com/go-playground/validator/v10/translations/en"
	zhT "github.com/go-playground/validator/v10/translations/zh"
)

type ValidationErrors = validator.ValidationErrors

var (
	_v *validator.Validate

	_t ut.Translator

	_tZH ut.Translator
	_tEN ut.Translator
)

func init() {
	_v = validator.New()

	zh := zhL.New()
	en := enL.New()

	uti := ut.New(zh, en)

	_tZH, _ = uti.GetTranslator(zh.Locale())
	_ = zhT.RegisterDefaultTranslations(_v, _tZH)

	_tEN, _ = uti.GetTranslator(en.Locale())
	_ = enT.RegisterDefaultTranslations(_v, _tEN)

	SetT(EN())
}

func V() *validator.Validate {
	return _v
}

func T() ut.Translator {
	return _t
}

func ZH() ut.Translator {
	return _tZH
}

func EN() ut.Translator {
	return _tEN
}

func SetT(t ut.Translator) {
	_t = t
}
