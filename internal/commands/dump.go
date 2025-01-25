package commands

import (
	"example/gocli/internal/common"
	"fmt"

	"github.com/spf13/cobra"
)

func NewDumpCommand() *cobra.Command {
	var verbose bool
	var tagsFile string

	cmd := &cobra.Command{
		Use:   "dump [args]",
		Short: "Dump command",
		Run: func(cmd *cobra.Command, args []string) {
			xmlData := common.GetXMLData(tagsFile)
			if verbose {
				fmt.Println("Verbose mode enabled")
			}
			fmt.Println("Dump command executed with args:", args)
			fmt.Println("XML Data:", string(xmlData))
		},
	}

	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose mode")
	cmd.Flags().StringVarP(&tagsFile, "tags", "t", "", "Path to XML tags file")

	return cmd
}
