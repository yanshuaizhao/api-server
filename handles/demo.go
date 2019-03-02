package handles

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"../models"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "hi, welcome!\n")
}

func DemoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	lists := []*models.Demo{}
	for _, item := range models.Demostore {
		lists = append(lists, item)
	}
	ResponseSuccess(w, lists)
}

func DemoShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, _ := strconv.Atoi(params.ByName("id"))
	demo, ok := models.Demostore[id]
	if !ok {
		ResponseError(w, http.StatusNotFound, "Record Not Found")
		return
	}
	ResponseSuccess(w, demo)
}

func DemoCreate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	item := &models.Demo{}
	if err := FromHandler(w, r, params, item); err != nil {
		ResponseError(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	models.Demostore[item.Id] = item
	ResponseSuccess(w, item)
}
