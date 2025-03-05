package common

import (
	"reflect"

	_validator "gopkg.in/validator.v2"
)

var (
	nonZeroValidator = "nonzero"
)

type Validator struct{}

type IValidatedStruct interface {
	Validate() error
}

func (v *Validator) SetupValidator() error {
	if err := _validator.SetValidationFunc(nonZeroValidator, v.NonZero); err != nil {
		return err
	}
	return nil
}

func (*Validator) NonZero(value any, _ string) error {
	rValue := reflect.ValueOf(value)
	if rValue.Kind() == reflect.Uint64 {
		if rValue.Uint() == 0 {
			return _validator.ErrZeroValue
		}
	}
	if rValue.Kind() == reflect.String {
		if rValue.String() == "" {
			return _validator.ErrZeroValue
		}
	}
	return nil
}
