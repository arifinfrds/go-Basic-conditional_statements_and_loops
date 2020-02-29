package main

import (
	"conditional_statements_and_loops/conditional_statements"
	"conditional_statements_and_loops/loops"
)

func main() {
	runConditionalStatements()
	runLoops()
}

func runConditionalStatements() {
	conditional_statements.PrintStatementIfNeeded()
}

func runLoops() {
	loops.PrintLoops()
}
