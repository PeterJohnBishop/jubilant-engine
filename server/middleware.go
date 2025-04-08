package server

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type ResponseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func (rw *ResponseWriterWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *ResponseWriterWrapper) Write(data []byte) (int, error) {
	rw.body.Write(data)
	return rw.ResponseWriter.Write(data)
}

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapper := &ResponseWriterWrapper{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
			body:           new(bytes.Buffer),
		}

		next(wrapper, r)

		duration := time.Since(start)
		logMsg := fmt.Sprintf("[%s] %s from %s - Status: %d, Duration: %s",
			r.Method, r.URL.Path, r.RemoteAddr, wrapper.statusCode, duration)

		if wrapper.statusCode >= 400 {
			logMsg += fmt.Sprintf(" - Response: %s", wrapper.body.String())
		}

		log.Println(logMsg)
	}
}

func KubernetesPod(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		podName := os.Getenv("HOSTNAME")
		w.Header().Set("X-Pod-Name", podName)
		fmt.Println("Request received at pod:", podName)
		next.ServeHTTP(w, r)
	})
}
