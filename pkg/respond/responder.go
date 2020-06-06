package respond

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func OK(response http.ResponseWriter, data interface{}) {
	sendJSONResponse(response, http.StatusOK, data)
}

func Created(response http.ResponseWriter, resourceID string, resource interface{}) {
	relativeLocation := fmt.Sprintf("/%s", resourceID)
	response.Header().Set("Location", relativeLocation)
	sendJSONResponse(response, http.StatusCreated, resource)
}

func MalformedRequest(response http.ResponseWriter) {
	body := []byte("{\"message\":\"malformed request\"}")
	sendResponse(response, http.StatusBadRequest, body)
}

func InvalidRequestParameters(response http.ResponseWriter, details interface{}) {
	sendJSONResponse(response, http.StatusBadRequest, details)
}

func Forbidden(response http.ResponseWriter) {
	body := []byte("{\"message\":\"forbidden\"}")
	sendResponse(response, http.StatusForbidden, body)
}

func UnsupportedContentType(response http.ResponseWriter) {
	body := []byte("{\"message\":\"unsupported content type\"}")
	sendResponse(response, http.StatusUnsupportedMediaType, body)
}

func InternalServerError(response http.ResponseWriter) {
	body := []byte("{\"message\":\"unexpected error\"}")
	sendResponse(response, http.StatusInternalServerError, body)
}

func sendJSONResponse(response http.ResponseWriter, statusCode int, body interface{}) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		InternalServerError(response)
		return
	}
	sendResponse(response, statusCode, bodyBytes)
}

func sendResponse(response http.ResponseWriter, statusCode int, body []byte) {
	response.WriteHeader(statusCode)
	_, err := response.Write(body)
	if err != nil {
		panic(err)
	}
}
