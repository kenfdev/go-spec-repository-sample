package domain

type InvoiceSpecification interface {
	SatisfyingElementsFrom(repository InvoiceRepository) ([]*Invoice, error)
}
