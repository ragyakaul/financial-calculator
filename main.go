package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MortgageData struct {
	Loan                    float64 `json:"loan"`
	MandatoryMonthlyPayment float64 `json:"mandatoryMonthlyPayment"`
	OptionalMonthlyPayment  float64 `json:"optionalMonthlyPayment"`
	InterestRate            float64 `json:"annualInterestRate"`
	YearsLeftOnLoan         int     `json:"yearsLeftOnLoan"`
}

type CompoundingData struct {
	InitialPortfolioValue      float64 `json:"initialPortfolioValue"`
	AnnualBaseContribution     float64 `json:"annualBaseContribution"`
	InflationRate              float64 `json:"inflationRate"`
	GrowthRate                 float64 `json:"growthRate"`
	YearsCompounding           int     `json:"yearsCompounding"`
	TargetFIREValue            float64 `json:"targetFIREValue"`
	AnnualWithdrawalPercentage float64 `json:annualWithdrawalPercentage"`
}

func mortgageFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "Invalid method in request")
		return
	}

	var data MortgageData

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
	}

	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	fmt.Println("Received data LOAN: ", data.Loan, " INTEREST RATE: ", data.InterestRate, " YEARS LEFT: ", data.YearsLeftOnLoan)
	interestValues, principalValues, err := mortgageCalculator(data.Loan, data.MandatoryMonthlyPayment, data.OptionalMonthlyPayment, data.InterestRate, data.YearsLeftOnLoan)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Marshal interestValues and principalValues (Go -> JSON) and send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"interestValues":  interestValues,
		"principalValues": principalValues,
	})
}

func compoundInterestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "Invalid method in request")
		return
	}

	var data CompoundingData

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
	}

	portfolioValues, err := compoundingCalculator(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Marshal portfolioValues (Go -> JSON) and send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(portfolioValues)

}

func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	http.Handle("/", fileServer)
	http.HandleFunc("/api/v1/mortgage", mortgageFormHandler)
	http.HandleFunc("/api/v1/compounding", compoundInterestHandler)
	fmt.Println("Listening on port 8000")
	http.ListenAndServe("localhost:8000", nil)

}
