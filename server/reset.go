package server

import (
	"net/http"
)

func (c *config) responseReset(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	c.fileServerHits = 0
	body := []byte("Counter reset to 0")
	w.Write(body)
}
