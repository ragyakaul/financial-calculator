package main

import "fmt"

func mortgageCalculator(loan float64, mandatoryMonthlyPayment float64, optionalMonthlyPayment float64, interestRate float64, yearsLeftOnLoan int) ([]float64, []float64, error) {

	if yearsLeftOnLoan <= 0 {
		err := fmt.Errorf("years left on loan must be greater 0")
		return nil, nil, err
	}

	remainingLoan := loan
	monthlyInterestRate := (interestRate / 100) / 12
	monthsLeftOnLoan := yearsLeftOnLoan * 12

	interestValues := make([]float64, int(monthsLeftOnLoan))
	principalValues := make([]float64, int(monthsLeftOnLoan))
	sumInterest := 0.0

	for i := 0; i < int(monthsLeftOnLoan); i++ {
		interestValues[i] = remainingLoan * monthlyInterestRate
		sumInterest += interestValues[i]
		principalValues[i] = mandatoryMonthlyPayment - interestValues[i]
		remainingLoan -= principalValues[i] + optionalMonthlyPayment
		if principalValues[i] > interestValues[i] {
			fmt.Println("Principal payment is greater than interest payment")
		} else {
			fmt.Println("Principal payment is less than interest payment")
		}
	}
	fmt.Println("Total interest paid: ", sumInterest)
	return interestValues, principalValues, nil
}
