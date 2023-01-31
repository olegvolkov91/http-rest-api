package apiserver

import "net/http"

type ResponseWriter struct {
	http.ResponseWriter
	code int
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.code = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}
