package common

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed assets/tags.xml
var defaultXMLData []byte

func GetXMLData(tagsFile string) []byte {
	if tagsFile != "" {
		xmlData, err := os.ReadFile(tagsFile)
		if err != nil {
			fmt.Println("Error reading tags file:", err)
			os.Exit(1)
		}
		return xmlData
	}
	return defaultXMLData
}
