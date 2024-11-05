package variable

import (
	"github.com/nilspolek/Golaxy/router"
)

type Variable struct {
	value interface{}
	r     *router.Router
}

func New(r *router.Router) *Variable {
	return &Variable{
		r: r,
	}
}

func (v *Variable) Set(val interface{}) {
	v.value = val
	v.r.Render()
}

func (v *Variable) Get() interface{} {
	return v.value
}
