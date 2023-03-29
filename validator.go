package svalidator

// New create Validator from Validate funcs.
func New[T any](funcs ...Validate[T]) *Validator[T] {
	return &Validator[T]{
		validFuncs: funcs,
	}
}

type AnyValidator interface {
	validateAny(v any) error
}

// Validator is a validator for generics type.
type Validator[T any] struct {
	validFuncs []Validate[T]
}

type Validate[T any] func(value T) error

// Validate validates value.
func (v *Validator[T]) Validate(value T) error {
	for _, f := range v.validFuncs {
		if err := f(value); err != nil {
			return &ErrValidate{Err: err, Input: value}
		}
	}
	return nil
}

// AppendValidate appends Validate func.
func (v *Validator[T]) AppendValidate(funcs ...Validate[T]) *Validator[T] {
	v.validFuncs = append(v.validFuncs, funcs...)
	return v
}

func (v *Validator[T]) validateAny(anyValue any) error {
	return v.Validate(anyValue.(T))
}
