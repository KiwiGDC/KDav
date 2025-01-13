package Auth

/*
*
Auth Error Sentinel :
Code prefix : 1
Generic usage
*/
var (
	// For AuthType Password, the token is not required.
	ErrAuthTypeNotRequireToken = AuthError{Code: 1, Message: "token is not required for this auth type"}
	// Same, for Token Authtype, the password is not required.
	ErrAuthTypeNotRequirePassword = AuthError{Code: 2, Message: "password is not required for this auth type"}
	// For anonymous, token and password and username is not required
	ErrAuthTypeAnonymousField = AuthError{Code: 3, Message: "token, password or username field is not allowed for this auth type 'Anonymous'"}
)
