package Clients

type Client interface {
	WebDavAuth
	New()
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
	AuthType AuthType

	// Optionnal field
	Username string
	Password string
	Token    string
}
