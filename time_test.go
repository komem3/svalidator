package svalidator_test

import (
	"testing"
	"time"

	"github.com/komem3/svalidator"
)

func TestTime_Validate(t *testing.T) {
	now := func() time.Time { return time.Date(2021, 1, 1, 0, 0, 1, 0, time.Local) }
	type (
		args struct {
			validator *svalidator.TimeValidator
			input     time.Time
		}
	)
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			"same time",
			args{
				validator: svalidator.Time().Equal(now),
				input:     now(),
			},
			nil,
		},
		{
			"not same time",
			args{
				validator: svalidator.Time().Equal(now),
				input:     now().Add(time.Second),
			},
			svalidator.ErrNotEqual,
		},
		{
			"same date",
			args{
				validator: svalidator.Time().EqualDate(now),
				input:     now().Add(time.Second),
			},
			nil,
		},
		{
			"not same date",
			args{
				validator: svalidator.Time().EqualDate(now),
				input:     now().AddDate(0, 0, 1),
			},
			svalidator.ErrNotEqual,
		},
		{
			"after time",
			args{
				validator: svalidator.Time().After(now),
				input:     now().Add(time.Second),
			},
			nil,
		},
		{
			"not after time",
			args{
				validator: svalidator.Time().After(now),
				input:     now(),
			},
			svalidator.ErrTooSmall,
		},
		{
			"after date",
			args{
				validator: svalidator.Time().AfterDate(now),
				input:     now().AddDate(0, 0, 1),
			},
			nil,
		},
		{
			"not after date",
			args{
				validator: svalidator.Time().AfterDate(now),
				input:     now().Add(time.Second),
			},
			svalidator.ErrTooSmall,
		},
		{
			"equal or after time",
			args{
				validator: svalidator.Time().EqOrAfter(now),
				input:     now().Add(time.Second),
			},
			nil,
		},
		{
			"equal or after time validate when input same time",
			args{
				validator: svalidator.Time().EqOrAfter(now),
				input:     now(),
			},
			nil,
		},
		{
			"not equal or after time",
			args{
				validator: svalidator.Time().EqOrAfter(now),
				input:     now().Add(-time.Second),
			},
			svalidator.ErrTooSmall,
		},
		{
			"equal or after date",
			args{
				validator: svalidator.Time().EqOrAfterDate(now),
				input:     now().Add(-time.Second),
			},
			nil,
		},
		{
			"not equal or after date",
			args{
				validator: svalidator.Time().EqOrAfterDate(now),
				input:     now().AddDate(0, 0, -1),
			},
			svalidator.ErrTooSmall,
		},
		{
			"before time",
			args{
				validator: svalidator.Time().Before(now),
				input:     now().Add(-time.Second),
			},
			nil,
		},
		{
			"not before time",
			args{
				validator: svalidator.Time().Before(now),
				input:     now(),
			},
			svalidator.ErrTooBig,
		},
		{
			"before date",
			args{
				validator: svalidator.Time().BeforeDate(now),
				input:     now().AddDate(0, 0, -1),
			},
			nil,
		},
		{
			"not before date",
			args{
				validator: svalidator.Time().BeforeDate(now),
				input:     now().Add(-time.Second),
			},
			svalidator.ErrTooBig,
		},
		{
			"equal or before time",
			args{
				validator: svalidator.Time().EqOrBefore(now),
				input:     now().Add(-time.Second),
			},
			nil,
		},
		{
			"equal or after time validate when input same time",
			args{
				validator: svalidator.Time().EqOrBefore(now),
				input:     now(),
			},
			nil,
		},
		{
			"not equal or before time",
			args{
				validator: svalidator.Time().EqOrBefore(now),
				input:     now().Add(time.Second),
			},
			svalidator.ErrTooBig,
		},
		{
			"equal or before date",
			args{
				validator: svalidator.Time().EqOrBeforeDate(now),
				input:     now().Add(time.Second),
			},
			nil,
		},
		{
			"not equal or before date",
			args{
				validator: svalidator.Time().EqOrBeforeDate(now),
				input:     now().AddDate(0, 0, 1),
			},
			svalidator.ErrTooBig,
		},
		{
			"pass required",
			args{
				validator: svalidator.Time().Required(),
				input:     now(),
			},
			nil,
		},
		{
			"empty",
			args{
				validator: svalidator.Time().Required(),
				input:     time.Time{},
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

func TestPointerTime_Validate(t *testing.T) {
	now := func() time.Time { return time.Date(2021, 1, 1, 0, 0, 1, 0, time.Local) }
	type (
		args struct {
			validator *svalidator.PointerTimeValidator
			input     *time.Time
		}
	)
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			"same time",
			args{
				validator: svalidator.PointerTime().Equal(now),
				input:     pointer(now()),
			},
			nil,
		},
		{
			"not same time",
			args{
				validator: svalidator.PointerTime().Equal(now),
				input:     pointer(now().Add(time.Second)),
			},
			svalidator.ErrNotEqual,
		},
		{
			"same date",
			args{
				validator: svalidator.PointerTime().EqualDate(now),
				input:     pointer(now().Add(time.Second)),
			},
			nil,
		},
		{
			"not same date",
			args{
				validator: svalidator.PointerTime().EqualDate(now),
				input:     pointer(now().AddDate(0, 0, 1)),
			},
			svalidator.ErrNotEqual,
		},
		{
			"after time",
			args{
				validator: svalidator.PointerTime().After(now),
				input:     pointer(now().Add(time.Second)),
			},
			nil,
		},
		{
			"not after time",
			args{
				validator: svalidator.PointerTime().After(now),
				input:     pointer(now()),
			},
			svalidator.ErrTooSmall,
		},
		{
			"after date",
			args{
				validator: svalidator.PointerTime().AfterDate(now),
				input:     pointer(now().AddDate(0, 0, 1)),
			},
			nil,
		},
		{
			"not after date",
			args{
				validator: svalidator.PointerTime().AfterDate(now),
				input:     pointer(now().Add(time.Second)),
			},
			svalidator.ErrTooSmall,
		},
		{
			"equal or after time",
			args{
				validator: svalidator.PointerTime().EqOrAfter(now),
				input:     pointer(now().Add(time.Second)),
			},
			nil,
		},
		{
			"equal or after time validate when input same time",
			args{
				validator: svalidator.PointerTime().EqOrAfter(now),
				input:     pointer(now()),
			},
			nil,
		},
		{
			"not equal or after time",
			args{
				validator: svalidator.PointerTime().EqOrAfter(now),
				input:     pointer(now().Add(-time.Second)),
			},
			svalidator.ErrTooSmall,
		},
		{
			"equal or after date",
			args{
				validator: svalidator.PointerTime().EqOrAfterDate(now),
				input:     pointer(now().Add(-time.Second)),
			},
			nil,
		},
		{
			"not equal or after date",
			args{
				validator: svalidator.PointerTime().EqOrAfterDate(now),
				input:     pointer(now().AddDate(0, 0, -1)),
			},
			svalidator.ErrTooSmall,
		},
		{
			"before time",
			args{
				validator: svalidator.PointerTime().Before(now),
				input:     pointer(now().Add(-time.Second)),
			},
			nil,
		},
		{
			"not before time",
			args{
				validator: svalidator.PointerTime().Before(now),
				input:     pointer(now()),
			},
			svalidator.ErrTooBig,
		},
		{
			"before date",
			args{
				validator: svalidator.PointerTime().BeforeDate(now),
				input:     pointer(now().AddDate(0, 0, -1)),
			},
			nil,
		},
		{
			"not before date",
			args{
				validator: svalidator.PointerTime().BeforeDate(now),
				input:     pointer(now().Add(-time.Second)),
			},
			svalidator.ErrTooBig,
		},
		{
			"equal or before time",
			args{
				validator: svalidator.PointerTime().EqOrBefore(now),
				input:     pointer(now().Add(-time.Second)),
			},
			nil,
		},
		{
			"equal or after time validate when input same time",
			args{
				validator: svalidator.PointerTime().EqOrBefore(now),
				input:     pointer(now()),
			},
			nil,
		},
		{
			"not equal or before time",
			args{
				validator: svalidator.PointerTime().EqOrBefore(now),
				input:     pointer(now().Add(time.Second)),
			},
			svalidator.ErrTooBig,
		},
		{
			"equal or before date",
			args{
				validator: svalidator.PointerTime().EqOrBeforeDate(now),
				input:     pointer(now().Add(time.Second)),
			},
			nil,
		},
		{
			"not equal or before date",
			args{
				validator: svalidator.PointerTime().EqOrBeforeDate(now),
				input:     pointer(now().AddDate(0, 0, 1)),
			},
			svalidator.ErrTooBig,
		},
		{
			"pass required",
			args{
				validator: svalidator.PointerTime().Required(),
				input:     pointer(now()),
			},
			nil,
		},
		{
			"empty",
			args{
				validator: svalidator.PointerTime().Required(),
				input:     pointer(time.Time{}),
			},
			svalidator.ErrEmpty,
		},
		{
			"nil is required error",
			args{
				validator: svalidator.PointerTime().Required(),
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
