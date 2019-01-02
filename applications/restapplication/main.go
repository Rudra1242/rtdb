package main

import (
	"net/http"
	"os"
	"rtdb/adapter/OrderItem"
	"rtdb/applications/restapplication/orderitemhandler"
	"rtdb/utils/loggerutils"
	"rtdb/utils/mongoutils"
	"time"

	"github.com/gorilla/mux"
)

func init() {
	/*
	   Safety net for 'too many open files' issue on legacy code.
	   Set a sane timeout duration for the http.DefaultClient, to ensure idle connections are terminated.
	   Reference: https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	   https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	*/
	http.DefaultClient.Timeout = time.Minute * 10
}

func main() {
	mongoSession, _ := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"), "notauthenticated", nil)
	mongoRepo := OrderItem.NewMongoRepository(mongoSession, "agrostar_csr_dashboard")
	OrderItemService := OrderItem.NewService(mongoRepo)

	orderitemhandler := &orderitemhandler.OrderItemHandler{
		OrderItemService: *OrderItemService,
	}
	logger := loggerutils.GetLogger()
	logger.Infof("Setting up resources.")
	h := mux.NewRouter()

	h.Handle("/rtdb/orderitem/", orderitemhandler)

	logger.Infof("Resource Setup Done.")
	logger.Error(http.ListenAndServe(":6677", h))
}
