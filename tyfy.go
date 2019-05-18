package tyfy

const version = "1.0.0"

var (
	truthy = []string{"1", "true", "t", "yes", "y", "ok"}
	falsy  = []string{"0", "false", "f", "no", "n"}
)

// Truthy return the truthy value list
func Truthy() []string {
	return truthy
}

// Falsy return the falsy value list
func Falsy() []string {
	return falsy
}

// IsTruthy check the value is truthy
func IsTruthy(s string) bool {
	for _, t := range truthy {
		if s == t {
			return true
		}
	}
	return false
}

// IsFalsy check the value is falsy
func IsFalsy(s string) bool {
	for _, f := range falsy {
		if s == f {
			return true
		}
	}
	return false
}

// ExtendTruthy allow you to extend the truthy value list
func ExtendTruthy(ss []string) {
	truthy = append(truthy, ss...)
}

// ExtendFalsy allow you to extend the falsy value list
func ExtendFalsy(ss []string) {
	falsy = append(falsy, ss...)
}
