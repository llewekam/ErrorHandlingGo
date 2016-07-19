package main

import "fmt"
import "os"
import "token"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from panic", r)
		}
	}()

	tokenString := commandLineToken()

	if msg, err := token.DoTokenStuff(tokenString); err == nil {
		// Do stuff with the successful token
		fmt.Println(msg)
	} else {
		// handle the token error
		fmt.Println("An error has occured")

		switch err {
		case token.ErrTokenExpired:
			fmt.Println("Token has expired")
			return

		case token.ErrTokenInvalid:
			fmt.Println("The token is invalid")
			return

		default:
			fmt.Println("An unknown error has occured", err)
		}

	}
}

func commandLineToken() string {
	if len(os.Args) > 1 {
		args := os.Args[1:]

		return args[0]
	}

	return "default"
}
