package Auth

import "fmt"

type AuthError struct {
	Code    int
	Message string
}

/*
Create a AuthError, code erreur always prefixed with "1"
@return string
*/
func (e AuthError) Error() string {
	return fmt.Sprintf("AuthError : %s (Error code: 1%d)", e.Message, e.Code)
}
