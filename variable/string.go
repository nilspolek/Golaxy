package variable

import "github.com/nilspolek/Golaxy/router"

type String struct {
	val string
	r   router.Router
}

func NewString(r router.Router) *String {
	return &String{
		r: r,
	}
}

func (v *String) Set(val string) {
	v.val = val
	v.r.Render()
}

func (v *String) Get() string {
	return v.val
}
