package main

import "fmt"

func compoundingCalculator(compoundingData CompoundingData) ([]float64, error) {
	portfolioValues := make([]float64, compoundingData.YearsCompounding)
	currentPortfolioValue := compoundingData.InitialPortfolioValue
	fireReached := false
	for i := 0; i < compoundingData.YearsCompounding; i++ {

		netGrowth := (compoundingData.GrowthRate - compoundingData.InflationRate) / 100
		if currentPortfolioValue < compoundingData.TargetFIREValue && !fireReached {
			fmt.Println("Year Less Than 1 Million: ", i)
			amountToAdd := compoundingData.AnnualBaseContribution + (currentPortfolioValue * netGrowth)
			currentPortfolioValue = currentPortfolioValue + amountToAdd
			portfolioValues[i] = currentPortfolioValue
		} else {
			fireReached = true
			fmt.Println("Year More Than 1 Million: ", i)
			amountToAdd := currentPortfolioValue * netGrowth                                               // 1500 x 0.04
			amountToSubtract := currentPortfolioValue * (compoundingData.AnnualWithdrawalPercentage / 100) // 1500 * 0.04
			currentPortfolioValue = currentPortfolioValue + amountToAdd - amountToSubtract                 // 1500 + 60 - 60
			portfolioValues[i] = currentPortfolioValue
		}

		fmt.Println("Current Portfolio Value: ", portfolioValues[i])
	}
	return portfolioValues, nil
}
