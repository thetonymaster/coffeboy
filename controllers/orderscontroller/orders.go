package orderscontroller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/coopernurse/gorp"
	"github.com/crowdint/coffeboy/models/orders"
	"github.com/gorilla/mux"
)

func Save(dbmap *gorp.DbMap, w http.ResponseWriter, r *http.Request) {
	order, err := getOrderFromRequest(r)
	if err != nil {
		writeInternalError(err, w)
		return
	}

	err = order.Save(dbmap)
	if err != nil {
		writeInternalError(err, w)
		return
	}

	body, err := order.Marshal()
	if err != nil {
		writeInternalError(err, w)
		return
	}

	writeOkResponse(w, http.StatusCreated, body)
}

func Get(dbmap *gorp.DbMap, w http.ResponseWriter, r *http.Request) {
	order, err := getOrderFromDB(dbmap, r)
	if err != nil {
		writeNotFound(err, w)
		return
	}

	body, err := order.Marshal()
	if err != nil {
		writeInternalError(err, w)
		return
	}

	writeOkResponse(w, http.StatusOK, body)
}

func Update(dbmap *gorp.DbMap, w http.ResponseWriter, r *http.Request) {
	order, err := getOrderFromRequest(r)
	if err != nil {
		writeInternalError(err, w)
		return
	}

	err = order.Update(dbmap)
	if err != nil {
		writeInternalError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Delete(dbmap *gorp.DbMap, w http.ResponseWriter, r *http.Request) {
	order, err := getOrderFromDB(dbmap, r)
	if err != nil {
		writeNotFound(err, w)
		return
	}

	err = order.Delete(dbmap)
	if err != nil {
		writeInternalError(err, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func getOrderFromRequest(r *http.Request) (*orders.Order, error) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var order orders.Order
	err = json.Unmarshal(body, &order)

	return &order, err
}

func getOrderFromDB(dbmap *gorp.DbMap, r *http.Request) (*orders.Order, error) {
	params := mux.Vars(r)

	var id string

	if tempId, exists := params["id"]; !exists {
		return nil, errors.New("Id param missing")
	} else {
		id = tempId
	}

	return orders.GetOrder(id, dbmap)
}

func writeInternalError(err error, w http.ResponseWriter) {
	writeError(err, w, http.StatusInternalServerError)
}

func writeNotFound(err error, w http.ResponseWriter) {
	writeError(err, w, http.StatusNotFound)
}

func writeError(err error, w http.ResponseWriter, httpStatus int) {
	fmt.Printf("ERROR: %s\n", err.Error())
	w.WriteHeader(httpStatus)
}

func writeOkResponse(w http.ResponseWriter, httpStatus int, body []byte) {
	w.WriteHeader(httpStatus)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
