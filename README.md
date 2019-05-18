<br><br>

<h1 align="center">TyFy</h1>

<p align="center">
  <a href="/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg"/></a>
  <a href="https://godoc.org/github.com/mingrammer/tyfy"><img src="https://godoc.org/github.com/mingrammer/tyfy?status.svg"/></a>
  <a href="https://goreportcard.com/report/github.com/mingrammer/tyfy"><img src="https://goreportcard.com/badge/github.com/mingrammer/tyfy"/></a>
  <a href="https://travis-ci.org/mingrammer/tyfy"><img src="https://travis-ci.org/mingrammer/tyfy.svg?branch=master"/></a>
</p>

<p align="center">
Truthy and falsy text validators
</p>
<br><br><br>

> A truthy value is a value that is considered true, and a falsy value is a value that is considered false.

tyfy (short of "truthy and falsy") provides two simple validators for the **truthy** and **falsy** text values (string only). It is useful for some programs that accept the human-readable true/false value as input from users, for example, the command line tools with a prompt such as `(y/n)`.

The default values of the truthy and falsy are `"1", "true", "t", "yes", "y", "ok"` and `"0", "false", "f", "no", "n"`, respectively.

## Installation

```
go get github.com/mingrammer/tyfy
```

## Usage

> See details in [GoDoc](https://godoc.org/github.com/mingrammer/tyfy)

```go
IsTruthy("yes")
IsTruthy("no")
// true
// false

IsFalsy("n")
IsFalsy("y")
// true
// false
```

You can extend the set of values as needed.

```go
ExtendTruthy([]string{"o"})
ExtendFalsy([]string{"x"})
```

## License

MIT
