# svalidator

svalidator is a validate package.
This package can validate against already defined types.

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/komem3/svalidator.svg)](https://pkg.go.dev/github.com/komem3/svalidator)

## Example

```go
package main

import (
	"fmt"

	"github.com/komem3/svalidator"
)

type Sample struct {
	ID   SampleID
	Name string
}

var validator = svalidator.Object(svalidator.ValidatorMap[Sample]{
	"ID":   svalidator.UString[SampleID]().Required(),
	"Name": svalidator.String().Max(255),
})


func main() {
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
```
