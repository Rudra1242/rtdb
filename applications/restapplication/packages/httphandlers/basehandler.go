package srvchandlers

import (
	"net/http"

	"rtdb/applications/restapplication/packages/resputl"

	"rtdb/utils/loggerutils"

	"go.uber.org/zap"
)

//var logger *log.Logger
var logger *zap.SugaredLogger

func init() {
	// Package level global logger. Will be available in all the handlers.
	// And GetLogger is run only once when the package is imported.
	// But will be slightly confusing to understand where this logger is declared
	// when we check in other files.
	logger = loggerutils.GetLogger()
}

type BaseHandler struct{}

func (p *BaseHandler) GetOne(r *http.Request, id string) resputl.SrvcRes {
	return resputl.ResponseNotImplemented(nil)
}

func (p *BaseHandler) Get(r *http.Request) resputl.SrvcRes {
	return resputl.ResponseNotImplemented(nil)
}

func (p *BaseHandler) Put(r *http.Request) resputl.SrvcRes {
	return resputl.ResponseNotImplemented(nil)
}

func (p *BaseHandler) Post(r *http.Request) resputl.SrvcRes {
	return resputl.ResponseNotImplemented(nil)
}

func (p *BaseHandler) Delete(r *http.Request) resputl.SrvcRes {
	return resputl.ResponseNotImplemented(nil)
}

func (p *BaseHandler) Patch(r *http.Request) resputl.SrvcRes {
	return resputl.ResponseNotImplemented(nil)
}

func (p *BaseHandler) Options(r *http.Request) resputl.SrvcRes {
	return resputl.OptionsResponseOK("OK")
}
