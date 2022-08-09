package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	r "github.com/mingue/payroll/pkg/infra/repositories"
	i "github.com/mingue/payroll/pkg/invoice"
)

type GetInvoiceHandler struct{}

func (d GetInvoiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := NewInvoiceService()
	invoices := s.GetAllByFinancialYear(uint16(time.Now().Year()))
	fmt.Printf("Obtained %v invoices\n", len(invoices))

	panic("Panic!! trying to recover")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(invoices)

	if err != nil {
		log.Fatalf("Couldn't serialize invoices %v", err.Error())
	}
}

func NewInvoiceService() i.Service {
	return i.NewService(r.InvoiceRepository{})
}
