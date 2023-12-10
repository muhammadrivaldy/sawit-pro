package utils

import (
	"regexp"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func ValidatorNew() *validator.Validate {

	validatorObj := validator.New()
	validatorObj.RegisterValidation("phone-number", validateIndonesiaPhoneNumber)

	en := en.New()
	ut := ut.New(en, en)
	trans, _ := ut.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(validatorObj, trans)

	translateOverride(trans, validatorObj)

	return validatorObj

}

func validateIndonesiaPhoneNumber(fl validator.FieldLevel) bool {

	rgx, _ := regexp.Compile(`^\+62\d{9,12}$`)

	return rgx.MatchString(fl.Field().String())

}

func translateOverride(trans ut.Translator, validatorObj *validator.Validate) {

	validatorObj.RegisterTranslation("phone-number", trans, func(ut ut.Translator) error {
		return ut.Add("phone-number", "{0} format must be valid!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("phone-number", fe.Field())
		return t
	})

}
