package request

type Authenticate struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
