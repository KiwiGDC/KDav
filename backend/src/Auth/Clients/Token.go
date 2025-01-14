package Clients

// New /*
/*
Constructor of Auth
@return: struct WebDavAuth
*/
func New(authType AuthType) WebDavAuth {
	auth := WebDavAuth{AuthType: authType}
	return auth
}

func (auth WebDavAuth) Validate(usernameField string, passwordField string) (bool, error) {
	userValidate, err := auth.validateUsernameField(usernameField)
	if err != nil {
		return false, err
	}
	passValidate, err := auth.validatePasswordField(passwordField)
	return userValidate && passValidate, nil
}
func (auth WebDavAuth) validatePasswordField(field string) (bool, error) {

}
func (auth WebDavAuth) validateUsernameField(field string) (bool, error) {

}
