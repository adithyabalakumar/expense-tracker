.PHONY: build
build: expense-tracker

.PHONY: expense-tracker
expense-tracker:
	go build -o $@ main.go expenses.go
	@echo "$@: build successful"

.PHONY: clean
clean:
	$(RM) expense-tracker