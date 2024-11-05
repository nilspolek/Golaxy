package variable

import "github.com/nilspolek/Golaxy/router"

type Integer struct {
	val int
	r   router.Router
}

func NewInt(r router.Router) *Integer {
	return &Integer{
		r: r,
	}
}

func (v *Integer) Set(val int) {
	v.val = val
	v.r.Render()
}

func (v *Integer) Get() int {
	return v.val
}
