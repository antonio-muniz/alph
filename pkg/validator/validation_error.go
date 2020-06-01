package validator

import "fmt"

type Error struct {
	Type    string                 `json:"type"`
	Field   string                 `json:"field"`
	Value   interface{}            `json:"value,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("'%s'='%v': %s", e.Field, e.Value, e.Type)
}
