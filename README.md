goutils is a set of common utilities for the Go language

[![CircleCI](https://circleci.com/gh/palsivertsen/goutils.svg?style=svg)](https://circleci.com/gh/palsivertsen/goutils)

[![Code Climate](https://codeclimate.com/github/palsivertsen/goutils/badges/gpa.svg)](https://codeclimate.com/github/palsivertsen/goutils)

[![Test Coverage](https://codeclimate.com/github/palsivertsen/goutils/badges/coverage.svg)](https://codeclimate.com/github/palsivertsen/goutils/coverage)

[![Issue Count](https://codeclimate.com/github/palsivertsen/goutils/badges/issue_count.svg)](https://codeclimate.com/github/palsivertsen/goutils)

# Usage

## Validators

Some simple validations:

```go
qwerty := validators.String("qwerty")
fmt.Println("Is a valid hex string:", qwerty.IsHex())
fmt.Println("Matches 'qwer' pattern:", qwerty.IsMatch("qwer"))
```

Validators can also be chained:

```go
hex := validators.String("abc123").
    Hex().
    Match("ABC").
    Match("321")
fmt.Println("Has errors:", hex.HasErrors())
fmt.Println("First error was:", hex.FirstError())
fmt.Println("All errors:", hex.Errors())
```

## Converters

Convert from string to time.Time:

```go
fmt.Println(converters.String("1502835623824559499").MustTimeFromNsec())
fmt.Println(converters.String("1502835623").MustTimeFromSec())
```
