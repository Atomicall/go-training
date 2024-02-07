package subtask1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	s "strings"
	u "unicode"
)

func CheckPalindrome(str *string) bool {
	runes := []rune(s.ToLower(*str))
	head, tail := 0, len(runes)-1
	res := true

	for head < tail {

		for !isAlpanumeric(runes[head]) {
			head++
		}

		for !isAlpanumeric(runes[tail]) {
			tail--
		}

		if runes[head] != runes[tail] {
			res = false
			break
		}
		head++
		tail--
	}
	return res
}

// usage example
func main() {
	stringToCheck := readString()
	fmt.Println("String to check:", stringToCheck)
	fmt.Println("Is palindrome:", CheckPalindrome(&stringToCheck))
}

func isAlpanumeric(r rune) bool {
	return u.IsLetter(r) || u.IsDigit(r)
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
