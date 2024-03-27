package domain

type LoanStrategy struct{}

func NewLoanStrategy() Strategy {
	return &LoanStrategy{}
}

// Generate implements Strategy.
func (l *LoanStrategy) Generate(income int, age int, location string) (loans []Loan) {
	if income <= 3000 {
		loans = append(loans, personal)
		loans = append(loans, guaranteed)
	}

	if income > 3000 && income < 5000 && age < 30 && location == "SP" {
		loans = append(loans, personal)
		loans = append(loans, guaranteed)
	}

	if income >= 5000 {
		loans = append(loans, consigment)
	}

	return
}
