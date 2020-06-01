package message

type PasswordAuthRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type PasswordAuthResponse struct {
	AccessToken string `json:"access_token"`
}
