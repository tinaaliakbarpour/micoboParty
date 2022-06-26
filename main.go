package main

import (
	"fmt"
	"html"
	"micobianParty/cmd"

	"github.com/logrusorgru/aurora"
)

// root execute command with cobra
func main() {
	if err := cmd.Runner.RootCmd().Execute(); err != nil {
		fmt.Printf("\n %v Failed to run command: %v %v\n\n ", aurora.White(html.UnescapeString("&#x274C;")), err, aurora.White(html.UnescapeString("&#x274C;")))
	}
}
