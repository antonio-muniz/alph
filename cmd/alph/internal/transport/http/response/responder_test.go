package response_test

import (
	"net/http/httptest"
	"testing"

	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/response"
	"github.com/antonio-muniz/alph/test/helpers"
	"github.com/stretchr/testify/require"
)

func TestCreated(t *testing.T) {
	httpResponse := httptest.NewRecorder()
	resource := map[string]interface{}{"field": "value"}

	response.Created(httpResponse, "resource-id", resource)

	responseBody := helpers.DeserializeHttpResponseBody(t, httpResponse)
	require.Equal(t, resource, responseBody)
	require.Equal(t, "/resource-id", httpResponse.Header().Get("Location"))
}

func TestMalformedRequestBody(t *testing.T) {}

func TestInvalidRequestParameters(t *testing.T) {}

func TestForbidden(t *testing.T) {}

func TestInternalServerError(t *testing.T) {}
