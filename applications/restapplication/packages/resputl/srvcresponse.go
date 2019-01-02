package resputl

import (
	"encoding/json"
	"fmt"
	"net/http"

	customerrors "rtdb/applications/restapplication/packages/errors"
	"rtdb/utils/loggerutils"
	mail "rtdb/utils/mail"

	"gopkg.in/mgo.v2/bson"
)

type EmptyStruct struct {
}

type SrvcRes struct {
	Code     int
	Response interface{}
	Message  string
	Headers  map[string]string
}

func marshalResponse(r interface{}) ([]byte, error) {
	return json.MarshalIndent(r, "", "")
}

func (s *SrvcRes) RenderResponse(w http.ResponseWriter) {
	if s.Headers == nil {
		s.Headers = map[string]string{"Content-Type": "application/json",
			"Access-Control-Allow-Headers": "source,x-authorization-token,Content-Type",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "*"}
	}
	for h, val := range s.Headers {
		w.Header().Set(h, val)
	}
	var statusBool bool
	switch s.Code {
	case http.StatusOK:
		statusBool = true
	default:
		statusBool = false
	}

	formatted := bson.M{
		"responseData": s.Response,
		"message":      s.Message,
		"status":       statusBool}

	data, _ := marshalResponse(formatted)
	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	w.WriteHeader(s.Code)
	fmt.Fprint(w, string(data))
}

func Simple200OK(message string) SrvcRes {

	var inf EmptyStruct
	return SrvcRes{http.StatusOK, inf, message, nil}
}

//Simple404Response is given if a requested object is not found
func Simple404Response(message string) SrvcRes {

	var inf EmptyStruct
	return SrvcRes{http.StatusNotFound, inf, message, nil}
}

//Simple422Response is given if a requested object is not processable
func Simple422Response(message string) SrvcRes {

	var inf EmptyStruct
	return SrvcRes{http.StatusUnprocessableEntity, inf, message, nil}
}

//PreconditionFailed gives 412 response
func PreconditionFailed(message string) SrvcRes {
	var inf EmptyStruct
	return SrvcRes{http.StatusPreconditionFailed, inf, message, nil}
}

func OptionsResponseOK(message string) SrvcRes {

	var inf EmptyStruct
	return SrvcRes{http.StatusOK, inf, message, nil}
}

func SimpleBadRequest(message string) SrvcRes {
	return SrvcRes{http.StatusBadRequest, "{}", message, nil}
}

func Response200OK(response interface{}) SrvcRes {
	return SrvcRes{http.StatusOK, response, "OK", nil}
}

func ResponseNotImplemented(response interface{}) SrvcRes {
	return SrvcRes{http.StatusNotImplemented, "{}", "Method not implementd", nil}
}

func ReponseCustomError(err error) SrvcRes {
	//err := errors.
	//return SrvcRes{}
	//var cusErr *customerrors.CustomError
	var inf EmptyStruct
	cusErr := err.(*customerrors.CustomError)
	return SrvcRes{cusErr.GetStatusCode(), inf, cusErr.GetMessage(), nil}
}

func ReponseInternalError() SrvcRes {
	//err := errors.
	//return SrvcRes{}
	var inf EmptyStruct
	return SrvcRes{http.StatusInternalServerError, inf, "Internal Server Error", nil}
}

func sendErrorEmail(err *customerrors.CustomError, inf interface{}) {
	subject := fmt.Sprintf("OrderHold Error: %v", err.GetStatusCode())
	text := fmt.Sprintf("Order Hold Error: %v <br> Message: %v <br> Request: %s", err.GetStatusCode(), err.GetMessage(), inf)
	mail.SendEmail(subject, text)
}
func internalErrorEmail(err error, inf interface{}) {
	subject := fmt.Sprintf("OrderHold Internal Server Error")
	text := fmt.Sprintf("Order Hold Internal Error: %s <br>Request: %s", err, inf)
	mail.SendEmail(subject, text)
}

func ProcessError(err error, inf interface{}) SrvcRes {
	logger := loggerutils.GetLogger()
	logger.Errorf("Processing Error: %s", err)

	switch err.(type) {
	case *customerrors.CustomError:
		err1 := err.(*customerrors.CustomError)
		logger.Errorf("Error %v", err1)
		if true { //|| err1.GetStatusCode() != 404 {
			//go sendErrorEmail(err1, inf)
		}
		return ReponseCustomError(err)
	default:
		//go internalErrorEmail(err, inf)
		return ReponseInternalError()

	}
}
