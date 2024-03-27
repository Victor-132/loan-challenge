package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/Victor-132/loan-challenge.git/application/usecase"
	"github.com/Victor-132/loan-challenge.git/domain"
	"github.com/Victor-132/loan-challenge.git/infra/httpserver"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name           string
	Input          usecase.Input
	ExpectedOutput usecase.Output
}

var fa httpserver.HttpServer[*fiber.Ctx]
var testCases []TestCase

func TestMain(t *testing.T) {
	fa = httpserver.NewFiberAdapter()
	s := domain.NewLoanStrategy()
	uc := usecase.NewGenerateLoan(s)

	httpserver.NewController(fa, uc)

	testCases = append(testCases, TestCase{
		Name: "Grant the personal loan and the guaranteed loan if the client's salary is equal to or less than R$3000",
		Input: usecase.Input{
			Age:      30,
			Document: "275.484.389-23",
			Name:     "Vuxaywua Zukiagou",
			Income:   3000.00,
			Location: "RJ",
		},
		ExpectedOutput: usecase.Output{
			Customer: "Vuxaywua Zukiagou",
			Loans: []usecase.OutputLoan{
				{Type: "PERSONAL", InterestRate: 4},
				{Type: "GUARANTEED", InterestRate: 3},
			},
		},
	})

	testCases = append(testCases, TestCase{
		Name: "Grant the personal loan and the guaranteed loan if the client's salary is between R$3000 and R$5000, if the client is under 30 years old and lives in São Paulo (SP).",
		Input: usecase.Input{
			Age:      26,
			Document: "275.484.389-23",
			Name:     "Vuxaywua Zukiagou",
			Income:   4000.00,
			Location: "SP",
		},
		ExpectedOutput: usecase.Output{
			Customer: "Vuxaywua Zukiagou",
			Loans: []usecase.OutputLoan{
				{Type: "PERSONAL", InterestRate: 4},
				{Type: "GUARANTEED", InterestRate: 3},
			},
		},
	})

	testCases = append(testCases, TestCase{
		Name: "Grant the payroll loan if the client's salary is equal to or greater than R$5000.",
		Input: usecase.Input{
			Age:      26,
			Document: "275.484.389-23",
			Name:     "Vuxaywua Zukiagou",
			Income:   5000.00,
			Location: "SP",
		},
		ExpectedOutput: usecase.Output{
			Customer: "Vuxaywua Zukiagou",
			Loans: []usecase.OutputLoan{
				{Type: "CONSIGNMENT", InterestRate: 2},
			},
		},
	})
}

func Test(t *testing.T) {
	var buf bytes.Buffer
	for _, tc := range testCases {
		t.Run(tc.Name, func(tt *testing.T) {
			if err := json.NewEncoder(&buf).Encode(tc.Input); err != nil {
				tt.Error(err)
			}

			req := httptest.NewRequest(fiber.MethodPost, "/customer-loans", &buf)
			req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

			res, err := fa.Test(req)

			if err != nil {
				tt.Error(err)
			}

			defer res.Body.Close()

			assert.Equal(t, fiber.StatusOK, res.StatusCode)

			var out usecase.Output
			if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
				tt.Error(err)
			}

			assert.Equal(t, tc.ExpectedOutput, out)
		})
	}
}
