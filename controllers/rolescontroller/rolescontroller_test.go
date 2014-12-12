package rolescontroller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/coopernurse/gorp"
	. "github.com/crowdint/coffeboy/controllers/rolescontroller"
	"github.com/crowdint/coffeboy/models/roles"
	"github.com/gorilla/mux"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rolescontroller", func() {
	var (
		dbmap *gorp.DbMap
		r     *mux.Router
	)

	BeforeSuite(func() {
		dbmap, _ = roles.InitDb()
		r = CreateHandler(CreateDbMapHandlerToHTTPHandler(dbmap))
	})

	Describe("Create a role", func() {
		Context("Send a POST request", func() {

			It("Should return status created", func() {

				record := httptest.NewRecorder()
				message := `{"name":"user"}`
				req, err := http.NewRequest("POST", "/role", bytes.NewReader([]byte(message)))
				r.ServeHTTP(record, req)

				Expect(err).To(BeNil())
				Expect(record.Code).To(Equal(http.StatusCreated))

			})
		})
	})

	Describe("Get a role from the database", func() {
		Context("Send a GET request", func() {
			It("Should return a role JSON object", func() {
				record := httptest.NewRecorder()
				message := `{"name":"user"}`
				req, err := http.NewRequest("POST", "/role", bytes.NewReader([]byte(message)))
				r.ServeHTTP(record, req)

				Expect(err).To(BeNil())
				Expect(record.Code).To(Equal(http.StatusCreated))

				var role roles.Role

				err = json.Unmarshal(record.Body.Bytes(), &role)

				url := fmt.Sprintf("/role/%d", role.ID)

				record2 := httptest.NewRecorder()
				req, err = http.NewRequest("GET", url, nil)
				r.ServeHTTP(record2, req)

				Expect(record2.Code).To(Equal(http.StatusOK))

			})
		})
	})

	Describe("Update a role from the database", func() {
		Context("Send a PUT request", func() {
			It("Should return a role JSON object", func() {
				record := httptest.NewRecorder()
				message := `{"name":"user"}`
				req, err := http.NewRequest("POST", "/role", bytes.NewReader([]byte(message)))
				r.ServeHTTP(record, req)

				Expect(err).To(BeNil())
				Expect(record.Code).To(Equal(http.StatusCreated))

				var role roles.Role

				err = json.Unmarshal(record.Body.Bytes(), &role)

				url := fmt.Sprintf("/role/%d", role.ID)

				role.Name = "breakfast at tiffany's"

				body, err := json.Marshal(role)

				req2, err := http.NewRequest("PUT", url, bytes.NewReader([]byte(body)))
				record2 := httptest.NewRecorder()
				r.ServeHTTP(record2, req2)

				Expect(record2.Code).To(Equal(http.StatusOK))

				req3, err := http.NewRequest("GET", url, nil)
				record3 := httptest.NewRecorder()
				r.ServeHTTP(record3, req3)

				var product2 roles.Role
				err = json.Unmarshal(record3.Body.Bytes(), &product2)

				Expect(role).To(Equal(product2))

			})
		})
	})

	Describe("Delete a role from the database", func() {
		Context("Send a DELETE request", func() {
			It("Should delete the object", func() {
				record := httptest.NewRecorder()
				message := `{"name":"user"}`
				req, err := http.NewRequest("POST", "/role", bytes.NewReader([]byte(message)))
				r.ServeHTTP(record, req)

				Expect(err).To(BeNil())
				Expect(record.Code).To(Equal(http.StatusCreated))

				var role roles.Role

				err = json.Unmarshal(record.Body.Bytes(), &role)

				url := fmt.Sprintf("/role/%d", role.ID)

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
	r.HandleFunc("/role", f(Save)).Methods("POST")
	r.HandleFunc("/role/{id}", f(Get)).Methods("GET")
	r.HandleFunc("/role/{id}", f(Update)).Methods("PUT")
	r.HandleFunc("/role/{id}", f(Delete)).Methods("DELETE")

	return r
}
