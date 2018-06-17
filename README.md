# goutils

goutils is a set of common utilities for the Go language

[![CircleCI](https://circleci.com/gh/palsivertsen/goutils.svg?style=svg)](https://circleci.com/gh/palsivertsen/goutils) [![Code Climate](https://codeclimate.com/github/palsivertsen/goutils/badges/gpa.svg)](https://codeclimate.com/github/palsivertsen/goutils) [![Test Coverage](https://codeclimate.com/github/palsivertsen/goutils/badges/coverage.svg)](https://codeclimate.com/github/palsivertsen/goutils/coverage) [![Issue Count](https://codeclimate.com/github/palsivertsen/goutils/badges/issue_count.svg)](https://codeclimate.com/github/palsivertsen/goutils)

This library provides utilities for validation and convertion of types.
See[documentation](https://godoc.org/github.com/palsivertsen/goutils)

## Validators

To perform validations on a type you first need a validator

```go
myValidator := validators.String("qwerty")
```

Simple validations can be done using the functions prefixed with `IsXXX` or `HasXXX`. These functions returns a `bool`. This is most useful when you need to do a single validation validation

```go
myValidator.IsHex() // false
```

Validators can also be chained to do multiple validations in one go. Use the same functions as above, but without the `IsXXX` or `HasXXX` prefix and end with the special `HasErrors()` function

```go
myValidator.Hex().
	MaxLen(5).
	MinLen(2).
	HasErrors() // true
```

The check the results of the validations use the helper functions from the [validators.Validator](https://github.com/palsivertsen/goutils/blob/master/validators/validator.go) type

```go
myValidator.FirstError() // First registered error (Hex)
myValidator.Errors()     // All errors [Hex, MinLen]
```

### Implementations

goutils currently has validation implementations for the following types:
* string
* url.URL

Planned validators
* int (8/16/32/64)
* uint (8/16/32/64)
* float (32/64)

More validation functions and types to come.

## Converters

Convert from string to time.Time:

```go
fmt.Println(converters.String("1502835623824559499").MustTimeFromNsec())
fmt.Println(converters.String("1502835623").MustTimeFromSec())
```

## Contributions

Feature requests, feedback and pull requests welcome!
