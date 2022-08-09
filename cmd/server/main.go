package main

import (
	"log"
	"net/http"
	"os"

	handlers "github.com/mingue/payroll/cmd/server/handlers"
	h "github.com/mingue/payroll/pkg/infra/http"
)

func main() {
	logger := log.New(os.Stdout, "Payroll", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

	http.Handle("/invoices", GetInvoicesHandler(logger))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Couldn't start the server: %v", err.Error())
	}
}

func GetInvoicesHandler(logger *log.Logger) http.Handler {
	getInvoicesDecorated := h.NewRequestLoggingDecorator(handlers.NewInvoiceHandler(), logger)
	getInvoicesDecorated = h.NewPanicRecoveryDecorator(getInvoicesDecorated, logger)
	return getInvoicesDecorated
}
