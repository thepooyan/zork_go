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
