package domain

type Loan struct {
	Type         string
	InterestRate int
}

var (
	consigment = Loan{Type: "CONSIGNMENT", InterestRate: 2}
	guaranteed = Loan{Type: "GUARANTEED", InterestRate: 3}
	personal   = Loan{Type: "PERSONAL", InterestRate: 4}
)
