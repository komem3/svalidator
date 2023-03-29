package svalidator

type OrderedNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type NumberValidator[T OrderedNumber] struct {
	*Validator[T]
}

func Number[T OrderedNumber]() *NumberValidator[T] {
	return &NumberValidator[T]{Validator: New[T]()}
}

func (n *NumberValidator[T]) Min(num T) *NumberValidator[T] {
	return n.Append(func(value T) error {
		if num > value {
			return ErrTooSmall
		}
		return nil
	})
}

func (n *NumberValidator[T]) Max(num T) *NumberValidator[T] {
	return n.Append(func(value T) error {
		if num < value {
			return ErrTooBig
		}
		return nil
	})
}

func (n *NumberValidator[T]) Equal(num T) *NumberValidator[T] {
	return n.Append(func(value T) error {
		if num != value {
			return ErrNotEqual
		}
		return nil
	})
}

func (n *NumberValidator[T]) Append(validates ...Validate[T]) *NumberValidator[T] {
	n.validFuncs = append(n.validFuncs, validates...)
	return n
}

type PointerNumberValidator[T OrderedNumber] struct {
	*Validator[*T]
}

func PointerNumber[T OrderedNumber](validates ...Validate[*T]) *PointerNumberValidator[T] {
	return &PointerNumberValidator[T]{Validator: New[*T](validates...)}
}

func (n *PointerNumberValidator[T]) Min(num T) *PointerNumberValidator[T] {
	return n.Append(func(value *T) error {
		if value != nil && num > *value {
			return ErrTooSmall
		}
		return nil
	})
}

func (n *PointerNumberValidator[T]) Max(num T) *PointerNumberValidator[T] {
	return n.Append(func(value *T) error {
		if value != nil && num < *value {
			return ErrTooBig
		}
		return nil
	})
}

func (n *PointerNumberValidator[T]) Equal(num T) *PointerNumberValidator[T] {
	return n.Append(func(value *T) error {
		if value != nil && num != *value {
			return ErrNotEqual
		}
		return nil
	})
}

func (n *PointerNumberValidator[T]) Required() *PointerNumberValidator[T] {
	return n.Append(func(value *T) error {
		if value == nil {
			return ErrEmpty
		}
		return nil
	})
}

func (n *PointerNumberValidator[T]) Append(validates ...Validate[*T]) *PointerNumberValidator[T] {
	n.validFuncs = append(n.validFuncs, validates...)
	return n
}
