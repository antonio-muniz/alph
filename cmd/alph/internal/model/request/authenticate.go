package request

type Authenticate struct {
	Username string `json:"subject_id"`
	Password string `json:"password"`
}
