package utils

import (
	"regexp"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	goutil "github.com/muhammadrivaldy/go-util"
)

func NewValidation() goutil.Validation {

	validation, _ := goutil.NewValidation()
	validation.RegisterValidation(validateIndonesiaPhoneNumber, validatePassword)
	validation.RegisterTranslation(translationIndonesiaPhoneNumber, translationPassword)

	return validation

}

func validateIndonesiaPhoneNumber(v *validator.Validate) (err error) {

	err = v.RegisterValidation("phone-number", func(fl validator.FieldLevel) bool {
		rgx, _ := regexp.Compile(`^\+62\d{9,12}$`)
		return rgx.MatchString(fl.Field().String())
	})
	if err != nil {
		return err
	}

	return nil

}

func validatePassword(v *validator.Validate) (err error) {

	err = v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		return regexp.MustCompile(`[A-Z]`).MatchString(password) &&
			regexp.MustCompile(`\d`).MatchString(password) &&
			regexp.MustCompile(`[^A-Za-z0-9]`).MatchString(password)
	})
	if err != nil {
		return err
	}

	return nil

}

func translationIndonesiaPhoneNumber(v *validator.Validate, trans *ut.Translator) error {

	return v.RegisterTranslation("phone-number", *trans, func(ut ut.Translator) error {
		return ut.Add("phone-number", "{0} format must be valid", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("phone-number", fe.Field())
		return t
	})

}

func translationPassword(v *validator.Validate, trans *ut.Translator) error {

	return v.RegisterTranslation("password", *trans, func(ut ut.Translator) error {
		return ut.Add("password", "{0} format must be valid", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("password", fe.Field())
		return t
	})

}
