package domain

type Strategy interface {
	Generate(income, age int, location string) []Loan
}
