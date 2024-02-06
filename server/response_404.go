package server

import (
	"fmt"
	"net/http"
)

func response404(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "404 not found")
}
