package srvchandlers

import (
	"net/http"

	"rtdb/applications/restapplication/packages/mthdroutr"
	"rtdb/applications/restapplication/packages/resputl"
	"rtdb/utils/loggerutils"
)

// PingHandler is a Basic ping utility for the service
type PingHandler struct {
	BaseHandler
}

func (p *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := mthdroutr.RouteApiCall(p, r)
	response.RenderResponse(w)
}

// Get function for PingHandler
func (p *PingHandler) Get(r *http.Request) resputl.SrvcRes {
	logger := loggerutils.GetLogger()
	logger.Infof("Got Ping Request")
	return resputl.Response200OK("OK")
}
