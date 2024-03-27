package main

import (
	"github.com/Victor-132/loan-challenge.git/application/usecase"
	"github.com/Victor-132/loan-challenge.git/domain"
	"github.com/Victor-132/loan-challenge.git/infra/httpserver"
)

func main() {
	fa := httpserver.NewFiberAdapter()
	s := domain.NewLoanStrategy()
	uc := usecase.NewGenerateLoan(s)

	httpserver.NewController(fa, uc)

	fa.Listen(3000)
}
