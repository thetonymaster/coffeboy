package categoriescontroller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/coopernurse/gorp"
	"github.com/crowdint/coffeboy/models/categories"
	"github.com/gorilla/mux"
)

func Save(dbmap *gorp.DbMap, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	var category categories.Category
	err = json.Unmarshal(body, &category)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	err = category.Save(dbmap)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	body, err = json.Marshal(category)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

}

func Get(dbmap *gorp.DbMap, w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	category, err := categories.Get(id, dbmap)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	body, err := json.Marshal(category)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)

}

func GetAll(dbmap *gorp.DbMap, w http.ResponseWriter, r *http.Request) {
	cats, err := categories.GetAll(dbmap)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(cats)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Update(dbmap *gorp.DbMap, w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var category categories.Category
	err = json.Unmarshal(body, &category)

	if err != nil {
		panic(err)
	}

	err = category.Update(dbmap)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Delete(dbmap *gorp.DbMap, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	category, err := categories.Get(id, dbmap)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = category.Delete(dbmap)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
