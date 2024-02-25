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

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("db > ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input[0] == '.' {
			switch doMetaCommand(input) {
			case MetaCommandSuccess:
				continue
			case MetaCommandUnrecognized:
				fmt.Printf("Unrecognized command '%s'.\n", input)
				continue
			}
		}

		var statement Statement
		switch prepareStatement(input, &statement) {
		case PrepareSuccessStatement:
			break
		case PrepareUnrecognizedStatement:
			fmt.Printf("Unrecognized keyword at the start of '%s'.\n", input)
			continue
		}

		executeStatement(&statement)
		fmt.Printf("Executed.\n")
	}
}

func doMetaCommand(input string) MetaCommandResult {
	if input == ".exit" {
		os.Exit(int(ExitSuccess))
	}
	return MetaCommandUnrecognized
}

func prepareStatement(input string, statement *Statement) PrepareResultStatement {
	if strings.HasPrefix(input, "insert") {
		statement.Type = StatementInsert
		return PrepareSuccessStatement
	} else if strings.HasPrefix(input, "select") {
		statement.Type = StatementSelect
		return PrepareSuccessStatement
	}

	return PrepareUnrecognizedStatement
}

func executeStatement(statement *Statement) {
	switch statement.Type {
	case StatementInsert:
		fmt.Println("This is where we should do an insert.")
		break
	case StatementSelect:
		fmt.Println("This is where we should do a select.")
	}
}
