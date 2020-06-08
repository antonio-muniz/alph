package message

type PasswordAuthRequest struct {
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required"`
	ClientID     string `json:"client_id" validate:"required"`
	ClientSecret string `json:"client_secret" validate:"required"`
}

func (r PasswordAuthRequest) NewPointer() interface{} {
	return &PasswordAuthRequest{}
}

func (r PasswordAuthRequest) Dereference(pointer interface{}) interface{} {
	return *(pointer.(*PasswordAuthRequest))
}

type PasswordAuthResponse struct {
	AccessToken string `json:"access_token"`
}
