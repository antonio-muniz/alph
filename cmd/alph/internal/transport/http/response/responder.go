package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Created(httpResponse http.ResponseWriter, resourceID string, resource interface{}) {
	body, err := json.Marshal(resource)
	if err != nil {
		httpResponse.WriteHeader(http.StatusInternalServerError)
		_, err := httpResponse.Write([]byte("{}"))
		if err != nil {
			panic(err)
		}
	}
	relativeLocation := fmt.Sprintf("/%s", resourceID)
	httpResponse.Header().Set("Location", relativeLocation)
	httpResponse.WriteHeader(http.StatusCreated)
	httpResponse.Write(body)
}
