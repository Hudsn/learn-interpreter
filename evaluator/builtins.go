package evaluator

import (
	"fmt"

	"github.com/hudsn/learn-interpreter/object"
)

var builtins = map[string]*object.Builtin{
	"puts": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d. want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				if len(arg.Elements) == 0 {
					return NULL
				}
				return arg.Elements[0]
			default:
				return newError("argument to `first` must be ARRAY, got %s", arg.Type())
			}
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d. want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				if len(arg.Elements) <= 0 {
					return NULL
				}
				return arg.Elements[len(arg.Elements)-1]
			default:
				return newError("argument to `last` must be ARRAY, got %s", arg.Type())
			}
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d. want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				length := len(arg.Elements)
				if length <= 0 {
					return NULL
				}
				newElements := make([]object.Object, length-1)
				copy(newElements, arg.Elements[1:])
				return &object.Array{Elements: newElements}
			default:
				return newError("argument to `rest` must be ARRAY, got %s", arg.Type())
			}
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d. want=2", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				length := len(arg.Elements)
				newElements := make([]object.Object, length+1)
				copy(newElements, arg.Elements)
				newElements[length] = args[1]
				return &object.Array{Elements: newElements}
			default:
				return newError("argument to `push` must be ARRAY, got %s", arg.Type())
			}
		},
	},
}
