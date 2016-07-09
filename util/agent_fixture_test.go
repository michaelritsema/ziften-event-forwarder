package util

import (
	"testing"
)

func TestPrintMessages(t *testing.T) {
	ExampleMessage("../agent_messages.txt")
	t.Log("...")
}
