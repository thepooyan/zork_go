package Game

import "fmt"

func StartGame() {
  fmt.Println("starting game...")
  loop()
}

func loop() {
  for {
    res := getUserInput("Hi!")
    analyzeResponse(res);
  }
}

func getUserInput(prompt string) string {
  fmt.Println(prompt)
  fmt.Print(" => ")
  var response string
  fmt.Scan(&response)
  return response
}
