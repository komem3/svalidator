package svalidator

import (
	"fmt"
	"reflect"
)

// ObjectValidator is validator for struct object.
type ObjectValidator[T any] struct {
	*Validator[T]
}

type ValidatorMap[T any] map[string]AnyValidator

type MapValidator struct {
	*Validator[map[string]any]
}

type AnyValidatorMap ValidatorMap[any]

// SafeObject returns ObjectValidator.
//
// This methods uses the value of the parameter to validate the structure
// of the argument ValidatorMap. If fails validation, this method returns error.
func SafeObject[T any](object ValidatorMap[T]) (*ObjectValidator[T], error) {
	var typ T
	rv := reflect.ValueOf(typ)
	if rv.Kind() != reflect.Struct {
		return nil, fmt.Errorf("only allow struct, but input %s", rv.Kind())
	}

	t := rv.Type()
	for field, validator := range object {
		stField, exist := t.FieldByName(field)
		if !exist {
			return nil, fmt.Errorf("%s does not exists in %T", field, typ)
		}

		vv := reflect.TypeOf(validator)
		vFunc, exist := vv.MethodByName("Validate")
		if !exist {
			panic(fmt.Sprintf("%s does not implement Validate", field))
		}

		if arg := vFunc.Type.In(1); stField.Type != arg {
			return nil, fmt.Errorf("struct field type is %s, but field Validator type is %s", stField.Type, arg)
		}
	}

	return &ObjectValidator[T]{
		Validator: New(func(value T) error {
			return object.validate(value)
		}),
	}, nil
}

// Object returns ObjectValidator.
//
// This validates structures as same safeObject, but will panic in case of error.
// Therefore, this is intended for global use.
func Object[T any](object ValidatorMap[T]) *ObjectValidator[T] {
	v, err := SafeObject(object)
	if err != nil {
		panic(err)
	}
	return v
}

func Map(object AnyValidatorMap) *MapValidator {
	return &MapValidator{
		Validator: New(func(value map[string]any) error {
			return object.validate(value)
		}),
	}
}

func (v ValidatorMap[T]) validate(value T) error {
	rv := reflect.ValueOf(value)
	rt := rv.Type()
	var merr []*ErrObjectField
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		validator, exists := v[field.Name]
		if !exists {
			continue
		}
		if err := validator.validateAny(rv.FieldByName(field.Name).Interface()); err != nil {
			merr = append(merr, newErrObjectField(field.Name, err))
		}
	}
	return newErrObject(merr...)
}

func (v AnyValidatorMap) validate(value map[string]any) error {
	var merr []*ErrObjectField
	for field, validator := range v {
		fieldValue, exists := value[field]
		if !exists {
			return fmt.Errorf("search %s field: %w", field, ErrNotExistsField)
		}
		vv := reflect.TypeOf(validator)
		vFunc, exist := vv.MethodByName("Validate")
		if !exist {
			panic(fmt.Sprintf("%s does not implement Validate", field))
		}

		fieldType := reflect.TypeOf(fieldValue)
		if arg := vFunc.Type.In(1); fieldType != arg {
			return fmt.Errorf("input %s type, but expected %s: %w", fieldType, arg, ErrInvalidType)
		}

		if err := validator.validateAny(fieldValue); err != nil {
			merr = append(merr, newErrObjectField(field, err))
		}
	}
	return newErrObject(merr...)
}
