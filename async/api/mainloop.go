package api

// #include <pulse/mainloop-api.h>
// #cgo LDFLAGS: -lpulse
import "C"
import "unsafe"

type MainloopAPI struct {
	api *C.pa_mainloop_api
}

func New(api unsafe.Pointer) *MainloopAPI {
	return &MainloopAPI{api: (*C.pa_mainloop_api)(api)}
}

func (m *MainloopAPI) Get() unsafe.Pointer {
	return unsafe.Pointer(m.api)
}
