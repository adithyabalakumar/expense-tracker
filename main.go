package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func printUsage() {
	fmt.Println(`
Usage: expense-tracker [COMMAND] [ARGUMENTS]

COMMAND:
	add: To add an expense.
		Example: expense-tracker add -description <DESCRIPTION> -amount <AMOUNT>
	update: To update the description of an existing expense. Must pass the id of the expense and a new description and amount.
		Example: expense-tracker update -id <ID> -description <DESCRIPTION> -amount <AMOUNT>
	delete: To delete an existing expense. Must pass the id of the task to be deleted.
		Example: expense-tracker delete -id <ID>
	list: List all expenses.
		Example: expense-tracker list
	summary: Summary of expenses.
		Example: expense-tracker summary -month <MONTH>
	`)
}

func main() {
	// subcommand: add
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addDescription := addCmd.String("description", "", "description")
	addAmount := addCmd.String("amount", "", "amount")

	// subcommand: update
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateId := updateCmd.Int("id", 0, "id")
	updateDescription := updateCmd.String("description", "", "description")
	updateAmount := updateCmd.String("amount", "", "amount")

	// subcommand: delete
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteId := deleteCmd.Int("id", 0, "id")

	// subcommand: summary
	summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
	summaryMonth := summaryCmd.Int("month", 0, "month")

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// load expenses
	expenseList := loadExpenses()

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		addExpense(&expenseList, *addDescription, *addAmount)
	case "update":
		updateCmd.Parse(os.Args[2:])
		updateExpense(&expenseList, *updateId, *updateDescription, *updateAmount)
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		deleteExpense(&expenseList, *deleteId)
	case "list":
		listExpenses(&expenseList)
	case "summary":
		summaryCmd.Parse(os.Args[2:])
		month := time.Month(*summaryMonth)
		expensesSummary(&expenseList, month)
	default:
		printUsage()
		os.Exit(1)
	}
}
