package webby

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
	http.ResponseWriter
	Status int
	Bytes  int
}
func (sr *StatusResponse) Write(buf []byte) (bytes int, err error) {
	if sr.Status == 0 {
		sr.Status = 200
	}
	bytes, err = sr.ResponseWriter.Write(buf)
	sr.Bytes += bytes
	return
}
func (sr *StatusResponse) WriteHeader(status int) {
	sr.Status = status
	sr.ResponseWriter.WriteHeader(status)
}

type LogHandler struct {
	http.Handler
}
func (lh LogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sr := &StatusResponse{w, 0, 0}
	lh.Handler.ServeHTTP(sr, r)
	log.Printf("%d %d %s %s %s %s", sr.Status, sr.Bytes,
		r.RemoteAddr, r.Host, r.Method, r.RequestURI)
}
