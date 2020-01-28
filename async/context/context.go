package context

// #include <pulse/context.h>
// #cgo LDFLAGS: -lpulse
import "C"
import (
	"unsafe"

	"github.com/kchugalinskiy/pulseaudio/async/api"
)

type Context struct {
	ctx *C.pa_context
}

func New(a api.MainloopAPI, name string) *Context {
	return &Context{ctx: C.pa_context_new((*C.pa_mainloop_api)(a.Get()), C.CString(name))}
}

func (c *Context) Close() {
	C.pa_context_disconnect(c.ctx)
	c.ctx = nil
}

func (c *Context) Get() unsafe.Pointer {
	return unsafe.Pointer(c.ctx)
}
