package currenttime_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"

	. "github.com/crowdint/coffeboy/controllers/currenttime"
	"github.com/gorilla/mux"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Currenttime", func() {

	var (
		r *mux.Router
	)

	BeforeSuite(func() {
		r = CreateHandler()
	})

	Describe("The endpoint should return the time", func() {
		Context("In a 24 hours time format", func() {
			It("Should return a json with the time", func() {
				record := httptest.NewRecorder()
				req, err := http.NewRequest("GET", "/current_time", nil)
				r.ServeHTTP(record, req)

				Ω(err).Should(BeNil())

				tm := Time{}

				err = json.Unmarshal(record.Body.Bytes(), &tm)
				_, err = time.Parse("15:04:05", tm.Current)
				Ω(err).Should(BeNil())
			})
		})
	})
})

func CreateHandler() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/current_time", Get).Methods("GET")

	return r
}
