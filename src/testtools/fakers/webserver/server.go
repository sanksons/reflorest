package webserver

import (
	"net/http"
	"net/http/httptest"

	"github.com/sanksons/reflorest/src/core/service"
)

type TestWebserver struct {
	Ws *service.Webserver
}

func Initialize(a func()) *TestWebserver {
	webServer := new(service.Webserver)
	service.SetAppMode(service.MODE_TEST)
	//service.Register()
	webServer.PreStart(a, func() {})
	ts := new(TestWebserver)
	ts.Ws = webServer
	return ts
}

func (this *TestWebserver) Response(req *http.Request) *httptest.ResponseRecorder {

	w := httptest.NewRecorder()

	this.Ws.ServiceHandler(w, req)

	return w
}
