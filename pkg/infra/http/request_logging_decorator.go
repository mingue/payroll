package http

import (
	"log"
	"net/http"
)

type RequestLoggingDecorator struct {
	l *log.Logger
	f http.Handler
}

// ServeHTTP implements http.Handler
func (d RequestLoggingDecorator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.l.Printf("Processing request to: %v - %v", r.Method, r.URL)
	d.f.ServeHTTP(w, r)
	d.l.Printf("Processed request to: %v - %v", r.Method, r.URL)
}

func NewRequestLoggingDecorator(f http.Handler, l *log.Logger) http.Handler {
	return RequestLoggingDecorator{
		l: l,
		f: f,
	}
}
