// Package http implements handlers used by rebasebot http server
package http

import (
	"fmt"
	"net/http"
)

func Status(w http.ResponseWriter, r *http.Request) {
	logRequest(r)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK\n")

	logResponse(r, http.StatusOK)
}
