package handles

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"io"
	"../common"
)

// Writes the response as a standard JSON response with StatusOK
func ResponseSuccess(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&common.JsonResponse{Data: m}); err != nil {
		ResponseError(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

// Writes the error response as a Standard API JSON response with a response code
func ResponseError(w http.ResponseWriter, errorCode int, errorMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorCode)
	json.NewEncoder(w).Encode(&common.JsonErrorResponse{Error: &common.ApiError{Status: errorCode, Title: errorMsg}})
}

func FromHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params, model interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}
	if err := r.Body.Close(); err != nil {
		return err
	}
	if err := json.Unmarshal(body, model); err != nil {
		return err
	}
	return nil
}
