package Clients

type Client interface {
	Validate(usernameField string, passwordField string) (bool, error)
	validatePasswordField(field string) (bool, error)
	validateUsernameField(field string) (bool, error)
}

type AuthType string

const (
	TOKEN     AuthType = "TOKEN"
	PASSWORD  AuthType = "PASSWORD"
	ANONYMOUS AuthType = "ANONYMOUS"
)

type WebDavAuth struct {
	AuthType      AuthType
	UsernameField string
	PasswordField string
}
