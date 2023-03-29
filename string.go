package svalidator

import (
	"regexp"
	"unicode/utf8"
)

// UStringValidator is a validator for underlying string type.
type UStringValidator[T ~string] struct {
	*Validator[T]
}

func UString[T ~string]() *UStringValidator[T] {
	return &UStringValidator[T]{
		Validator: New[T](),
	}
}

type StringValidator = UStringValidator[string]

// String is a validator for primitive string type.
func String() *StringValidator {
	return UString[string]()
}

func (s *UStringValidator[T]) Max(m int) *UStringValidator[T] {
	return s.AppendValidate(func(value T) error {
		if utf8.RuneCountInString(string(value)) > m {
			return ErrTooBig
		}
		return nil
	})
}

func (s *UStringValidator[T]) Min(m int) *UStringValidator[T] {
	return s.AppendValidate(func(value T) error {
		if utf8.RuneCountInString(string(value)) < m {
			return ErrTooSmall
		}
		return nil
	})
}

func (s *UStringValidator[T]) Required() *UStringValidator[T] {
	return s.AppendValidate(func(value T) error {
		if len(value) == 0 {
			return ErrEmpty
		}
		return nil
	})
}

func (s *UStringValidator[T]) MatchRegex(re *regexp.Regexp) *UStringValidator[T] {
	return s.AppendValidate(func(value T) error {
		if !re.MatchString(string(value)) {
			return ErrMismatchPattern
		}
		return nil
	})
}

func (s *UStringValidator[T]) Equal(str T) *UStringValidator[T] {
	return s.AppendValidate(func(value T) error {
		if value != str {
			return ErrNotEqual
		}
		return nil
	})
}

func (s *UStringValidator[T]) Enum(enum []T) *UStringValidator[T] {
	return s.AppendValidate(func(value T) error {
		for _, t := range enum {
			if value == t {
				return nil
			}
		}
		return ErrMismatchPattern
	})
}

func (s *UStringValidator[T]) AppendValidate(funcs ...Validate[T]) *UStringValidator[T] {
	s.Validator = s.Validator.AppendValidate(funcs...)
	return s
}

// PointerUStringValidator is a validator for pointer of underlying string type.
type PointerUStringValidator[T ~string] struct {
	*Validator[*T]
}

func PointerUString[T ~string]() *PointerUStringValidator[T] {
	return &PointerUStringValidator[T]{
		Validator: New[*T](),
	}
}

type PointerStringValidator = PointerUStringValidator[string]

// PointerString is a validator for pointer of primitive string type.
func PointerString() *PointerStringValidator {
	return PointerUString[string]()
}

// Max adds max validate.
// If input is nil, returns nil.
func (s *PointerUStringValidator[T]) Max(m int) *PointerUStringValidator[T] {
	return s.AppendValidate(func(value *T) error {
		if value != nil && utf8.RuneCountInString(string(*value)) > m {
			return ErrTooBig
		}
		return nil
	})
}

// Min adds min validate.
// If input is nil, returns nil.
func (s *PointerUStringValidator[T]) Min(m int) *PointerUStringValidator[T] {
	return s.AppendValidate(func(value *T) error {
		if value != nil && utf8.RuneCountInString(string(*value)) < m {
			return ErrTooSmall
		}
		return nil
	})
}

func (s *PointerUStringValidator[T]) Required() *PointerUStringValidator[T] {
	return s.AppendValidate(func(value *T) error {
		if value == nil {
			return ErrEmpty
		}
		return nil
	})
}

func (s *PointerUStringValidator[T]) MatchRegex(re *regexp.Regexp) *PointerUStringValidator[T] {
	return s.AppendValidate(func(value *T) error {
		if value != nil && !re.MatchString(string(*value)) {
			return ErrMismatchPattern
		}
		return nil
	})
}

func (s *PointerUStringValidator[T]) Equal(str T) *PointerUStringValidator[T] {
	return s.AppendValidate(func(value *T) error {
		if value != nil && *value != str {
			return ErrNotEqual
		}
		return nil
	})
}

func (s *PointerUStringValidator[T]) Enum(enum []T) *PointerUStringValidator[T] {
	return s.AppendValidate(func(value *T) error {
		if value == nil {
			return nil
		}
		for _, t := range enum {
			if *value == t {
				return nil
			}
		}
		return ErrMismatchPattern
	})
}

func (s *PointerUStringValidator[T]) AppendValidate(funcs ...Validate[*T]) *PointerUStringValidator[T] {
	s.Validator = s.Validator.AppendValidate(funcs...)
	return s
}
