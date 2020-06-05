package respond_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/antonio-muniz/alph/pkg/validator"

	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/respond"
	"github.com/antonio-muniz/alph/test/helpers"
	"github.com/stretchr/testify/require"
)

func TestOK(t *testing.T) {
	httpResponse := httptest.NewRecorder()
	data := map[string]interface{}{"field": "value"}

	respond.OK(httpResponse, data)

	require.Equal(t, http.StatusOK, httpResponse.Code)
	responseBody := helpers.DeserializeHttpResponseBody(t, httpResponse)
	require.Equal(t, data, responseBody)
}

func TestCreated(t *testing.T) {
	httpResponse := httptest.NewRecorder()
	resource := map[string]interface{}{"field": "value"}

	respond.Created(httpResponse, "resource-id", resource)

	require.Equal(t, http.StatusCreated, httpResponse.Code)
	responseBody := helpers.DeserializeHttpResponseBody(t, httpResponse)
	require.Equal(t, resource, responseBody)
	require.Equal(t, "/resource-id", httpResponse.Header().Get("Location"))
}

func TestMalformedRequest(t *testing.T) {
	httpResponse := httptest.NewRecorder()

	respond.MalformedRequest(httpResponse)

	require.Equal(t, http.StatusBadRequest, httpResponse.Code)
	responseBody := helpers.DeserializeHttpResponseBody(t, httpResponse)
	require.Equal(t, map[string]interface{}{"message": "malformed request"}, responseBody)
}

func TestInvalidRequestParameters(t *testing.T) {
	httpResponse := httptest.NewRecorder()
	validationResult := validator.Result{
		Errors: []validator.Error{
			{
				Type:  "MISSING",
				Field: "required_field",
			},
		},
	}

	respond.InvalidRequestParameters(httpResponse, validationResult)

	require.Equal(t, http.StatusBadRequest, httpResponse.Code)
	responseBody := helpers.DeserializeHttpResponseBody(t, httpResponse)
	require.Equal(t, map[string]interface{}{
		"validation_errors": []interface{}{
			map[string]interface{}{
				"type":  "MISSING",
				"field": "required_field",
			},
		},
	}, responseBody)
}

func TestForbidden(t *testing.T) {
	httpResponse := httptest.NewRecorder()

	respond.Forbidden(httpResponse)

	require.Equal(t, http.StatusForbidden, httpResponse.Code)
	responseBody := helpers.DeserializeHttpResponseBody(t, httpResponse)
	require.Equal(t, map[string]interface{}{"message": "forbidden"}, responseBody)
}

func TestInternalServerError(t *testing.T) {
	httpResponse := httptest.NewRecorder()

	respond.InternalServerError(httpResponse)

	require.Equal(t, http.StatusInternalServerError, httpResponse.Code)
	responseBody := helpers.DeserializeHttpResponseBody(t, httpResponse)
	require.Equal(t, map[string]interface{}{"message": "unexpected error"}, responseBody)
}
