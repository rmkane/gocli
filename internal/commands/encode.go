package commands

import (
	"example/gocli/internal/common"
	"fmt"

	"github.com/spf13/cobra"
)

func NewEncodeCommand() *cobra.Command {
	var tagsFile string

	cmd := &cobra.Command{
		Use:   "encode [args]",
		Short: "Encode command",
		Run: func(cmd *cobra.Command, args []string) {
			xmlData := common.GetXMLData(tagsFile)
			fmt.Println("Encode command executed with args:", args)
			fmt.Println("XML Data:", string(xmlData))
		},
	}

	cmd.Flags().StringVarP(&tagsFile, "tags", "t", "", "Path to XML tags file")

	return cmd
}
