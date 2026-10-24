package common

import (
	"errors"
	"fmt"
)

type errMissingParam struct {
	param string
}

func (e errMissingParam) Error() string {
	return fmt.Errorf("%s of %s", ErrMissingParamCmp.Error(), e.param).Error()
}

func (e errMissingParam) Unwrap() error {
	return ErrMissingParamCmp
}

func ErrMissingParam(param string) error {
	return &errMissingParam{
		param: param,
	}
}

type errNoLocation struct {
	vehicleID string
}

func (e errNoLocation) Error() string {
	return fmt.Errorf("%s for %s", ErrNoLocationCmp.Error(), e.vehicleID).Error()
}

func (e errNoLocation) Unwrap() error {
	return ErrNoLocationCmp
}

func ErrNoLocation(vehicleID string) error {
	return &errNoLocation{
		vehicleID: vehicleID,
	}
}

var (
	ErrMissingParamCmp = errors.New("missing param")
	ErrNoLocationCmp   = errors.New("no location found")

	ErrDefaultError = errors.New("something error happen")
)

const (
	SuccessCodeOK = 200000

	ErrorCodeMissingParam = 400001
	ErrorCodeNoLocation   = 400002

	ErrorCodeInternalError = 500000
	ErrorCodePanicRecover  = 500999
)
