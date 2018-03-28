package webserver

import (
	"net/http"
	"net/http/httptest"

	"github.com/sanksons/reflorest/src/core/service"
)

type TestWebserver struct {
}

func (ws *TestWebserver) Response(req *http.Request) *httptest.ResponseRecorder {

	w := httptest.NewRecorder()
	webServer := new(service.Webserver)
	webServer.ServiceHandler(w, req)

	return w
}
