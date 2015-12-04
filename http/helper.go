package http

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func logRequest(r *http.Request) {
	ip := ipAddr(r.RemoteAddr)

	log.Printf(
		"http.request.received method: %s, path: %s, ip: %s, client: %s\n",
		r.Method,
		r.RequestURI,
		ip,
		generateClientID(ip),
	)
}

func logResponse(r *http.Request, status int, startedAt time.Time) {
	ip := ipAddr(r.RemoteAddr)

	log.Printf(
		"http.request.finished method: %s, path: %s, ip: %s, client: %s, status: %d, time: %v\n",
		r.Method,
		r.RequestURI,
		ip,
		generateClientID(ip),
		status,
		time.Now().Sub(startedAt),
	)
}

func ipAddr(remoteAddr string) string {
	return strings.Split(remoteAddr, ":")[0]
}

func generateClientID(key string) string {
	md5Hash := md5.New()
	io.WriteString(md5Hash, key)
	return hex.EncodeToString(md5Hash.Sum(nil))
}
