package server

import "net/http"

func responseHealthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	body := []byte("OK")
	w.Write(body)
}
