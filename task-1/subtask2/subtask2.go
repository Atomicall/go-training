package subtask2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	s "strings"
)

func Smain() {
	stringToCheck := readString()
	fmt.Println("String to process:", stringToCheck)
}

func RemoveRepeated(strings []string) []string {
	return []string{}
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
