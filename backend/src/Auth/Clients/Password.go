package Clients

// New /*
/*
Constructor of Auth
@return: struct WebDavAuth
*/

type PasswordClient struct {
	auth WebDavAuth
}

func NewPasswordClient() TokenClient {
	return TokenClient{WebDavAuth{AuthType: PASSWORD}}
}

func (auth PasswordClient) Validate(usernameField string, passwordField string) (bool, error) {
	userValidate, err := auth.validateUsernameField(usernameField)
	if err != nil {
		return false, err
	}
	passValidate, _ := auth.validatePasswordField(passwordField)
	return userValidate && passValidate, nil
}

// TODO: add verification field with DB CONNEXION
func (auth PasswordClient) validatePasswordField(field string) (bool, error) {
	return true, nil
}

// TODO: add verification field with DB CONNEXION
func (auth PasswordClient) validateUsernameField(field string) (bool, error) {
	return true, nil
}
