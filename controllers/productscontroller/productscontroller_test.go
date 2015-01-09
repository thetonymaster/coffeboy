package productscontroller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/coopernurse/gorp"
	. "github.com/crowdint/coffeboy/controllers/productscontroller"
	"github.com/crowdint/coffeboy/models/products"
	"github.com/gorilla/mux"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Productscontroller", func() {

	var (
		dbmap *gorp.DbMap
		r     *mux.Router
	)

	BeforeSuite(func() {
		dbmap, _ = products.InitDb()
		r = CreateHandler(CreateDbMapHandlerToHTTPHandler(dbmap))
	})

	Describe("Create a product", func() {
		Context("Send a POST request", func() {

			It("Should return status created", func() {

				record := httptest.NewRecorder()
				message := `{"category-id": 1, "name":"breakfast", "price": 1.00, "image-url": "http://foo/200x300", "stock":9999 }`
				req, err := http.NewRequest("POST", "/product", bytes.NewReader([]byte(message)))
				r.ServeHTTP(record, req)

				Expect(err).To(BeNil())
				Expect(record.Code).To(Equal(http.StatusCreated))

			})
		})
	})

	Describe("Get a product from the database", func() {
		Context("Send a GET request", func() {
			It("Should return a product JSON object", func() {
				record := httptest.NewRecorder()
				message := `{"category-id": 1, "name":"breakfast", "price": 1.00, "image-url": "http://foo/200x300", "stock":9999 }`
				req, err := http.NewRequest("POST", "/product", bytes.NewReader([]byte(message)))
				r.ServeHTTP(record, req)

				Expect(err).To(BeNil())
				Expect(record.Code).To(Equal(http.StatusCreated))

				var product products.Product

				err = json.Unmarshal(record.Body.Bytes(), &product)

				url := fmt.Sprintf("/product/%d", product.ID)

				record2 := httptest.NewRecorder()
				req, err = http.NewRequest("GET", url, nil)
				r.ServeHTTP(record2, req)

				Expect(record2.Code).To(Equal(http.StatusOK))

			})
		})
	})

	Describe("Upload a product with an image", func() {

	})

	Describe("Update a product from the database", func() {
		Context("Send a PUT request", func() {
			It("Should return a product JSON object", func() {
				record := httptest.NewRecorder()
				message := `{"category-id": 1, "name":"breakfast", "price": 1.00, "image-url": "http://foo/200x300", "stock":9999}`
				req, err := http.NewRequest("POST", "/product", bytes.NewReader([]byte(message)))
				r.ServeHTTP(record, req)

				Expect(err).To(BeNil())
				Expect(record.Code).To(Equal(http.StatusCreated))

				var product products.Product

				err = json.Unmarshal(record.Body.Bytes(), &product)

				url := fmt.Sprintf("/product/%d", product.ID)

				product.Name = "breakfast at tiffany's"

				body, err := json.Marshal(product)

				req2, err := http.NewRequest("PUT", url, bytes.NewReader([]byte(body)))
				record2 := httptest.NewRecorder()
				r.ServeHTTP(record2, req2)

				Expect(record2.Code).To(Equal(http.StatusOK))

				req3, err := http.NewRequest("GET", url, nil)
				record3 := httptest.NewRecorder()
				r.ServeHTTP(record3, req3)

				var product2 products.Product
				err = json.Unmarshal(record3.Body.Bytes(), &product2)

				Expect(product).To(Equal(product2))

			})
		})
	})

	Describe("Delete a product from the database", func() {
		Context("Send a DELETE request", func() {
			It("Should delete the object", func() {
				record := httptest.NewRecorder()
				message := `{"category-id": 1, "name":"breakfast", "price": 1.00, "image-url": "http://foo/200x300", "stock":9999}`
				req, err := http.NewRequest("POST", "/product", bytes.NewReader([]byte(message)))
				r.ServeHTTP(record, req)

				Expect(err).To(BeNil())
				Expect(record.Code).To(Equal(http.StatusCreated))

				var product products.Product

				err = json.Unmarshal(record.Body.Bytes(), &product)

				url := fmt.Sprintf("/product/%d", product.ID)

				record2 := httptest.NewRecorder()
				req, err = http.NewRequest("DELETE", url, nil)
				r.ServeHTTP(record2, req)

				Expect(record2.Code).To(Equal(http.StatusNoContent))

				record3 := httptest.NewRecorder()
				req, err = http.NewRequest("GET", url, nil)
				r.ServeHTTP(record3, req)
				Expect(record3.Code).To(Equal(http.StatusNotFound))

			})
		})
	})

	Describe("Get all products from the database", func() {
		Context("Send a GET request", func() {
			It("Should return saved products", func() {
				record := httptest.NewRecorder()
				message := `{"name":"breakfast"}`
				req, err := http.NewRequest("POST", "/product", bytes.NewReader([]byte(message)))
				r.ServeHTTP(record, req)

				Expect(err).To(BeNil())
				Expect(record.Code).To(Equal(http.StatusCreated))

				record2 := httptest.NewRecorder()
				message = `{"name":"dinner"}`
				req2, err := http.NewRequest("POST", "/product", bytes.NewReader([]byte(message)))
				r.ServeHTTP(record2, req2)

				Expect(err).To(BeNil())
				Expect(record2.Code).To(Equal(http.StatusCreated))

				record3 := httptest.NewRecorder()
				message = `{"name":"drink"}`
				req3, err := http.NewRequest("POST", "/product", bytes.NewReader([]byte(message)))
				r.ServeHTTP(record3, req3)

				Expect(err).To(BeNil())
				Expect(record3.Code).To(Equal(http.StatusCreated))

				record4 := httptest.NewRecorder()
				req4, err := http.NewRequest("GET", "/product", nil)
				r.ServeHTTP(record4, req4)

				Expect(err).To(BeNil())
				Expect(record4.Code).To(Equal(http.StatusOK))

				var products []products.Product

				err = json.Unmarshal(record4.Body.Bytes(), &products)
				Expect(err).To(BeNil())

				Expect(len(products)).Should(Equal(3))

			})
		})
	})

	AfterEach(func() {
		dbmap.TruncateTables()
	})

	AfterSuite(func() {
		dbmap.DropTables()
		dbmap.Db.Close()
	})

})

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
	r.HandleFunc("/product", f(Save)).Methods("POST")
	r.HandleFunc("/product", f(GetAll)).Methods("GET")
	r.HandleFunc("/product/{id}", f(Get)).Methods("GET")
	r.HandleFunc("/product/{id}", f(Update)).Methods("PUT")
	r.HandleFunc("/product/{id}", f(Delete)).Methods("DELETE")

	return r
}
