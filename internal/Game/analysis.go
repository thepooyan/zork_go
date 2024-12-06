package Game

import "strings"


func analyzeResponse(response string) {

  words := strings.Fields(response);
  
  println("len", len(words))
}
