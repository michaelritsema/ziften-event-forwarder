package util

import (
	"bufio"
	"os"
	"path/filepath"
)

func ExampleMessage() []string {
	cd, _ := os.Getwd()
	fullpath := filepath.Join(cd, "../testdata", "agent_messages.txt")
	//fmt.Println(fullpath)
	f, _ := os.Open(fullpath)
	reader := bufio.NewReader(f)
	var lines []string

	for {
		bytes, _, e := reader.ReadLine()
		if e != nil {
			//fmt.Println(e)
			break
		}
		l := string(bytes)
		lines = append(lines, l)
	}

	return lines
}
