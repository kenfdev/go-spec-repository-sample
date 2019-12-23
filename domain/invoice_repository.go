package domain

import "time"

type InvoiceRepository interface {
	SelectSatisfying(spec InvoiceSpecification) ([]*Invoice, error)
	SelectWhereGracePeriodPast(aDate time.Time) ([]*Invoice, error)
}
