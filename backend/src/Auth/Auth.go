package Auth

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

// New /*
/*
Constructor of Auth
@return: struct WebDavAuth
*/
func New(authType AuthType) WebDavAuth {
	auth := WebDavAuth{AuthType: authType}
	return auth
}

func (auth WebDavAuth) WithToken(token string) (WebDavAuth, error) {
	if auth.AuthType != TOKEN || auth.AuthType == ANONYMOUS {
		return WebDavAuth{}, ErrAuthTypeNotRequireToken
	}
	auth.Token = token
	return auth, nil

}

func (auth WebDavAuth) WithPassword(password string) (WebDavAuth, error) {
	if auth.AuthType != PASSWORD || auth.AuthType == ANONYMOUS {
		return WebDavAuth{}, ErrAuthTypeNotRequirePassword
	}
	auth.Password = password
	return auth, nil
}

func (auth WebDavAuth) WithUsername(username string) (WebDavAuth, error) {
	if auth.AuthType == ANONYMOUS {
		return WebDavAuth{}, ErrAuthTypeAnonymousField
	}
	auth.Username = username
	return auth, nil
}
