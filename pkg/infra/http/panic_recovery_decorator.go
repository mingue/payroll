package http

import (
	"fmt"
	"log"
	"net/http"
)

type PanicRecoveryDecorator struct {
	l *log.Logger
	f http.Handler
}

// ServeHTTP implements http.Handler
func (d PanicRecoveryDecorator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			m := fmt.Sprintf("Ooops panic processing request %v - %v : %v", r.Method, r.URL, err)
			d.l.Print(m)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, m)
		}
	}()

	d.f.ServeHTTP(w, r)
}

func NewPanicRecoveryDecorator(f http.Handler, l *log.Logger) http.Handler {
	return PanicRecoveryDecorator{
		l: l,
		f: f,
	}
}
