package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http/middleware"
	fixtures "github.com/antonio-muniz/alph/test/fixtures/http"
	"github.com/stretchr/testify/require"
)

func TestPanicRecovery(t *testing.T) {
	scenarios := []struct {
		description        string
		handler            http.Handler
		expectedStatusCode int
	}{
		{
			description:        "does_nothing_when_handler_returns_normally",
			handler:            fixtures.NewNoContentHandler(),
			expectedStatusCode: http.StatusNoContent,
		},
		{
			description:        "returns_internal_server_error_when_handler_panics",
			handler:            fixtures.NewPanicHandler(),
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			middleware := middleware.PanicRecovery(scenario.handler)
			request, err := http.NewRequest(http.MethodPost, "/fake", nil)
			require.NoError(t, err)
			response := httptest.NewRecorder()
			require.NotPanics(t, func() { middleware.ServeHTTP(response, request) })
			require.Equal(t, scenario.expectedStatusCode, response.Code)
		})
	}
}
