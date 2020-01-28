package loop

// #include <pulse/mainloop.h>
// #cgo LDFLAGS: -lpulse-mainloop-glib -lpulse
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/kchugalinskiy/pulseaudio/async/api"

	"github.com/kchugalinskiy/pulseaudio/errs"
)

// Loop struct implements async main loop API
// See https://www.freedesktop.org/software/pulseaudio/doxygen/mainloop.html
type Loop struct {
	loop *C.pa_mainloop
}

func New() (*Loop, error) {
	return &Loop{loop: C.pa_mainloop_new()}, nil
}

func (l *Loop) Iterate(block int) (int, error) {
	var code C.int
	processed := C.pa_mainloop_iterate(l.loop, C.int(block), &code)
	if processed < 0 {
		return int(processed), fmt.Errorf("running loop: %v", errs.FromCode(int(code)))
	}

	return int(processed), nil
}

func (l *Loop) Run() error {
	var code C.int
	if C.pa_mainloop_run(l.loop, &code) < 0 {
		return fmt.Errorf("running loop: %v", errs.FromCode(int(code)))
	}

	return nil
}

func (l *Loop) Close() {
	C.pa_mainloop_free(l.loop)
	l.loop = nil
}

func (l *Loop) GetAPI() *api.MainloopAPI {
	return api.New((unsafe.Pointer(C.pa_mainloop_get_api(l.loop))))
}
