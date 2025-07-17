package object

type Environment struct {
	store map[string]Object
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

func (e *Environment) Get(name string) (Object, bool) {
	o, ok := e.store[name]
	return o, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
