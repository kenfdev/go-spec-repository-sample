package inmem

import (
	"fmt"
	"github.com/kenfdev/go-spec-repository-sample/domain"
	"time"
)

type InMemoryInvoiceRepository struct {
	invoices []*domain.Invoice
}

func NewInMemoryInvoiceRepository() *InMemoryInvoiceRepository {
	dd1, _ := time.Parse(time.RFC3339, "2014-11-11T11:11:11.111Z")
	inv1 := &domain.Invoice{
		ID:      "inv1",
		Total:   100,
		DueDate: dd1,
	}
	dd2, _ := time.Parse(time.RFC3339, "2020-11-11T11:11:11.111Z")
	inv2 := &domain.Invoice{
		ID:      "inv2",
		Total:   200,
		DueDate: dd2,
	}
	return &InMemoryInvoiceRepository{
		invoices: []*domain.Invoice{inv1, inv2},
	}
}

func (i *InMemoryInvoiceRepository) SelectWhereGracePeriodPast(aDate time.Time) ([]*domain.Invoice, error) {
	sql := i.whereGracePeriodPastSQL(aDate)
	fmt.Printf("%s\n", sql)
	// emulate an SQL execution
	fmt.Println("Executing SQL...")
	time.Sleep(3 * time.Second)
	fmt.Println("Execution Finished...")

	var invoices []*domain.Invoice
	for _, inv := range i.invoices {
		if inv.DueDate.Unix() < aDate.Unix() {
			invoices = append(invoices, inv)
		}
	}

	return invoices, nil
}

func (i *InMemoryInvoiceRepository) whereGracePeriodPastSQL(aDate time.Time) string {
	return `SELECT * FROM INVOICE, CUSTOMER
 WHERE INVOICE.CUST_ID = CUSTOMER.ID
 AND INVOICE.DUE_DATE + CUSTOMER.GRACE_PERIOD < ` + aDate.Format("2006-01-02T15:04:05")
}

func (i *InMemoryInvoiceRepository) SelectSatisfying(spec domain.InvoiceSpecification) ([]*domain.Invoice, error) {
	return spec.SatisfyingElementsFrom(i)
}
