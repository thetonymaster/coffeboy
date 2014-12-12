package userscontroller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/coopernurse/gorp"
	. "github.com/crowdint/coffeboy/controllers/userscontroller"
	"github.com/crowdint/coffeboy/models/users"
	"github.com/gorilla/mux"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Userscontroller", func() {

	var (
		dbmap *gorp.DbMap
		r     *mux.Router
	)

	BeforeSuite(func() {
		dbmap, _ = users.InitDb()
		r = CreateHandler(CreateDbMapHandlerToHTTPHandler(dbmap))
	})

	Describe("Create a user", func() {
		Context("Send a POST request", func() {

			It("Should return status created", func() {

				record := httptest.NewRecorder()
				message := `{ "role-id": 1, "name": "Nig", "last-name": "nog", "password": "12345", "image-url": "http://foo/200x300", "channel": "9999", "uuid": "uuid", "device-token": "lalalla", "email": "1234@1234.com" }`
				req, err := http.NewRequest("POST", "/user", bytes.NewReader([]byte(message)))
				r.ServeHTTP(record, req)

				Expect(err).To(BeNil())
				Expect(record.Code).To(Equal(http.StatusCreated))

			})
		})
	})

	Describe("Get a user from the database", func() {
		Context("Send a GET request", func() {
			It("Should return a user JSON object", func() {
				record := httptest.NewRecorder()
				message := `{ "role-id": 1, "name": "Nig", "last-name": "nog", "password": "12345", "image-url": "http://foo/200x300", "channel": "9999", "uuid": "uuid", "device-token": "lalalla", "email": "1234@1234.com" }`
				req, err := http.NewRequest("POST", "/user", bytes.NewReader([]byte(message)))
				r.ServeHTTP(record, req)

				Expect(err).To(BeNil())
				Expect(record.Code).To(Equal(http.StatusCreated))

				var user users.User

				err = json.Unmarshal(record.Body.Bytes(), &user)

				url := fmt.Sprintf("/user/%d", user.ID)

				record2 := httptest.NewRecorder()
				req, err = http.NewRequest("GET", url, nil)
				r.ServeHTTP(record2, req)

				Expect(record2.Code).To(Equal(http.StatusOK))

			})
		})
	})

	Describe("Update a user from the database", func() {
		Context("Send a PUT request", func() {
			It("Should return a user JSON object", func() {
				record := httptest.NewRecorder()
				message := `{ "role-id": 1, "name": "Nig", "last-name": "nog", "password": "12345", "image-url": "http://foo/200x300", "channel": "9999", "uuid": "uuid", "device-token": "lalalla", "email": "1234@1234.com" }`
				req, err := http.NewRequest("POST", "/user", bytes.NewReader([]byte(message)))
				r.ServeHTTP(record, req)

				Expect(err).To(BeNil())
				Expect(record.Code).To(Equal(http.StatusCreated))

				var user users.User

				err = json.Unmarshal(record.Body.Bytes(), &user)

				url := fmt.Sprintf("/user/%d", user.ID)

				user.Name = "breakfast at tiffany's"

				body, err := json.Marshal(user)

				req2, err := http.NewRequest("PUT", url, bytes.NewReader([]byte(body)))
				record2 := httptest.NewRecorder()
				r.ServeHTTP(record2, req2)

				Expect(record2.Code).To(Equal(http.StatusOK))

				req3, err := http.NewRequest("GET", url, nil)
				record3 := httptest.NewRecorder()
				r.ServeHTTP(record3, req3)

				var product2 users.User
				err = json.Unmarshal(record3.Body.Bytes(), &product2)

				Expect(user).To(Equal(product2))

			})
		})
	})

	Describe("Delete a user from the database", func() {
		Context("Send a DELETE request", func() {
			It("Should delete the object", func() {
				record := httptest.NewRecorder()
				message := `{ "role-id": 1, "name": "Nig", "last-name": "nog", "password": "12345", "image-url": "http://foo/200x300", "channel": "9999", "uuid": "uuid", "device-token": "lalalla", "email": "1234@1234.com" }`
				req, err := http.NewRequest("POST", "/user", bytes.NewReader([]byte(message)))
				r.ServeHTTP(record, req)

				Expect(err).To(BeNil())
				Expect(record.Code).To(Equal(http.StatusCreated))

				var user users.User

				err = json.Unmarshal(record.Body.Bytes(), &user)

				url := fmt.Sprintf("/user/%d", user.ID)

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
	r.HandleFunc("/user", f(Save)).Methods("POST")
	r.HandleFunc("/user/{id}", f(Get)).Methods("GET")
	r.HandleFunc("/user/{id}", f(Update)).Methods("PUT")
	r.HandleFunc("/user/{id}", f(Delete)).Methods("DELETE")

	return r
}
