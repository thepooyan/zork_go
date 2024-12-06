package Game

import "fmt"

func Outputln(write string) {
  fmt.Println(write);
}

func Output(write string) {
  fmt.Print(write)
}

func Input() string {
  var response string
  fmt.Scan(&response)
  return response
}
