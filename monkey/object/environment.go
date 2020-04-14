package object

type Environment struct {
	store map[string]Object
	outer *Environment // 拡張元の環境
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	// 内部スコープで見つからない && 外部に環境がある場合
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}