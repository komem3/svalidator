/*
svalidator is a validate package.
This package can validate against already defined types.

# Usage

By declaring validation globally,
you can validate the type before the code is actually called.

	package main

	import "github.com/komem3/svalidator"

	type Sample struct {
		Name string
	}
	var validator = svalidator.Object(svalidator.ValidatorMap[Sample]{
		"Name": svalidator.String().Max(255),
	})

	func main() {
		err := validator.Validate(Sample{
			ID:   "id",
			Name: "sample",
		})
		// ...
	}

The above code works fine, but will panic if the type is different as shown below.

	type Sample struct {
		Name int
	}
	var validator = svalidator.Object(svalidator.ValidatorMap[Sample]{ // panic!
		"Name": svalidator.String().Max(255),
	})

Of course, each validator can also be used directly.

	 	validator := svalidator.Number[int]().Min(2).Max(4)
		err := validator.Validate(4)

# Custom Validator

Also, you can use your definition in two way.

The first, easiest way is to use the [New] function.

	type Sample struct{ ID string }
	var validator = svalidator.Object(svalidator.ValidatorMap[Sample]{
		"ID": svalidator.New(func(value string) error {
			if value == "id" {
				return fmt.Errorf("invalid error")				
			}
			return nil
		}),
	})

Other way, you can also define a new validator as a structure.

	type StringID string

	type StringIDValidator struct {
		*Validator[StringID]
	}

	func NewStringIDValidator() *StringIDValidator {
		return &StringIDValidator{
			Validator: New(),
		}
	}

	func (s *StringIDValidator) Required() *StringIDValidator {
		s.Validator = s.Validator.AppendValidate(func(value StringIDValidator) error {
			if len(value) == 0 {
				return fmt.Errorf("id is required")		
			}
			return nil
		})
	}

	type Sample struct{ ID StringID }

	var validator = svalidator.Object(svalidator.ValidatorMap[Sample]{
		"ID": NewStringIDValidator().Required(),
	})

*/
package svalidator
