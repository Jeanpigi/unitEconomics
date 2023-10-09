package main

import "fmt"

func main() {
	var monthlyRevenue float64
	var costOfGoodsSold float64
	var monthlyMarketingExpense float64
	var numberUnitSold float64

	fmt.Println("Welcome to a program for Unit Economics:")
	fmt.Print("Write your Monthly Revenue: $")
	fmt.Scanf("%f", &monthlyRevenue)
	fmt.Print("Write your Cost of Goods Sold: $")
	fmt.Scanf("%f", &costOfGoodsSold)
	fmt.Print("Write your Monthly Marketing Expense: $")
	fmt.Scanf("%f", &monthlyMarketingExpense)
	fmt.Print("Write your Number of Units Sold: ")
	fmt.Scanf("%f", &numberUnitSold)

	// Check for division by zero
	if numberUnitSold == 0 {
		fmt.Println("Number of Units Sold cannot be zero.")
		return
	}

	cac := monthlyMarketingExpense / numberUnitSold
	arpu := monthlyRevenue / numberUnitSold
	grossMarginUnit := (monthlyRevenue - costOfGoodsSold) / numberUnitSold

	fmt.Printf("CAC: $%.2f\n", cac)
	fmt.Printf("ARPU: $%.2f\n", arpu)
	fmt.Printf("Gross Margin per Unit: $%.2f\n", grossMarginUnit)
}
