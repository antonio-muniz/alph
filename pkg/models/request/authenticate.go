package request

type Authenticate struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}
