package internalhelpers

import (
	nethttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/antonio-muniz/alph/cmd/alph/internal/transport/http"
	"github.com/antonio-muniz/alph/pkg/system"
)

func ExecuteHttpRequest(
	t *testing.T,
	sys system.System,
	request *nethttp.Request,
) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	router := http.Router(sys)
	router.ServeHTTP(response, request)
	return response
}
