package categoriescontroller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/coopernurse/gorp"
	. "github.com/crowdint/coffeboy/controllers/categoriescontroller"
	"github.com/crowdint/coffeboy/models/categories"
	"github.com/gorilla/mux"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Categoriescontroller", func() {

	var (
		dbmap *gorp.DbMap
		r     *mux.Router
	)

	BeforeSuite(func() {
		dbmap, _ = categories.InitDb()
		r = CreateHandler(CreateDbMapHandlerToHTTPHandler(dbmap))
	})

	Describe("Create a category", func() {
		Context("Send a POST request", func() {

			It("Should return status created", func() {

				record := httptest.NewRecorder()
				message := `{"name":"breakfast"}`
				req, err := http.NewRequest("POST", "/category", bytes.NewReader([]byte(message)))
				r.ServeHTTP(record, req)

				Expect(err).To(BeNil())
				Expect(record.Code).To(Equal(http.StatusCreated))

			})
		})
	})

	Describe("Get a category from the database", func() {
		Context("Send a GET request", func() {
			It("Should return a category JSON object", func() {
				record := httptest.NewRecorder()
				message := `{"name":"breakfast"}`
				req, err := http.NewRequest("POST", "/category", bytes.NewReader([]byte(message)))
				r.ServeHTTP(record, req)

				Expect(err).To(BeNil())
				Expect(record.Code).To(Equal(http.StatusCreated))

				var category categories.Category

				err = json.Unmarshal(record.Body.Bytes(), &category)

				url := fmt.Sprintf("/category/%d", category.ID)

				record2 := httptest.NewRecorder()
				req, err = http.NewRequest("GET", url, nil)
				r.ServeHTTP(record2, req)

				Expect(record2.Code).To(Equal(http.StatusOK))

			})
		})
	})

	Describe("Update a category from the database", func() {
		Context("Send a PUT request", func() {
			It("Should return a category JSON object", func() {
				record := httptest.NewRecorder()
				message := `{"name":"breakfast"}`
				req, err := http.NewRequest("POST", "/category", bytes.NewReader([]byte(message)))
				r.ServeHTTP(record, req)

				Expect(err).To(BeNil())
				Expect(record.Code).To(Equal(http.StatusCreated))

				var category categories.Category

				err = json.Unmarshal(record.Body.Bytes(), &category)

				url := fmt.Sprintf("/category/%d", category.ID)

				category.Name = "breakfast at tiffany's"

				body, err := json.Marshal(category)

				req2, err := http.NewRequest("PUT", url, bytes.NewReader([]byte(body)))
				record2 := httptest.NewRecorder()
				r.ServeHTTP(record2, req2)

				Expect(record2.Code).To(Equal(http.StatusOK))

				req3, err := http.NewRequest("GET", url, nil)
				record3 := httptest.NewRecorder()
				r.ServeHTTP(record3, req3)

				var category2 categories.Category
				err = json.Unmarshal(record3.Body.Bytes(), &category2)

				Expect(category).To(Equal(category2))

			})
		})
	})

	Describe("Delete a category from the database", func() {
		Context("Send a DELETE request", func() {
			It("Should delete the object", func() {
				record := httptest.NewRecorder()
				message := `{"name":"breakfast"}`
				req, err := http.NewRequest("POST", "/category", bytes.NewReader([]byte(message)))
				r.ServeHTTP(record, req)

				Expect(err).To(BeNil())
				Expect(record.Code).To(Equal(http.StatusCreated))

				var category categories.Category

				err = json.Unmarshal(record.Body.Bytes(), &category)

				url := fmt.Sprintf("/category/%d", category.ID)

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
	r.HandleFunc("/category", f(Save)).Methods("POST")
	r.HandleFunc("/category/{id}", f(Get)).Methods("GET")
	r.HandleFunc("/category/{id}", f(Update)).Methods("PUT")
	r.HandleFunc("/category/{id}", f(Delete)).Methods("DELETE")

	return r
}
