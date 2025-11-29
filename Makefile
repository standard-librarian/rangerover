URL  ?= https://http.cat/429
DIST ?= /Users/medhatmohammed/Documents/goprojects/rangerover/data/file.txt
N    ?= 4
S    ?= 1023
R    ?= 5

# --- Phony Targets (tasks that are not actual files) ---

.PHONY: run run-help run-args test clean tidy lint

# The default goal when running 'make' without arguments
.DEFAULT_GOAL := run-args

run:
	go run .
run-args:
	go run . -url $(URL) -dist $(DIST) -n $(N) -s $(S) -r $(R)
run-help:
	go run . --help

tidy:
	go mod tidy
