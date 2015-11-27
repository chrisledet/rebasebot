package http

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"net/http"
)

func logRequest(r *http.Request) {
	log.Printf(
		"http.request.received method: %s, path: %s, ip: %s, client: %s\n",
		r.Method,
		r.RequestURI,
		r.RemoteAddr,
		generateClientId(r.RemoteAddr),
	)
}

func logResponse(r *http.Request, status int) {
	log.Printf(
		"http.request.finished method: %s, path: %s, ip: %s, client: %s, status: %d\n",
		r.Method,
		r.RequestURI,
		r.RemoteAddr,
		generateClientId(r.RemoteAddr),
		status,
	)
}

func generateClientId(remoteIp string) string {
	md5Hash := md5.New()
	io.WriteString(md5Hash, remoteIp)
	return hex.EncodeToString(md5Hash.Sum(nil))
}
