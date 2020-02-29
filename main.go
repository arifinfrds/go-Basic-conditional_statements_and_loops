package main

import (
	"conditional_statements_and_loops/conditional_statements"
	"conditional_statements_and_loops/loops"
	"conditional_statements_and_loops/switch_statement"
)

func main() {
	runConditionalStatements()
	runLoops()
	runSwitchStatement()
}

func runConditionalStatements() {
	conditional_statements.PrintStatementIfNeeded()
}

func runLoops() {
	loops.PrintLoops()
}

func runSwitchStatement() {
	switch_statement.PrintSwitchStatement()
}
