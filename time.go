package svalidator

import (
	"time"
)

// TimeValidator is a validator for time.Time.
type TimeValidator struct {
	*Validator[time.Time]
}

func Time() *TimeValidator {
	return &TimeValidator{
		Validator: New[time.Time](),
	}
}

// After add a validate whether input value is after target.
func (t *TimeValidator) After(target func() time.Time) *TimeValidator {
	return t.AppendValidate(func(value time.Time) error {
		if value.After(target()) {
			return nil
		}
		return ErrTooSmall
	})
}

// After add a validate whether input value is after target date.
func (t *TimeValidator) AfterDate(target func() time.Time) *TimeValidator {
	return t.AppendValidate(func(value time.Time) error {
		tim := target()
		if time.Date(value.Year(), value.Month(), value.Day(), 0, 0, 0, 0, value.Location()).After(time.Date(tim.Year(), tim.Month(), tim.Day(), 0, 0, 0, 0, tim.Location())) {
			return nil
		}
		return ErrTooSmall
	})
}

// EqOrAfter add a validate whether input value is after or equal to target.
func (t *TimeValidator) EqOrAfter(target func() time.Time) *TimeValidator {
	return t.AppendValidate(func(value time.Time) error {
		if target().After(value) {
			return ErrTooSmall
		}
		return nil
	})
}

// EqOrAfter add a validate whether input value is after or equal to target date.
func (t *TimeValidator) EqOrAfterDate(target func() time.Time) *TimeValidator {
	return t.AppendValidate(func(value time.Time) error {
		tim := target()
		if time.Date(tim.Year(), tim.Month(), tim.Day(), 0, 0, 0, 0, tim.Location()).After(time.Date(value.Year(), value.Month(), value.Day(), 0, 0, 0, 0, value.Location())) {
			return ErrTooSmall
		}
		return nil
	})
}

// Before add a validate whether input value is before target.
func (t *TimeValidator) Before(target func() time.Time) *TimeValidator {
	return t.AppendValidate(func(value time.Time) error {
		if value.Before(target()) {
			return nil
		}
		return ErrTooBig
	})
}

// BeforeDate add a validate whether input value is before target date.
func (t *TimeValidator) BeforeDate(target func() time.Time) *TimeValidator {
	return t.AppendValidate(func(value time.Time) error {
		tim := target()
		if time.Date(value.Year(), value.Month(), value.Day(), 0, 0, 0, 0, value.Location()).Before(time.Date(tim.Year(), tim.Month(), tim.Day(), 0, 0, 0, 0, tim.Location())) {
			return nil
		}
		return ErrTooBig
	})
}

// EqOrBefore add a validate whether input value is before or equal to target.
func (t *TimeValidator) EqOrBefore(target func() time.Time) *TimeValidator {
	return t.AppendValidate(func(value time.Time) error {
		if target().Before(value) {
			return ErrTooBig
		}
		return nil
	})
}

// EqOrBeforeDate add a validate whether input value is before or equal to target date.
func (t *TimeValidator) EqOrBeforeDate(target func() time.Time) *TimeValidator {
	return t.AppendValidate(func(value time.Time) error {
		tim := target()
		if time.Date(tim.Year(), tim.Month(), tim.Day(), 0, 0, 0, 0, tim.Location()).Before(time.Date(value.Year(), value.Month(), value.Day(), 0, 0, 0, 0, value.Location())) {
			return ErrTooBig
		}
		return nil
	})
}

func (t *TimeValidator) Required() *TimeValidator {
	return t.AppendValidate(func(value time.Time) error {
		if value.IsZero() {
			return ErrEmpty
		}
		return nil
	})
}

func (t *TimeValidator) Equal(target func() time.Time) *TimeValidator {
	return t.AppendValidate(func(value time.Time) error {
		if value.Equal(target()) {
			return nil
		}
		return ErrNotEqual
	})
}

func (t *TimeValidator) EqualDate(target func() time.Time) *TimeValidator {
	return t.AppendValidate(func(value time.Time) error {
		tim := target()
		if tim.Year() == value.Year() && tim.Month() == value.Month() && tim.Day() == value.Day() {
			return nil
		}
		return ErrNotEqual
	})
}

func (t *TimeValidator) AppendValidate(funcs ...Validate[time.Time]) *TimeValidator {
	t.Validator = t.Validator.AppendValidate(funcs...)
	return t
}

// PointerTimeValidator is a validator for *time.Time.
type PointerTimeValidator struct {
	*Validator[*time.Time]
}

