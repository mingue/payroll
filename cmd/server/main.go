package main

import (
	"log"
	"net/http"
	"os"

	handlers "github.com/mingue/payroll/cmd/server/handlers"
	h "github.com/mingue/payroll/pkg/infra/http"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

	var getInvoicesDecorated http.Handler
	getInvoicesDecorated = h.NewRequestLoggingDecorator(handlers.GetInvoiceHandler{}, logger)
	getInvoicesDecorated = h.NewPanicRecoveryDecorator(getInvoicesDecorated, logger)

	http.Handle("/invoices", getInvoicesDecorated)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Couldn't start the server: %v", err.Error())
	}
}
