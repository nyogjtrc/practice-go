package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendPost(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			t.Error(err)
		}

		j, err := json.Marshal(r.Form)
		if err != nil {
			t.Error(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	}

	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	host = ts.URL

	err := sendPOSTform()
	if err != nil {
		t.Error(err)
	}
}
