package v10

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	v10 "github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
	"github.com/naufalfmm/fleet-management-system/internal/pkg/validator"
	"github.com/shopspring/decimal"
)

type v10Validator struct {
	v *v10.Validate
}

func NewV10() (validator.Validator, error) {
	langEn := en.New()
	langId := id.New()
	uni := ut.New(langEn, langEn, langId)
	trans, _ := uni.GetTranslator("en")

	validate := v10.New()
	if err := enTrans.RegisterDefaultTranslations(validate, trans); err != nil {
		return nil, err
	}

	if err := validate.RegisterValidation("dgte", ValidateDecimalGTE); err != nil {
		return nil, err
	}

	if err := validate.RegisterValidation("dlte", ValidateDecimalLTE); err != nil {
		return nil, err
	}

	if err := validate.RegisterValidation("dreq", ValidateDecimalReq); err != nil {
		return nil, err
	}

	return &v10Validator{
		v: validate,
	}, nil
}

func (v *v10Validator) ValidateStruct(i any) error {
	if err := v.v.Struct(i); err != nil {
		return err
	}
	return nil
}

func (v *v10Validator) Engine() any {
	return v.v
}

// ValidateDecimalGTE memvalidasi tag `dgte`
var ValidateDecimalGTE v10.Func = func(fl v10.FieldLevel) bool {
	field, ok := fl.Field().Interface().(decimal.Decimal)
	if !ok {
		return false
	}

	paramValue, err := decimal.NewFromString(fl.Param())
	if err != nil {
		return false // Parameter tag tidak valid
	}

	// Cmp mengembalikan 1 jika field > param, 0 jika sama, -1 jika field < param
	return field.Cmp(paramValue) >= 0
}

// ValidateDecimalLTE memvalidasi tag `dlte`
var ValidateDecimalLTE v10.Func = func(fl v10.FieldLevel) bool {
	field, ok := fl.Field().Interface().(decimal.Decimal)
	if !ok {
		return false
	}

	paramValue, err := decimal.NewFromString(fl.Param())
	if err != nil {
		return false
	}

	return field.Cmp(paramValue) <= 0
}

// ValidateDecimalReq memvalidasi tag `dreq` (tidak boleh nol/kosong)
var ValidateDecimalReq v10.Func = func(fl v10.FieldLevel) bool {
	field, ok := fl.Field().Interface().(decimal.Decimal)
	if !ok {
		return false
	}
	return !field.IsZero()
}
