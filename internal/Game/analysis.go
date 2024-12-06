package Game

import (
	"reflect"
	"strings"
)


func analyzeResponse(response string) {

  words := strings.Fields(response);
  
  println(isViableAction(words[0]))
}

func isViableAction(str string) bool {
  val := reflect.ValueOf(Action{})
  method := val.MethodByName(str)
  return method.IsValid()
}
