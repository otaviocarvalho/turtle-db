package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MetaCommandResult int
type ExitResult int
type PrepareResultStatement int
type StatementType int

const (
	MetaCommandUnrecognized MetaCommandResult = iota
	MetaCommandSuccess
)

const (
	ExitSuccess ExitResult = iota
	ExitFailure
)

const (
	PrepareSuccessStatement PrepareResultStatement = iota
	PrepareUnrecognizedStatement
)

const (
	StatementInsert StatementType = iota
	StatementSelect
)

type Statement struct {
	Type StatementType
}

var statementType = map[string]StatementType{
	"insert": StatementInsert,
	"select": StatementSelect,
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("db > ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.HasPrefix(input, ".") {
			switch doMetaCommand(input) {
			case MetaCommandSuccess:
				continue
			case MetaCommandUnrecognized:
				fmt.Printf("Unrecognized command '%s'.\n", input)
				continue
			}
		}

		var statement Statement
		if result, ok := prepareStatement(input, &statement); ok {
			executeStatement(&statement)
			fmt.Println("Executed.")
		} else {
			fmt.Printf("Unrecognized keyword at start of '%s'. PrepareResultStatement=%d\n", input, result)
		}
	}
}

func doMetaCommand(input string) MetaCommandResult {
	if input == ".exit" {
		os.Exit(int(ExitSuccess))
	}
	return MetaCommandUnrecognized
}

func prepareStatement(input string, statement *Statement) (PrepareResultStatement, bool) {
	for keyword, stmtType := range statementType {
		if strings.HasPrefix(input, keyword) {
			statement.Type = stmtType
			return PrepareSuccessStatement, true
		}
	}

	return PrepareUnrecognizedStatement, false
}

func executeStatement(statement *Statement) {
	switch statement.Type {
	case StatementInsert:
		fmt.Println("This is where we should do an insert.")
	case StatementSelect:
		fmt.Println("This is where we should do a select.")
	default:
		fmt.Println("Execution operation not handled.")
	}
}
