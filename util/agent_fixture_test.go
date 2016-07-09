package util

import (
	"testing"
)

func TestPrintMessages(t *testing.T) {
	for _, line := range ExampleMessage() {
		t.Log(line)
	}

}
