package main

import "fmt"

func mortgageCalculator(loan float64, interestRate float64, yearsLeftOnLoan float64) (float64, float64, float64, float64, error) {

	if yearsLeftOnLoan <= 0 {
		err := fmt.Errorf("Years left on loan must be greater 0")
		return 0, 0, 0, 0, err
	}
	currentMonthlyHomeInterestRate := (interestRate / 12) / 100
	monthlyInterestPayment := loan * currentMonthlyHomeInterestRate
	yearlyInterestPayment := monthlyInterestPayment * 12
	yearlyPrincipalPayment := loan / yearsLeftOnLoan
	monthlyPrincipalPayment := yearlyPrincipalPayment / 12
	monthlyTotalLoanPayment := monthlyPrincipalPayment + monthlyInterestPayment
	yearlyTotalLoanPayment := monthlyTotalLoanPayment * 12

	fmt.Println("Monthly Principal: ", monthlyPrincipalPayment)
	fmt.Println("Monthly Interest: ", monthlyInterestPayment)
	fmt.Println("Total Monthly Loan: ", monthlyTotalLoanPayment)
	fmt.Println("Yearly Principal: ", yearlyPrincipalPayment)
	fmt.Println("Yearly Interest: ", yearlyInterestPayment)
	fmt.Println("Total Yearly Loan: ", yearlyTotalLoanPayment)
	return monthlyPrincipalPayment, yearlyPrincipalPayment, monthlyInterestPayment, yearlyInterestPayment, nil
}
