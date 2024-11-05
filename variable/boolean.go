package variable

import "github.com/nilspolek/Golaxy/router"

type Bool struct {
	val bool
	r   router.Router
}

func NewBool(r router.Router) *Bool {
	return &Bool{
		r: r,
	}
}

func (v *Bool) Set(val bool) {
	v.val = val
	v.r.Render()
}

func (v *Bool) Get() bool {
	return v.val
}
