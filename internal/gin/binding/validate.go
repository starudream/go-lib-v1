package binding

type Validator interface {
	Struct(v any) error
}

var DefaultValidator Validator

func validate(v any) error {
	if DefaultValidator == nil {
		return nil
	}
	return DefaultValidator.Struct(v)
}
