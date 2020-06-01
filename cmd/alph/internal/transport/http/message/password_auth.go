package message

type PasswordAuthRequest struct {
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required"`
	ClientID     string `json:"client_id" validate:"required"`
	ClientSecret string `json:"client_secret" validate:"required"`
}

type PasswordAuthResponse struct {
	AccessToken string `json:"access_token"`
}
