package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	i "vm.io/payroll/invoice"
)

func main() {
	inv1 := i.Invoice{}
	inv1.ID = 1234

	fmt.Printf("Hello World!\n %v", inv1.ID)

	s := GetInvoiceService()
	invoices := s.GetAllByFinancialYear(uint16(time.Now().Year()))
	fmt.Printf("Obtained %v invoices", len(invoices))

	http.HandleFunc("/invoices", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err := json.NewEncoder(w).Encode(invoices)

		if err != nil {
			log.Fatalf("Couldn't serialize invoices %v", err.Error())	
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Couldn't start the server: %v", err.Error())
	}
}

func GetInvoiceService() i.InvoiceService {
	return i.NewInvoiceService(&i.InvoiceRepository{})
}
