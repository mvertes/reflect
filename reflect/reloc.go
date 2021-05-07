package reflect

import _ "unsafe"

//go:linkname resolveTypeOff reflect.resolveTypeOff
//go:linkname resolveNameOff reflect.resolveNameOff
//go:linkname resolveTextOff reflect.resolveTextOff
//go:linkname addReflectOff reflect.addReflectOff
//go:linkname typelinks reflect.typelinks
//go:linkname unsafe_New reflect.unsafe_New
//go:linkname typedmemmove reflect.typedmemmove
//go:linkname typedmemclr reflect.typedmemclr
//go:linkname memmove reflect.memmove
//go:linkname typedmemmovepartial reflect.typedmemmovepartial
//go:linkname chanlen reflect.chanlen
//go:linkname maplen reflect.maplen
//go:linkname ifaceE2I reflect.ifaceE2I
