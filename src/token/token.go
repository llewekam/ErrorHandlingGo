package token

import "errors"

var (
	// ErrTokenInvalid indicates Invalid Token
	ErrTokenInvalid = errors.New("Token is invalid")

	// ErrTokenExpired indicates Expired token
	ErrTokenExpired = errors.New("Token has expired")

	// ErrTokenMismatch mismatch
	ErrTokenMismatch = errors.New("Token does not match something")
)

// DoTokenStuff an example
func DoTokenStuff(token string) (string, error) {
	if token == "" {
		panic("Token is unexpectedly empty")
	}

	if token == "invalid" {
		return "", ErrTokenInvalid
	}

	if token == "expired" {
		return "", ErrTokenExpired
	}

	if token == "mismatch" {
		return "", ErrTokenMismatch
	}

	return "Token done", nil
}
