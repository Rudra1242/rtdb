package orderitemhandler

import (
	"net/http"
	"rtdb/adapter/OrderItem"
	bshandler "rtdb/applications/restapplication/packages/httphandlers"
	"rtdb/applications/restapplication/packages/mthdroutr"
	"rtdb/applications/restapplication/packages/resputl"
	"rtdb/engine/orderitem"
	"rtdb/utils/loggerutils"
)

// OrderItemHandler ...
type OrderItemHandler struct {
	bshandler.BaseHandler
	OrderItemService OrderItem.Service
}

func (o *OrderItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := mthdroutr.RouteApiCall(o, r)
	response.RenderResponse(w)
}

// Get function for PingHandler
func (o *OrderItemHandler) Get(r *http.Request) resputl.SrvcRes {
	logger := loggerutils.GetLogger()
	logger.Infof("Got OrderItem count Request")
	orderitemAccess := orderitem.OrderItemAccess{
		OrderItemService: o.OrderItemService,
	}
	count := orderitemAccess.GetAllOrderCount()
	var response OrderItemResponse
	response.Count = count
	return resputl.Response200OK(response)
}
