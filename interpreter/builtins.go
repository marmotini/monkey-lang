package interpreter

import (
	"fmt"

	"github.com/marmotini/ngiri-lang/object"
)

var builtins = map[string]*object.BuiltIn{
	"len": &object.BuiltIn{
		FN: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"puts": &object.BuiltIn{
		FN: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},
}
