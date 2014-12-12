package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/coopernurse/gorp"
	"github.com/crowdint/coffeboy/controllers/categoriescontroller"
	"github.com/crowdint/coffeboy/controllers/productscontroller"
	"github.com/crowdint/coffeboy/models/categories"
	"github.com/crowdint/coffeboy/models/products"
	"github.com/gorilla/mux"
)

var dbmap *gorp.DbMap
var dberr error

func init() {
	dbmap, dberr = InitDb()
	if dberr != nil {
		panic(dberr)
	}
}

func main() {
	r := CreateHandler(CreateDbMapHandlerToHTTPHandler(dbmap))
	http.Handle("/", r)
	log.Println("Starting server...")
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request)
type DbMapHandlerFunc func(dbmap *gorp.DbMap, w http.ResponseWriter, r *http.Request)
type DbMapHandlerToHTTPHandlerHOF func(f DbMapHandlerFunc) HandlerFunc

func CreateDbMapHandlerToHTTPHandler(dbmap *gorp.DbMap) DbMapHandlerToHTTPHandlerHOF {
	return func(f DbMapHandlerFunc) HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			f(dbmap, w, r)
		}
	}
}

func CreateHandler(f DbMapHandlerToHTTPHandlerHOF) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/product", f(productscontroller.Save)).Methods("POST")
	r.HandleFunc("/product/{id}", f(productscontroller.Get)).Methods("GET")
	r.HandleFunc("/product/{id}", f(productscontroller.Update)).Methods("PUT")
	r.HandleFunc("/product/{id}", f(productscontroller.Delete)).Methods("DELETE")

	r.HandleFunc("/category", f(categoriescontroller.Save)).Methods("POST")
	r.HandleFunc("/category/{id}", f(categoriescontroller.Get)).Methods("GET")
	r.HandleFunc("/category/{id}", f(categoriescontroller.Update)).Methods("PUT")
	r.HandleFunc("/category/{id}", f(categoriescontroller.Delete)).Methods("DELETE")

	return r
}

func InitDb() (*gorp.DbMap, error) {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		return nil, err
	}

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(products.Product{}, "products").SetKeys(true, "id")
	dbmap.AddTableWithName(categories.Category{}, "categories").SetKeys(true, "id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		return nil, err
	}

	return dbmap, nil
}