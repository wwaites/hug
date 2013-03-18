package hutil

import (
	"encoding/json"
	"log"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, r interface{}) {
	w.Header().Set("Content-Type", "text/javascript")
	b, err := json.Marshal(r)
	if err == nil {
		w.Write(b)
	}
	return
}

type StatusResponse struct {
	Status int
	Bytes  int
	writer http.ResponseWriter
}
func (sr *StatusResponse) Header() http.Header {
	return sr.writer.Header()
}
func (sr *StatusResponse) Write(buf []byte) (bytes int, err error) {
	if sr.Status == 0 {
		sr.Status = 200
	}
	bytes, err = sr.writer.Write(buf)
	sr.Bytes += bytes
	return
}
func (sr *StatusResponse) WriteHeader(status int) {
	sr.Status = status
	sr.writer.WriteHeader(status)
}

type LogHandler struct {
	h http.Handler
}
func (lh LogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sr := &StatusResponse{writer: w}
	lh.h.ServeHTTP(sr, r)
	log.Printf("%d %d %s %s %s %s", sr.Status, sr.Bytes,
		r.RemoteAddr, r.Host, r.Method, r.RequestURI)
}
func NewLogHandler(h http.Handler) http.Handler {
	return LogHandler{h}
}
func LogHandlerFunc(f func(http.ResponseWriter, *http.Request)) http.Handler {
	return LogHandler{http.HandlerFunc(f)}
}

