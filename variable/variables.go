package variables

import (
	"github.com/nilspolek/Golaxy/router"
)

type Variable[T any] struct {
	value T
	r     *router.Router
}

func New[T any](r *router.Router) *Variable[T] {
	return &Variable[T]{
		r: r,
	}
}

func (v *Variable[T]) Set(val T) {
	v.value = val
	v.r.Render()
}

func (v *Variable[T]) Get() T {
	return v.value
}
