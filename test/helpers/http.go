package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func BuildHttpRequest(
	t *testing.T,
	method string,
	uri string,
	body interface{},
) *http.Request {
	var bodyBytes []byte
	switch typedBody := body.(type) {
	case []byte:
		bodyBytes = typedBody
	default:
		serializedBody, err := json.Marshal(body)
		require.NoError(t, err)
		bodyBytes = serializedBody
	}
	bodyReader := bytes.NewReader(bodyBytes)
	request, err := http.NewRequest(method, uri, bodyReader)
	require.NoError(t, err)
	request.Header.Set("Content-Type", "application/json")
	return request
}

func DeserializeHttpResponseBody(
	t *testing.T,
	response *httptest.ResponseRecorder,
) map[string]interface{} {
	var responseBody map[string]interface{}
	err := json.NewDecoder(response.Body).Decode(&responseBody)
	require.NoError(t, err)
	return responseBody
}
