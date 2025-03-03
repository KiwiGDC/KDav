package Clients

// New /*
/*
Constructor of Auth
@return: struct WebDavAuth
*/

type TokenClient struct {
	WebDavAuth
}

func NewTokenClient() TokenClient {
	return TokenClient{WebDavAuth: WebDavAuth{AuthType: TOKEN}}
}

func (auth TokenClient) Validate(usernameField string, passwordField string) (bool, error) {
	userValidate, err := auth.validateUsernameField(usernameField)
	if err != nil {
		return false, err
	}
	passValidate, _ := auth.validatePasswordField(passwordField)
	return userValidate && passValidate, nil
}

// TODO: add verification field with DB CONNEXION
func (auth TokenClient) validatePasswordField(field string) (bool, error) {
	return true, nil
}

// TODO: add verification field with DB CONNEXION
func (auth TokenClient) validateUsernameField(field string) (bool, error) {
	return true, nil

}
