package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/antonio-muniz/alph/pkg/middleware"
	fixtures "github.com/antonio-muniz/alph/test/fixtures/http"
	"github.com/stretchr/testify/require"
)

func TestContentNegotiation(t *testing.T) {
	scenarios := []struct {
		description                 string
		handler                     http.Handler
		requestContentType          string
		expectedStatusCode          int
		expectedResponseContentType string
	}{
		{
			description:                 "accepts_content_type_application/json",
			handler:                     fixtures.NewNoContentHandler(),
			requestContentType:          "application/json",
			expectedStatusCode:          http.StatusNoContent,
			expectedResponseContentType: "application/json",
		},
		{
			description:                 "rejects_other_content_types",
			handler:                     fixtures.NewNoContentHandler(),
			requestContentType:          "text/plain",
			expectedStatusCode:          http.StatusUnsupportedMediaType,
			expectedResponseContentType: "application/json",
		},
		{
			description:                 "rejects_empty_content_type",
			handler:                     fixtures.NewNoContentHandler(),
			requestContentType:          "",
			expectedStatusCode:          http.StatusUnsupportedMediaType,
			expectedResponseContentType: "application/json",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			middleware := middleware.ContentNegotiation(scenario.handler)
			request, err := http.NewRequest(http.MethodPost, "/fake", nil)
			require.NoError(t, err)
			request.Header.Set("Content-Type", scenario.requestContentType)
			response := httptest.NewRecorder()
			middleware.ServeHTTP(response, request)
			require.Equal(t, scenario.expectedStatusCode, response.Code)
			responseContentType := response.Header().Get("Content-Type")
			require.Equal(t, scenario.expectedResponseContentType, responseContentType)
		})
	}
}
