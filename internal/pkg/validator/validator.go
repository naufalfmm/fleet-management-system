package validator

type Validator interface {
	ValidateStruct(i any) error
	Engine() any
}
