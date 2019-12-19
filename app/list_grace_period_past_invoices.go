package app

import (
	"github.com/kenfdev/go-spec-repository-sample/domain"
	"time"
)

type ListGracePeriodPastInvoiceResponse struct {
	Invoices []*domain.Invoice
}

type ListGracePeriodPastInvoiceUseCase struct {
	repository domain.InvoiceRepository
}

func NewListGracePeriodPastInvoiceUseCase(repository domain.InvoiceRepository) *ListGracePeriodPastInvoiceUseCase {
	return &ListGracePeriodPastInvoiceUseCase{repository: repository}
}

func (uc *ListGracePeriodPastInvoiceUseCase) Execute() (*ListGracePeriodPastInvoiceResponse, error) {
	now := time.Now()
	spec := domain.NewDelinquentInvoiceSpecification(now)
	invoices, _ := uc.repository.SelectSatisfying(spec)
	response := &ListGracePeriodPastInvoiceResponse{Invoices: invoices}
	return response, nil
}
