package main

import (
	"github.com/evercyan/mmmoney/command/loan"
	"github.com/evercyan/mmmoney/command/tax"
	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:     "mmmoney",
		Short:   "mmmoney: make more money",
		Version: "v0.0.1",
	}
	root.AddCommand(loan.Command)
	root.AddCommand(tax.Command)
	root.Execute()
}
