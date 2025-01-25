package main

import (
	"fmt"
	"os"

	"example/gocli/internal/commands"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "gocli"}

	rootCmd.AddCommand(commands.NewDumpCommand())
	rootCmd.AddCommand(commands.NewEncodeCommand())
	rootCmd.AddCommand(commands.NewDecodeCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
