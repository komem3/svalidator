package svalidator_test

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/komem3/svalidator"
)

type SampleID string

type Sample struct {
	ID   SampleID
	Name string
}

var validator = svalidator.Object(svalidator.ValidatorMap[Sample]{
	"ID":   svalidator.UString[SampleID]().Required(),
	"Name": svalidator.String().Max(255),
})

func ExampleObject() {
	err := validator.Validate(Sample{
		ID:   "id",
		Name: "sample",
	})
	fmt.Printf("err is %t\n", err != nil)

	err = validator.Validate(Sample{
		Name: "sample",
	})
	fmt.Printf("err is %t", err != nil)

	// Output:
	// err is false
	// err is true
}

func ExampleSafeObject() {
	_, err := svalidator.SafeObject(svalidator.ValidatorMap[Sample]{
		"ID":   svalidator.UString[SampleID]().Required(),
		"Name": svalidator.String().Max(255),
	})
	fmt.Printf("err is %t\n", err != nil)

	_, err = svalidator.SafeObject(svalidator.ValidatorMap[Sample]{
		"ID":   svalidator.Number[int]().Min(3),
		"Name": svalidator.String().Max(255),
	})
	fmt.Printf("err is %t", err != nil)

	// Output:
	// err is false
	// err is true
}

func ExampleString() {
	validator := svalidator.String().Required()

	err := validator.Validate("ok")
	fmt.Printf("err is %t\n", err != nil)

	err = validator.Validate("")
	fmt.Printf("err is %t\n", err != nil)

	// Output:
	// err is false
	// err is true
}

func ExampleNumber() {
	validator := svalidator.Number[int]().Min(2).Max(4)

	err := validator.Validate(4)
	fmt.Printf("err is %t\n", err != nil)

	err = validator.Validate(5)
	fmt.Printf("err is %t\n", err != nil)

	// Output:
	// err is false
	// err is true
}

func ExamplePointerNumber() {
	validator := svalidator.PointerNumber[int]().Required().Equal(4)

	err := validator.Validate(&[]int{4}[0])
	fmt.Printf("err is %t\n", err != nil)

	err = validator.Validate(nil)
	fmt.Printf("err is %t\n", err != nil)

	// Output:
	// err is false
	// err is true
}

func ExampleTime() {
	validator := svalidator.Time().EqOrAfter(time.Now)

	err := validator.Validate(time.Now().Add(time.Second))
	fmt.Printf("err is %t\n", err != nil)

	err = validator.Validate(time.Date(1900, 1, 10, 00, 0, 0, 0, time.Local))
	fmt.Printf("err is %t\n", err != nil)

	// Output:
	// err is false
	// err is true
}

func ExampleNew() {
	type Sample struct{ ID string }
	validator := svalidator.New(func(value Sample) error {
		if value.ID == "bad" {
			return fmt.Errorf("bad id")
		}
		return nil
	})

	err := validator.Validate(Sample{ID: "id"})
	fmt.Printf("err is %t\n", err != nil)

	err = validator.Validate(Sample{ID: "bad"})
	fmt.Printf("err is %t", err != nil)

	// Output:
	// err is false
	// err is true
}

func ExampleErrValidate() {
	validator := svalidator.String().Max(3)

	err := validator.Validate("hello")
	var verr *svalidator.ErrValidate
	if errors.As(err, &verr) {
		fmt.Printf("input is %s\n", verr.Input)
		fmt.Printf("required error is %t\n", errors.Is(err, svalidator.ErrTooBig))
	}
	// Output:
	// input is hello
	// required error is true
}

func ExampleErrObject() {
	err := validator.Validate(Sample{
		Name: strings.Repeat("あ", 256),
	})
	var verr svalidator.ErrObject
	if errors.As(err, &verr) {
		for _, ferr := range verr {
			fmt.Printf("%s field is error\n", ferr.Field)
		}
	}
	// Output:
	// ID field is error
	// Name field is error
}
