package rolescontroller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/coopernurse/gorp"
	"github.com/crowdint/coffeboy/models/roles"
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

	var role roles.Role
	err = json.Unmarshal(body, &role)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	err = role.Save(dbmap)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	body, err = json.Marshal(role)
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

	role, err := roles.Get(id, dbmap)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	body, err := json.Marshal(role)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)

}

func Update(dbmap *gorp.DbMap, w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var role roles.Role
	err = json.Unmarshal(body, &role)

	if err != nil {
		panic(err)
	}

	err = role.Update(dbmap)
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

	role, err := roles.Get(id, dbmap)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = role.Delete(dbmap)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
