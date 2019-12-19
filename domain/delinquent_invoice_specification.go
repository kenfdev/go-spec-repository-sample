package domain

import "time"

type DelinquentInvoiceSpecification struct {
	currentDate time.Time
}

func NewDelinquentInvoiceSpecification(currentDate time.Time) *DelinquentInvoiceSpecification {
	return &DelinquentInvoiceSpecification{currentDate: currentDate}
}

func (spec *DelinquentInvoiceSpecification) SatisfyingElementsFrom(repository InvoiceRepository) ([]*Invoice, error) {
	return repository.SelectWhereGracePeriodPast(spec.currentDate)
}
