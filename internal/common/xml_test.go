package common

import (
	"testing"
)

func TestGetXMLData(t *testing.T) {
	data := GetXMLData("")
	if len(data) == 0 {
		t.Error("Expected default XML data, got empty data")
	}
}
