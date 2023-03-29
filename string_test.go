package svalidator_test

import (
	"os"
	"regexp"
	"testing"

	"github.com/komem3/svalidator"
)

func TestUString_Validate(t *testing.T) {
	type ustring string
	type (
		args struct {
			validator *svalidator.UStringValidator[ustring]
			input     ustring
		}
	)
	for _, tt := range []struct {
		name string
		args args
		err  error
	}{
		{
			"pass",
			args{
				validator: svalidator.UString[ustring]().
					Required().
					Max(10).
					Min(2).
					MatchRegex(regexp.MustCompile("[[:alpha:]]")),
				input: "😄a",
			},
			nil,
		},
		{
			"equal error",
			args{
				validator: svalidator.UString[ustring]().Equal("ok"),
				input:     "",
			},
			svalidator.ErrNotEqual,
		},
		{
			"required error",
			args{
				validator: svalidator.UString[ustring]().Required(),
				input:     "",
			},
			svalidator.ErrEmpty,
		},
		{
			"min error",
			args{
				validator: svalidator.UString[ustring]().Min(3),
				input:     "😄a",
			},
			svalidator.ErrTooSmall,
		},
		{
			"max error",
			args{
				validator: svalidator.UString[ustring]().Max(3),
				input:     "😄😄aa",
			},
			svalidator.ErrTooBig,
		},
		{
			"regex error",
			args{
				validator: svalidator.UString[ustring]().MatchRegex(regexp.MustCompile("^[[:alpha:]]+$")),
				input:     "a00a",
			},
			svalidator.ErrMismatchPattern,
		},
		{
			"enum pass",
			args{
				validator: svalidator.UString[ustring]().Enum([]ustring{"hello", "world"}),
				input:     "hello",
			},
			nil,
		},
		{
			"enum error",
			args{
				validator: svalidator.UString[ustring]().Enum([]ustring{"hello", "world"}),
				input:     "a00a",
			},
			svalidator.ErrMismatchPattern,
		},
		{
			"custom error",
			args{
				validator: svalidator.UString[ustring]().AppendValidate(func(value ustring) error {
					return os.ErrNotExist
				}),
				input: "1234",
			},
			os.ErrNotExist,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.validator.Validate(tt.args.input)
			assertError(t, tt.err, err)
		})
	}
}

func TestPointerUString_Validate(t *testing.T) {
	type ustring string
	type (
		args struct {
			validator *svalidator.PointerUStringValidator[ustring]
			input     *ustring
		}
	)
	for _, tt := range []struct {
		name string
		args args
		err  error
	}{
		{
			"pass",
			args{
				validator: svalidator.PointerUString[ustring]().Required().Max(10).Min(2),
				input:     pointer[ustring]("😄a"),
			},
			nil,
		},
		{
			"required error",
			args{
				validator: svalidator.PointerUString[ustring]().Equal("ok"),
				input:     pointer[ustring]("bad"),
			},
			svalidator.ErrNotEqual,
		},
		{
			"required error",
			args{
				validator: svalidator.PointerUString[ustring]().Required(),
				input:     nil,
			},
			svalidator.ErrEmpty,
		},
		{
			"min error",
			args{
				validator: svalidator.PointerUString[ustring]().Min(1),
				input:     pointer[ustring](""),
			},
			svalidator.ErrTooSmall,
		},
		{
			"max error",
			args{
				validator: svalidator.PointerUString[ustring]().Max(3),
				input:     pointer[ustring]("😄😄aa"),
			},
			svalidator.ErrTooBig,
		},
		{
			"enum error",
			args{
				validator: svalidator.PointerUString[ustring]().Enum([]ustring{"hello", "world"}),
				input:     pointer[ustring]("ok"),
			},
			svalidator.ErrMismatchPattern,
		},
		{
			"custom error",
			args{
				validator: svalidator.PointerUString[ustring]().AppendValidate(func(value *ustring) error {
					return os.ErrNotExist
				}),
				input: pointer[ustring]("1234"),
			},
			os.ErrNotExist,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.validator.Validate(tt.args.input)
			assertError(t, tt.err, err)
		})
	}
}

func TestString_Validate(t *testing.T) {
	type (
		args struct {
			validator *svalidator.StringValidator
			input     string
		}
	)
	for _, tt := range []struct {
		name string
		args args
		err  error
	}{
		{
			"pass",
			args{
				validator: svalidator.String().Required().Max(10).Min(2),
				input:     "😄a",
			},
			nil,
		},
		{
			"required error",
			args{
				validator: svalidator.String().Required(),
				input:     "",
			},
			svalidator.ErrEmpty,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.validator.Validate(tt.args.input)
			assertError(t, tt.err, err)
		})
	}
}
