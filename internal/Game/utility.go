package Game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Outputln(write ...string) {
  join := strings.Join(write, "")
  fmt.Println(join)
}

func Output(write string) {
  fmt.Print(write)
}

func Input() string {
  reader := bufio.NewReader(os.Stdin)
	res, _ := reader.ReadString('\n')
  return res
}

func GetUserInput() string {
  Output(" => ")
  response := Input()
  return response
}

func fileExists(filename string) bool {
    _, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false // File does not exist
    }
    return err == nil // Return true if no error, false otherwise
}
