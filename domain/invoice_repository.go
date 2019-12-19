package domain

import "time"

type InvoiceRepository interface {
	SelectSatisfying(spec *DelinquentInvoiceSpecification) ([]*Invoice, error)
	SelectWhereGracePeriodPast(aDate time.Time) ([]*Invoice, error)
}
