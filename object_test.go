package svalidator_test

import (
	"fmt"
	"testing"

	"github.com/komem3/svalidator"
)

func TestSafeObject_BadStruct(t *testing.T) {
	type Sample struct {
		TestString1 string
		TestNumber  int
	}

	for _, tt := range []struct {
		name  string
		arg   svalidator.ValidatorMap[string]
		isErr bool
	}{
		{
			"type is not struct",
			svalidator.ValidatorMap[string]{
				"TestString1": svalidator.String().Required(),
			},
			true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := svalidator.SafeObject(tt.arg)
			assertIsError(t, tt.isErr, err)
		})
	}
}

func TestSafeObject_Error(t *testing.T) {
	type Sample struct {
		TestString1 string
		TestNumber  int
	}

	for _, tt := range []struct {
		name  string
		arg   svalidator.ValidatorMap[Sample]
		isErr bool
	}{
		{
			"pass",
			svalidator.ValidatorMap[Sample]{
				"TestString1": svalidator.String().Required(),
			},
			false,
		},
		{
			"invalid field name",
			svalidator.ValidatorMap[Sample]{
				"TestString": svalidator.String().Required(),
			},
			true,
		},
		{
			"invalid field type",
			svalidator.ValidatorMap[Sample]{
				"TestNumber": svalidator.String().Required(),
			},
			true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := svalidator.SafeObject(tt.arg)
			assertIsError(t, tt.isErr, err)
		})
	}
}

func TestObject_Validation(t *testing.T) {
	sampleErr := fmt.Errorf("bad")
	type stringType string
	type Nest struct {
		Name string
	}
	type Sample struct {
		TestString1    string
		TestString2    string
		TestString3    string
		TestStringType stringType
		Nest1          Nest
		Nest2          Nest
		IgnoreField    string
	}
	v := svalidator.Object(svalidator.ValidatorMap[Sample]{
		"TestString1":    svalidator.String().Min(4),
		"TestString2":    svalidator.String().Required(),
		"TestString3":    svalidator.String().Max(3),
		"TestStringType": svalidator.UString[stringType]().Max(1),
		"Nest1": svalidator.Object(svalidator.ValidatorMap[Nest]{
			"Name": svalidator.String().Max(4),
		}),
		"Nest2": svalidator.New(func(value Nest) error {
			if value.Name == "bad" {
				return sampleErr
			}
			return nil
		}),
	})

	for _, tt := range []struct {
		name string
		arg  Sample
		err  error
	}{
		{
			"pass",
			Sample{
				TestString1:    "1234",
				TestString2:    "1234",
				TestStringType: "a",
				Nest1:          Nest{Name: "1234"},
				Nest2:          Nest{Name: "12345"},
			},
			nil,
		},
		{
			"not pass",
			Sample{TestString1: "1234", TestString2: ""},
			svalidator.ErrObject{
				&svalidator.ErrObjectField{
					Field: "TestString2",
					Err:   svalidator.ErrEmpty,
				},
			},
		},
		{
			"multiple error",
			Sample{
				TestString1: "1234",
				TestString2: "",
				TestString3: "1244",
				Nest1:       Nest{Name: "12345"},
				Nest2:       Nest{Name: "bad"},
			},
			svalidator.ErrObject{
				&svalidator.ErrObjectField{
					Field: "TestString2",
					Err:   svalidator.ErrEmpty,
				},
				&svalidator.ErrObjectField{
					Field: "TestString3",
					Err:   svalidator.ErrTooBig,
				},
				&svalidator.ErrObjectField{
					Field: "Nest1",
					Err: svalidator.ErrObject{
						&svalidator.ErrObjectField{
							Field: "Name",
							Err:   svalidator.ErrTooBig,
						},
					},
				},
				&svalidator.ErrObjectField{
					Field: "Nest2",
					Err:   sampleErr,
				},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			err := v.Validate(tt.arg)
			assertError(t, tt.err, err)
		})
	}
}

func TestMap(t *testing.T) {
	type (
		args struct {
			validator *svalidator.MapValidator
			input     map[string]any
		}
	)
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			"pass",
			args{
				validator: svalidator.Map(svalidator.AnyValidatorMap{
					"ID":   svalidator.Number[int]().Min(1),
					"Name": svalidator.String().Max(255),
				}),
				input: map[string]any{
					"ID":   1,
					"Name": "taro",
				},
			},
			nil,
		},
		{
			"field does not exist",
			args{
				validator: svalidator.Map(svalidator.AnyValidatorMap{
					"ID":   svalidator.String().Required(),
					"Name": svalidator.String().Max(255),
				}),
				input: map[string]any{
					"ID": "test",
				},
			},
			svalidator.ErrNotExistsField,
		},
		{
			"invalid type",
			args{
				validator: svalidator.Map(svalidator.AnyValidatorMap{
					"ID": svalidator.Number[int]().Min(1),
				}),
				input: map[string]any{
					"ID": "test",
				},
			},
			svalidator.ErrInvalidType,
		},
		{
			"validation error",
			args{
				validator: svalidator.Map(svalidator.AnyValidatorMap{
					"ID": svalidator.Number[int]().Min(1),
				}),
				input: map[string]any{
					"ID": 0,
				},
			},
			svalidator.ErrObject{
				&svalidator.ErrObjectField{
					Field: "ID",
					Err:   svalidator.ErrTooSmall,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.validator.Validate(tt.args.input)
			assertError(t, tt.want, err)
		})
	}
}
