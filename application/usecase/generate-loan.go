package usecase

import "github.com/Victor-132/loan-challenge.git/domain"

type Input struct {
	Age      int
	Document string
	Name     string
	Income   int
	Location string
}

type Output struct {
	Customer string       `json:"customer"`
	Loans    []OutputLoan `json:"loans"`
}

type OutputLoan struct {
	Type         string `json:"type"`
	InterestRate int    `json:"interest_rate"`
}

type GenerateLoan struct {
	s domain.Strategy
}

func NewGenerateLoan(s domain.Strategy) UseCase[Input, Output] {
	return &GenerateLoan{s}
}

// Execute implements UseCase.
func (g *GenerateLoan) Execute(input Input) Output {
	var out Output = Output{
		Customer: input.Name,
	}

	for _, l := range g.s.Generate(input.Income, input.Age, input.Location) {
		out.Loans = append(out.Loans, OutputLoan(l))
	}

	return out
}
