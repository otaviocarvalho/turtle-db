package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("db > ")
        input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == ".exit" {
			os.Exit(0)
		} else {
			fmt.Printf("Unrecognized command `%s`.\n", input)
		}
	}
}
