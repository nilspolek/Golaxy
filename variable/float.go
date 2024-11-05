package variable

import "github.com/nilspolek/Golaxy/router"

type Float struct {
	val float64
	r   router.Router
}

func NewFloat(r router.Router) *String {
	return &String{
		r: r,
	}
}

func (v *Float) Set(val float64) {
	v.val = val
	v.r.Render()
}

func (v *Float) Get() float64 {
	return v.val
}
