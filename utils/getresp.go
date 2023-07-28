package utils

import (
	"io/ioutil"
	"net/http"
	"time"
)

func FordardGet(h string, timeout int, r *http.Request, w http.ResponseWriter) string {
	req, err := http.NewRequest("GET", h, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return ""
	}

	// set timeout
	client := &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "[ERROR] error occur when request"
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "[ERROR] error occur when getbody"
	}
	defer resp.Body.Close()
	return string(body)
}
