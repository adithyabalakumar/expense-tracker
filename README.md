# Expense Tracker

A simple CLI based expense tracker to manage your finances.

This is part of a list of [backend projects](https://roadmap.sh/backend/projects) to practice from [roadmaps.sh](https://roadmap.sh/).

The CLI implements listing, adding, updating and deleting expenses. As it stands it's a pretty minimal application that implements the necessary function as described by the [project page](https://roadmap.sh/projects/expense-tracker).

See [Usage](#usage) for examples.

## Build the project

Run `make build` to build the project.

## Usage
### Add expense
Adding an expense expects the `description` and `amount` for the expense as a positional argument.
```
# Adding a new expense
expense-tracker add --description "Lunch" --amount 20
```

### Update expense
Updating an expense expects the `id`, `description` and `amount` for the expense as positional arguments.
```
# Updating expense
expense-tracker update --id 2 --description "Dinner" --amount 10
```

### Delete expense
Deleting an expense expects the `id` of the expense as a positional argument.
```
# Deleting expense
expense-tracker delete --id 2
```

### Listing expenses
Existing expenses can be listed with the `list` sub-command.
```
# Listing all expenses
expense-tracker list
```

### Expenses summary
Summary of the expenses can be calculated using the `summary` sub-command. Optionally, a `month` can be passed to calculate the summary for just the `month` specified
```
# Summary of the expenses
expense-tracker summary

# Summary for the month of August
expense-tracker summary --month 8
``` 