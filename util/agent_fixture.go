package util

import (
	"bufio"
	"os"
)

func ExampleMessage(path string) []string {
	f, _ := os.Open(path)
	reader := bufio.NewReader(f)
	reader.ReadString('\n')
	return nil
}
