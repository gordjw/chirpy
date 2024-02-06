package server

import (
	"fmt"
	"net/http"
)

func (c *config) responseMetrics(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	body := []byte(fmt.Sprintf("Hits: %d", c.fileServerHits))
	w.Write(body)
}
