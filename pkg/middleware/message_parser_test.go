package middleware_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/antonio-muniz/alph/test/helpers"
	"github.com/stretchr/testify/require"

	"github.com/antonio-muniz/alph/pkg/middleware"
)

func TestMessageParser(t *testing.T) {
	scenarios := []struct {
		description          string
		requestBody          []byte
		messageType          middleware.MessageType
		handler              http.Handler
		expectedStatusCode   int
		expectedResponseBody map[string]interface{}
	}{
		{
			description:          "provides_parsed_message_in_request_context",
			requestBody:          serializeJSON(map[string]interface{}{"number": 123}),
			messageType:          incrementMessage{},
			handler:              incrementHandler{},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: map[string]interface{}{"message": "the number is 123"},
		},
		{
			description:          "rejects_a_malformed_request",
			requestBody:          []byte("Trust me, this is JSON!"),
			messageType:          incrementMessage{},
			handler:              incrementHandler{},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: map[string]interface{}{"message": "malformed request"},
		},
		{
			description:        "rejects_a_request_with_invalid_parameters",
			requestBody:        serializeJSON(map[string]interface{}{"number": -97}),
			messageType:        incrementMessage{},
			handler:            incrementHandler{},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponseBody: map[string]interface{}{
				"validation_errors": []interface{}{
					map[string]interface{}{
						"type":  "TOO_LOW",
						"field": "number",
						"value": float64(-97),
						"details": map[string]interface{}{
							"minimum": "0",
						},
					},
				},
			},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			middleware := middleware.MessageParser(
				scenario.messageType,
				scenario.handler,
			)
			request := helpers.BuildHttpRequest(t, http.MethodPost, "/", scenario.requestBody)
			response := httptest.NewRecorder()
			middleware.ServeHTTP(response, request)
			require.Equal(t, scenario.expectedStatusCode, response.Code)
			responseBody := helpers.DeserializeHttpResponseBody(t, response)
			require.Equal(t, scenario.expectedResponseBody, responseBody)
		})
	}
}

type incrementMessage struct {
	Number int `json:"number" validate:"gte=0"`
}

func (m incrementMessage) NewPointer() interface{} {
	return &incrementMessage{}
}

func (m incrementMessage) Dereference(pointer interface{}) interface{} {
	return *(pointer.(*incrementMessage))
}

type incrementHandler struct{}

func (h incrementHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	message := ctx.Value(middleware.Message).(incrementMessage)
	responseBody := []byte(fmt.Sprintf("{\"message\": \"the number is %d\"}", message.Number))
	response.Write(responseBody)
}

func serializeJSON(data interface{}) []byte {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return bytes
}