func PointerTime() *PointerTimeValidator {
	return &PointerTimeValidator{
		Validator: New[*time.Time](),
	}
}

// After add a validate whether input value is after target.
func (t *PointerTimeValidator) After(target func() time.Time) *PointerTimeValidator {
	return t.AppendValidate(func(value *time.Time) error {
		if value != nil && value.After(target()) {
			return nil
		}
		return ErrTooSmall
	})
}

// After add a validate whether input value is after target date.
func (t *PointerTimeValidator) AfterDate(target func() time.Time) *PointerTimeValidator {
	return t.AppendValidate(func(value *time.Time) error {
		tim := target()
		if value != nil && time.Date(value.Year(), value.Month(), value.Day(), 0, 0, 0, 0, value.Location()).After(time.Date(tim.Year(), tim.Month(), tim.Day(), 0, 0, 0, 0, tim.Location())) {
			return nil
		}
		return ErrTooSmall
	})
}

// EqOrAfter add a validate whether input value is after or equal to target.
func (t *PointerTimeValidator) EqOrAfter(target func() time.Time) *PointerTimeValidator {
	return t.AppendValidate(func(value *time.Time) error {
		if value != nil && target().After(*value) {
			return ErrTooSmall
		}
		return nil
	})
}

// EqOrAfter add a validate whether input value is after or equal to target date.
func (t *PointerTimeValidator) EqOrAfterDate(target func() time.Time) *PointerTimeValidator {
	return t.AppendValidate(func(value *time.Time) error {
		tim := target()
		if value != nil && time.Date(tim.Year(), tim.Month(), tim.Day(), 0, 0, 0, 0, tim.Location()).After(time.Date(value.Year(), value.Month(), value.Day(), 0, 0, 0, 0, value.Location())) {
			return ErrTooSmall
		}
		return nil
	})
}

// Before add a validate whether input value is before target.
func (t *PointerTimeValidator) Before(target func() time.Time) *PointerTimeValidator {
	return t.AppendValidate(func(value *time.Time) error {
		if value != nil && value.Before(target()) {
			return nil
		}
		return ErrTooBig
	})
}

// BeforeDate add a validate whether input value is before target date.
func (t *PointerTimeValidator) BeforeDate(target func() time.Time) *PointerTimeValidator {
	return t.AppendValidate(func(value *time.Time) error {
		tim := target()
		if value != nil && time.Date(value.Year(), value.Month(), value.Day(), 0, 0, 0, 0, value.Location()).Before(time.Date(tim.Year(), tim.Month(), tim.Day(), 0, 0, 0, 0, tim.Location())) {
			return nil
		}
		return ErrTooBig
	})
}

// EqOrBefore add a validate whether input value is before or equal to target.
func (t *PointerTimeValidator) EqOrBefore(target func() time.Time) *PointerTimeValidator {
	return t.AppendValidate(func(value *time.Time) error {
		if value != nil && target().Before(*value) {
			return ErrTooBig
		}
		return nil
	})
}

// EqOrBeforeDate add a validate whether input value is before or equal to target date.
func (t *PointerTimeValidator) EqOrBeforeDate(target func() time.Time) *PointerTimeValidator {
	return t.AppendValidate(func(value *time.Time) error {
		tim := target()
		if value != nil && time.Date(tim.Year(), tim.Month(), tim.Day(), 0, 0, 0, 0, tim.Location()).Before(time.Date(value.Year(), value.Month(), value.Day(), 0, 0, 0, 0, value.Location())) {
			return ErrTooBig
		}
		return nil
	})
}

func (t *PointerTimeValidator) Required() *PointerTimeValidator {
	return t.AppendValidate(func(value *time.Time) error {
		if value == nil || value.IsZero() {
			return ErrEmpty
		}
		return nil
	})
}

func (t *PointerTimeValidator) Equal(target func() time.Time) *PointerTimeValidator {
	return t.AppendValidate(func(value *time.Time) error {
		if value != nil && value.Equal(target()) {
			return nil
		}
		return ErrNotEqual
	})
}

func (t *PointerTimeValidator) EqualDate(target func() time.Time) *PointerTimeValidator {
	return t.AppendValidate(func(value *time.Time) error {
		tim := target()
		if value != nil && tim.Year() == value.Year() && tim.Month() == value.Month() && tim.Day() == value.Day() {
			return nil
		}
		return ErrNotEqual
	})
}

func (t *PointerTimeValidator) AppendValidate(funcs ...Validate[*time.Time]) *PointerTimeValidator {
	t.Validator = t.Validator.AppendValidate(funcs...)
	return t
}
