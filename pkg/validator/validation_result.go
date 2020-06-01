package validator

type Result struct {
	Errors []Error `json:"errors"`
}

func (r Result) Invalid() bool {
	return len(r.Errors) > 0
}
