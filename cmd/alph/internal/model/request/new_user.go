package request

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}