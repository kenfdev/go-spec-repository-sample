package main

import (
	"fmt"
	"github.com/kenfdev/go-spec-repository-sample/app"
	"github.com/kenfdev/go-spec-repository-sample/infrastructure/inmem"
)

func main() {

	repo := inmem.NewInMemoryInvoiceRepository()

	usecase := app.NewListGracePeriodPastInvoiceUseCase(repo)

	response, _ := usecase.Execute()

	fmt.Println("Results:")
	for _, inv := range response.Invoices {
		fmt.Printf("id: %s, total: %d, due_date: %s\n", inv.ID, inv.Total, inv.DueDate.Format("2006-01-02T15:04:05"))
	}
}
