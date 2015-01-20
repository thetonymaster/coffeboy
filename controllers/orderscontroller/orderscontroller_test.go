package orderscontroller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/coopernurse/gorp"
	. "github.com/crowdint/coffeboy/controllers/orderscontroller"
	"github.com/crowdint/coffeboy/models/orders"
	"github.com/gorilla/mux"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Orderscontroller", func() {

	var (
		dbmap *gorp.DbMap
		r     *mux.Router
	)

	BeforeSuite(func() {
		dbmap, _ = orders.InitDb()
		r = CreateHandler(CreateDbMapHandlerToHTTPHandler(dbmap))
	})

	Describe("Create an order", func() {
		Context("Send a POST request", func() {

			It("Should return status created", func() {

				record := httptest.NewRecorder()

				message := `{"id":"R9999","user_id":1,` +
					`"created_at":""` +
					`,"updated_at":"","completed_at":"","email":""` +
					`,"total_quantity":"","line_items":` +
					`[{"variant_id":"1","quantity":10},{"variant_id":"2","quantity":20}]}`

				req, err := http.NewRequest("POST", "/order", bytes.NewReader([]byte(message)))
				r.ServeHTTP(record, req)

				Expect(err).To(BeNil())
				Expect(record.Code).To(Equal(http.StatusCreated))

			})
		})
	})

	Describe("Get an order from the database", func() {
		Context("Send a GET request", func() {
			It("Should return an order JSON object", func() {
				record := httptest.NewRecorder()

				message := `{"id":"R9999","user_id":1,` +
					`"created_at":""` +
					`,"updated_at":"","completed_at":"","email":""` +
					`,"total_quantity":"","line_items":` +
					`[{"variant_id":"1","quantity":10},{"variant_id":"2","quantity":20}]}`

				req, err := http.NewRequest("POST", "/order", bytes.NewReader([]byte(message)))
				r.ServeHTTP(record, req)

				Expect(err).To(BeNil())
				Expect(record.Code).To(Equal(http.StatusCreated))

				var order orders.Order

				err = json.Unmarshal(record.Body.Bytes(), &order)
				Expect(err).To(BeNil())

				url := fmt.Sprintf("/order/%s", order.ID)

				record2 := httptest.NewRecorder()
				req, err = http.NewRequest("GET", url, nil)
				r.ServeHTTP(record2, req)

				Expect(record2.Code).To(Equal(http.StatusOK))

			})
		})
	})

	Describe("Update an order from the database", func() {
		Context("Send a PUT request", func() {
			It("Should return an order JSON object", func() {
				createRecord := httptest.NewRecorder()

				message := `{"id":"R9999","user_id":1,` +
					`"created_at":""` +
					`,"updated_at":"","completed_at":"","email":""` +
					`,"total_quantity":"","line_items":` +
					`[{"variant_id":"1","quantity":10},{"variant_id":"2","quantity":20}]}`

				createRequest, err := http.NewRequest("POST", "/order", bytes.NewReader([]byte(message)))
				r.ServeHTTP(createRecord, createRequest)

				Expect(err).To(BeNil())
				Expect(createRecord.Code).To(Equal(http.StatusCreated))

				var order orders.Order

				err = json.Unmarshal(createRecord.Body.Bytes(), &order)

				url := fmt.Sprintf("/order/%s", order.ID)

				order.Quantity = "100"

				body, err := json.Marshal(order)
				Expect(err).To(BeNil())

				updateRequest, err := http.NewRequest("PUT", url, bytes.NewReader([]byte(body)))
				updateRecord := httptest.NewRecorder()
				r.ServeHTTP(updateRecord, updateRequest)

				Expect(updateRecord.Code).To(Equal(http.StatusOK))

				getUpdatedRequest, err := http.NewRequest("GET", url, nil)
				getUpdateRecord := httptest.NewRecorder()
				r.ServeHTTP(getUpdateRecord, getUpdatedRequest)

				Expect(getUpdateRecord.Code).To(Equal(http.StatusOK))

				var updatedOrder orders.Order
				err = json.Unmarshal(getUpdateRecord.Body.Bytes(), &updatedOrder)

				fmt.Println(updatedOrder)

				Expect(updatedOrder.Quantity).To(Equal("100"))

			})
		})
	})

	Describe("Delete an order from the database", func() {
		Context("Send a DELETE request", func() {
			It("Should delete the object", func() {
				record := httptest.NewRecorder()

				message := `{"id":"R9999","user_id":1,` +
					`"created_at":""` +
					`,"updated_at":"","completed_at":"","email":""` +
					`,"total_quantity":"","line_items":` +
					`[{"variant_id":"1","quantity":10},{"variant_id":"2","quantity":20}]}`

				req, err := http.NewRequest("POST", "/order", bytes.NewReader([]byte(message)))
				r.ServeHTTP(record, req)

				Expect(err).To(BeNil())
				Expect(record.Code).To(Equal(http.StatusCreated))

				var order orders.Order

				err = json.Unmarshal(record.Body.Bytes(), &order)

				url := fmt.Sprintf("/order/%s", order.ID)

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
	r.HandleFunc("/order", f(Save)).Methods("POST")
	r.HandleFunc("/order/{id}", f(Get)).Methods("GET")
	r.HandleFunc("/order/{id}", f(Update)).Methods("PUT")
	r.HandleFunc("/order/{id}", f(Delete)).Methods("DELETE")

	return r
}
