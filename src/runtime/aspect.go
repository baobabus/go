package runtime

import (
	"unsafe"
)

type AspectMap map[uintptr]interface{}

var aspm *_type = func() *_type {
	var m AspectMap = make(map[uintptr]interface{})
	var i interface{} = m
	e := efaceOf(&i)
	return e._type.typeOff(e._type.ptrToThis)
}()

func switchE2A(e *eface, t **_type, inter *interfacetype) {
	if e._type == aspm {
		ex := *(*AspectMap)(e.data)
		if v, ok := ex[uintptr(unsafe.Pointer(&inter.typ))]; ok {
			*e = *efaceOf(&v)
			if t != nil {
				*t = e._type
			}
		}
	}
}
