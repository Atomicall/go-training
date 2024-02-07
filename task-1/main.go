package main

import (
	"bufio"
	"log"
	"os"
	s "strings"
	t2 "subtask2/subtask2"
)

func main() {
	t2.Smain()
}

func readString() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		result, err := reader.ReadString('\n')
		if err != nil {
			log.Print(err)
			continue
		}
		return s.TrimSpace(result)
	}
}
