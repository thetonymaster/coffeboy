package productscontroller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/coopernurse/gorp"
	"github.com/crowdint/coffeboy/models/products"
	"github.com/crowdint/coffeboy/utils"
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

	var product products.Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	err = product.Save(dbmap)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	if product.Image != nil {
		identifier := strconv.FormatInt(product.ID, 10)
		go utils.UploadImages(product.Image, identifier, "products")
		product.Image = nil
		product.ImageURL = "https://s3-us-west-2.amazonaws.com/coffeboy/products/" + identifier + ".jpg"

		images := map[string]string{
			"small":    "https://s3-us-west-2.amazonaws.com/coffeboy/products/" + identifier + "_small.jpg",
			"medium":   "https://s3-us-west-2.amazonaws.com/coffeboy/products/" + identifier + "_medium.jpg",
			"large":    "https://s3-us-west-2.amazonaws.com/coffeboy/products/" + identifier + "_large.jpg",
			"original": "https://s3-us-west-2.amazonaws.com/coffeboy/products/" + identifier + ".jpg",
		}
		product.Images = images

	}

	body, err = json.Marshal(product)
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

	product, err := products.Get(id, dbmap)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if product.ImageURL != "" {
		identifier := strconv.FormatInt(product.ID, 10)
		images := map[string]string{
			"small":    "https://s3-us-west-2.amazonaws.com/coffeboy/products/" + identifier + "_small.jpg",
			"medium":   "https://s3-us-west-2.amazonaws.com/coffeboy/products/" + identifier + "_medium.jpg",
			"large":    "https://s3-us-west-2.amazonaws.com/coffeboy/products/" + identifier + "_large.jpg",
			"original": "https://s3-us-west-2.amazonaws.com/coffeboy/products/" + identifier + ".jpg",
		}
		product.Images = images
	}

	body, err := json.Marshal(product)
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

	var product products.Product
	err = json.Unmarshal(body, &product)

	if err != nil {
		panic(err)
	}

	err = product.Update(dbmap)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if product.Image != nil {

		identifier := strconv.FormatInt(product.ID, 10)
		go utils.UploadImages(product.Image, identifier, "products")
		product.Image = nil
		product.ImageURL = "https://s3-us-west-2.amazonaws.com/coffeboy/products/" + identifier + ".jpg"

	}

	if product.ImageURL != "" {
		identifier := strconv.FormatInt(product.ID, 10)
		images := map[string]string{
			"small":    "https://s3-us-west-2.amazonaws.com/coffeboy/products/" + identifier + "_small.jpg",
			"medium":   "https://s3-us-west-2.amazonaws.com/coffeboy/products/" + identifier + "_medium.jpg",
			"large":    "https://s3-us-west-2.amazonaws.com/coffeboy/products/" + identifier + "_large.jpg",
			"original": "https://s3-us-west-2.amazonaws.com/coffeboy/products/" + identifier + ".jpg",
		}
		product.Images = images
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

	product, err := products.Get(id, dbmap)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = product.Delete(dbmap)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
