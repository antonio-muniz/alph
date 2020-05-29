package request

type Authenticate struct {
	SubjectID string `json:"subject_id"`
	Password  string `json:"password"`
}
