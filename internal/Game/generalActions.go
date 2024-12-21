package Game

import (
	"fmt"
	"reflect"
	"strings"
)

func (a *Game) NewSingleAction(actionName, methodName string, iface any, args ...string) func(actioargs ...any) {
	return func(actionArgs ...any) {
		if notEnoughArgs(2, args...) {
			Respond(actionName, " what?")
			return
		}

		target := strings.Join(args[1:], " ")
		o := a.findObjsInGame(target)

		switch len(o) {
      case 1:
        ifaceType := reflect.TypeOf(iface).Elem()
        objType := reflect.TypeOf(o[0])

        if !objType.Implements(ifaceType) {
          Respond("can't " + actionName + " the ", o[0].getDescription())
          return
        }

        objValue := reflect.ValueOf(o[0])
        method := objValue.MethodByName(methodName)
        if !method.IsValid() {
          panic(fmt.Sprintf("method \"%s\" not found", methodName))
        }

        methodArgs := make([]reflect.Value, len(actionArgs))
        for i,v := range actionArgs {
          methodArgs[i] = reflect.ValueOf(v)
        }

        method.Call(methodArgs)

      case 0:
        Respond("can't find any \"", target, "\"s")

      default:
        str := make([]string, 0)
        for _, i := range o {
          str = append(str, i.getDescription())
        }
        Respond(strings.Join(str, " or "), "?")
		}
	}
}
