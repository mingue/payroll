package http

import (
	"fmt"
	"log"
	"net/http"
)

type PanicRecoveryDecorator struct {
	logger *log.Logger
	handler http.Handler
}

func NewPanicRecoveryDecorator(f http.Handler, l *log.Logger) http.Handler {
	return PanicRecoveryDecorator{
		logger: l,
		handler: f,
	}
}

// ServeHTTP implements http.Handler
func (d PanicRecoveryDecorator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			m := fmt.Sprintf("Ooops panic processing request %v - %v : %v", r.Method, r.URL, err)
			d.logger.Print(m)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, m)
		}
	}()

	d.handler.ServeHTTP(w, r)
}
