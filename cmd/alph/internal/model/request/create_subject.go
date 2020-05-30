package request

type CreateSubject struct {
	SubjectID string `json:"subject_id"`
	Password  string `json:"password"`
}
