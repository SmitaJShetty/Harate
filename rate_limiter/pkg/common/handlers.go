package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vayu/internal/service"
	"vayu/pkg/apperror"
	
)

//AppHandlers construct for handlers
type AppHandlers struct {
	svc *service.RandomService
}

//GetHandler handler for GetHandler
func (h *AppHandlers) GetHandler(w http.ResponseWriter, req *http.Request) {
	res, err := h.svc.Get()
	if err != nil {
		SendErrorResponse(w, req, apperror.NewAppError(fmt.Sprintf("error occurred:%v", err), http.StatusInternalServerError))
		return
	}

	resp, respErr := json.Marshal(res)
	if respErr != nil {
		SendErrorResponse(w, req, apperror.NewAppError(fmt.Sprintf("error occurred:%v", respErr), http.StatusInternalServerError))
		return 
	}

	SendResult(w, req, []byte(resp))
	return
}

//NewHandlers return new Handlers
func NewHandlers() *AppHandlers {
	return &AppHandlers{
		svc: service.NewRandomService(),
	}
}