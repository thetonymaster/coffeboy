package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/coopernurse/gorp"
	"github.com/crowdint/coffeboy/controllers/categoriescontroller"
	"github.com/crowdint/coffeboy/controllers/currenttime"
	"github.com/crowdint/coffeboy/controllers/orderscontroller"
	"github.com/crowdint/coffeboy/controllers/productscontroller"
	"github.com/crowdint/coffeboy/controllers/rolescontroller"
	"github.com/crowdint/coffeboy/controllers/userscontroller"
	"github.com/crowdint/coffeboy/models/categories"
	"github.com/crowdint/coffeboy/models/orders"
	"github.com/crowdint/coffeboy/models/products"
	"github.com/crowdint/coffeboy/models/roles"
	"github.com/crowdint/coffeboy/models/users"
	"github.com/gorilla/mux"
	"github.com/jingweno/negroni-gorelic"
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
	n := negroni.New()
	n.Use(negronigorelic.New("dc91b004a0e39086cb75ea6a4442cc1a8509d6f7", "coffeboy", true))
	n.UseHandler(r)
	n.Run(":" + os.Getenv("PORT"))
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
	r.HandleFunc("/product", f(productscontroller.GetAll)).Methods("GET")
	r.HandleFunc("/product/{id}", f(productscontroller.Get)).Methods("GET")
	r.HandleFunc("/product/{id}", f(productscontroller.Update)).Methods("PUT")
	r.HandleFunc("/product/{id}", f(productscontroller.Delete)).Methods("DELETE")

	r.HandleFunc("/category", f(categoriescontroller.Save)).Methods("POST")
	r.HandleFunc("/category", f(categoriescontroller.GetAll)).Methods("GET")
	r.HandleFunc("/category/{id}", f(categoriescontroller.Get)).Methods("GET")
	r.HandleFunc("/category/{id}", f(categoriescontroller.Update)).Methods("PUT")
	r.HandleFunc("/category/{id}", f(categoriescontroller.Delete)).Methods("DELETE")

	r.HandleFunc("/role", f(rolescontroller.Save)).Methods("POST")
	r.HandleFunc("/role/{id}", f(rolescontroller.Get)).Methods("GET")
	r.HandleFunc("/role/{id}", f(rolescontroller.Update)).Methods("PUT")
	r.HandleFunc("/role/{id}", f(rolescontroller.Delete)).Methods("DELETE")

	r.HandleFunc("/user", f(userscontroller.Save)).Methods("POST")
	r.HandleFunc("/user/{id}", f(userscontroller.Get)).Methods("GET")
	r.HandleFunc("/user/{id}", f(userscontroller.Update)).Methods("PUT")
	r.HandleFunc("/user/{id}", f(userscontroller.Delete)).Methods("DELETE")

	r.HandleFunc("/order", f(orderscontroller.Save)).Methods("POST")
	r.HandleFunc("/order/{id}", f(orderscontroller.Get)).Methods("GET")
	r.HandleFunc("/order/{id}", f(orderscontroller.Update)).Methods("PUT")
	r.HandleFunc("/order/{id}", f(orderscontroller.Delete)).Methods("DELETE")

	r.HandleFunc("/current_time", currenttime.Get).Methods("GET")
	r.HandleFunc("/products", f(categoriescontroller.GetAll)).Methods("GET")

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
	dbmap.AddTableWithName(roles.Role{}, "roles").SetKeys(true, "id")
	dbmap.AddTableWithName(users.User{}, "users").SetKeys(true, "id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		return nil, err
	}

	return dbmap, nil
}
