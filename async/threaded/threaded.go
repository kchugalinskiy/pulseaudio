package threaded

// #include <pulse/thread-mainloop.h>
// #cgo LDFLAGS: -lpulse
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/kchugalinskiy/pulseaudio/async/api"
)

type MainLoop struct {
	loop *C.pa_threaded_mainloop
}

func New(name string) *MainLoop {
	m := &MainLoop{loop: C.pa_threaded_mainloop_new()}
	C.pa_threaded_mainloop_set_name(m.loop, C.CString(name))
	return m
}

func (m *MainLoop) Start() error {
	if code := C.pa_threaded_mainloop_start(m.loop); code < 0 {
		return fmt.Errorf("starting main loop: %v", code)
	}
	return nil
}

func (m *MainLoop) Stop() {
	C.pa_threaded_mainloop_stop(m.loop)
}

func (m *MainLoop) Lock() {
	C.pa_threaded_mainloop_lock(m.loop)
}

func (m *MainLoop) Unlock() {
	C.pa_threaded_mainloop_unlock(m.loop)
}

func (m *MainLoop) Accept() {
	C.pa_threaded_mainloop_accept(m.loop)
}

func (m *MainLoop) Wait() {
	C.pa_threaded_mainloop_wait(m.loop)
}

func (m *MainLoop) Close() {
	C.pa_threaded_mainloop_free(m.loop)
	m.loop = nil
}

func (m *MainLoop) GetAPI() *api.MainloopAPI {
	return api.New((unsafe.Pointer(C.pa_threaded_mainloop_get_api(m.loop))))
}
