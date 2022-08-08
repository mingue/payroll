package main

import (
	"fmt"
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
}

func GetInvoiceService() i.InvoiceService {
	return i.NewInvoiceService(&i.InvoiceRepository{})
}
