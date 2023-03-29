package svalidator_test

import (
	"testing"

	"github.com/komem3/svalidator"
)

func TestNumber_Validate(t *testing.T) {
	type (
		args struct {
			validator *svalidator.NumberValidator[int]
			input     int
		}
	)
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			"pass equal",
			args{
				validator: svalidator.Number[int]().Equal(2),
				input:     2,
			},
			nil,
		},
		{
			"not equal",
			args{
				validator: svalidator.Number[int]().Equal(2),
				input:     3,
			},
			svalidator.ErrNotEqual,
		},
		{
			"pass min",
			args{
				validator: svalidator.Number[int]().Min(2),
				input:     2,
			},
			nil,
		},
		{
			"too small",
			args{
				validator: svalidator.Number[int]().Min(2),
				input:     1,
			},
			svalidator.ErrTooSmall,
		},
		{
			"pass max",
			args{
				validator: svalidator.Number[int]().Max(2),
				input:     2,
			},
			nil,
		},
		{
			"too big",
			args{
				validator: svalidator.Number[int]().Max(2),
				input:     3,
			},
			svalidator.ErrTooBig,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.validator.Validate(tt.args.input)
			assertError(t, tt.want, err)
		})
	}
}

func TestPointerNumber_Validate(t *testing.T) {
	type (
		args struct {
			validator *svalidator.PointerNumberValidator[int]
			input     *int
		}
	)
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			"pass equal",
			args{
				validator: svalidator.PointerNumber[int]().Equal(2),
				input:     pointer(2),
			},
			nil,
		},
		{
			"not equal",
			args{
				validator: svalidator.PointerNumber[int]().Equal(2),
				input:     pointer(3),
			},
			svalidator.ErrNotEqual,
		},
		{
			"pass min",
			args{
				validator: svalidator.PointerNumber[int]().Min(2),
				input:     pointer(2),
			},
			nil,
		},
		{
			"too small",
			args{
				validator: svalidator.PointerNumber[int]().Min(2),
				input:     pointer(1),
			},
			svalidator.ErrTooSmall,
		},
		{
			"pass max",
			args{
				validator: svalidator.PointerNumber[int]().Max(2),
				input:     pointer(2),
			},
			nil,
		},
		{
			"too big",
			args{
				validator: svalidator.PointerNumber[int]().Max(2),
				input:     pointer(3),
			},
			svalidator.ErrTooBig,
		},
		{
			"pass required",
			args{
				validator: svalidator.PointerNumber[int]().Required(),
				input:     pointer(2),
			},
			nil,
		},
		{
			"empty",
			args{
				validator: svalidator.PointerNumber[int]().Required(),
				input:     nil,
			},
			svalidator.ErrEmpty,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.validator.Validate(tt.args.input)
			assertError(t, tt.want, err)
		})
	}
}
