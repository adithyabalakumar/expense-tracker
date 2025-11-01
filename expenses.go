package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
)

type Expense struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ExpenseList struct {
	Expenses []Expense `json:"expenses"`
	NextID   int       `json:"nextId"`
}

const expenseRegistry string = "expenses.json"

func loadExpenses() ExpenseList {
	tasks, err := os.ReadFile(expenseRegistry)
	if err != nil {
		return ExpenseList{Expenses: []Expense{}, NextID: 1}
	}

	var expenseList ExpenseList
	if err := json.Unmarshal(tasks, &expenseList); err != nil {
		return ExpenseList{Expenses: []Expense{}, NextID: 1}
	}
	return expenseList
}

func saveExpenses(expenseList *ExpenseList) error {
	jsonData, err := json.Marshal(expenseList)
	if err != nil {
		fmt.Printf("Error in converting expense list to JSON: %v\n", err)
		return err
	}

	err = os.WriteFile(expenseRegistry, jsonData, 0644)
	if err != nil {
		fmt.Printf("Writing to %s failed: %v\n", expenseRegistry, err)
		return err
	}
	return nil
}

func addExpense(expenseList *ExpenseList, description string, amount string) {
	if len(description) == 0 {
		fmt.Println("Description cannot be empty")
	}

	expenseAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		fmt.Println("Invalid amount")
		os.Exit(1)
	}

	if expenseAmount < 0 {
		fmt.Println("Expense amount cannot be less than 0")
		os.Exit(1)
	}

	expense := Expense{
		Id:          expenseList.NextID,
		Description: description,
		Amount:      expenseAmount,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	expenseList.Expenses = append(expenseList.Expenses, expense)
	expenseList.NextID++

	err = saveExpenses(expenseList)
	if err == nil {
		fmt.Printf("Expense with ID: %d saved to %s\n", expense.Id, expenseRegistry)
	}
}

func updateExpense(expenseList *ExpenseList, id int, description string, amount string) {
	expenseId := id
	expenseAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		fmt.Println("Invalid amount")
		os.Exit(1)
	}

	if expenseAmount < 0 {
		fmt.Println("Expense amount cannot be less than 0")
		os.Exit(1)
	}

	for i := range expenseList.Expenses {
		if expenseList.Expenses[i].Id == expenseId {
			expenseList.Expenses[i].Amount = expenseAmount
			if len(description) != 0 {
				expenseList.Expenses[i].Description = description
			}
			expenseList.Expenses[i].UpdatedAt = time.Now()
			err := saveExpenses(expenseList)
			if err == nil {
				fmt.Printf("Expense with ID: %d updated\n", expenseId)
				return
			}
		}
	}

	fmt.Printf("Expense with ID: %d not found\n", expenseId)
}

func deleteExpense(expenseList *ExpenseList, id int) {
	expenseId := id

	for i := range expenseList.Expenses {
		if expenseList.Expenses[i].Id == expenseId {
			expenseList.Expenses = append(expenseList.Expenses[:i], expenseList.Expenses[i+1:]...)
			err := saveExpenses(expenseList)
			if err == nil {
				fmt.Printf("Expense with ID: %d deleted\n", expenseId)
				return
			}
		}
	}

	fmt.Printf("Expense with ID: %d not found\n", expenseId)
}

func listExpenses(expenseList *ExpenseList) {
	if len(expenseList.Expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tDate\tDescription\tAmount")
	fmt.Fprintln(w, "----\t----\t----\t----")
	for _, expense := range expenseList.Expenses {
		fmt.Fprintf(w, "%d\t%s\t%s\t%.2f\n", expense.Id, expense.CreatedAt.Format("2006-01-02"), expense.Description, expense.Amount)
	}
	w.Flush()
}

func expensesSummary(expenseList *ExpenseList, month time.Month) {
	if len(expenseList.Expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	summary := 0.0
	currentYear := time.Now().Year()
	for _, expense := range expenseList.Expenses {
		if month == 0 || (expense.CreatedAt.Month() == month && expense.CreatedAt.Year() == currentYear) {
			summary += expense.Amount
		}
	}
	fmt.Printf("Total expenses: %.2f\n", summary)
}
