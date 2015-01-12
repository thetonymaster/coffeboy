package currenttime

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Time struct {
	Current string `json:"current_time"`
}

func Get(w http.ResponseWriter, r *http.Request) {
	hour := time.Now().Format("15:04:05")
	tm := Time{
		Current: hour,
	}

	response, err := json.Marshal(tm)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
