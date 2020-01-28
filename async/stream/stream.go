package stream

import "C"
import (
	"unsafe"

	"github.com/kchugalinskiy/pulseaudio/async/context"
	"github.com/kchugalinskiy/pulseaudio/channel"
	"github.com/kchugalinskiy/pulseaudio/sample"
)

// #include <pulse/stream.h>
// #include <pulse/def.h>
// #include <stdlib.h>
// #include "callback.h"
// #cgo LDFLAGS: -lpulse
import "C"

type Stream struct {
	s *C.pa_stream
}

func New(c *context.Context, name string, spec *sample.Spec, m channel.Map) *Stream {
	cc := (*C.pa_context)(c.Get())
	css := (*C.pa_sample_spec)(sample.Marshal(spec))
	var cm *C.pa_channel_map
	if m != nil {
		cm = (*C.pa_channel_map)(m.Marshal())
	}
	return &Stream{s: C.pa_stream_new(cc, C.CString(name), css, cm)}
}

func (s *Stream) Readable() int {
	return int(C.pa_stream_readable_size(s.s))
}

func (s *Stream) OnRead(cb OnRequest, userdata interface{}) {
	C.pa_stream_set_read_callback(s.s, C.onRequest, nil)
}

func (s *Stream) Read() []byte {
	var data unsafe.Pointer
	var size C.size_t
	C.pa_stream_peek(s.s, &data, &size)
	if size == 0 || data == nil {
		return nil
	}
	return C.GoBytes(data, C.int(size))
}

func (s *Stream) Writable() int {
	return int(C.pa_stream_writable_size(s.s))
}

func (s *Stream) OnWrite(cb OnRequest, userdata interface{}) {
	C.pa_stream_set_write_callback(s.s, C.onRequest, nil)
}

func (s *Stream) Write(b []byte, m SeekMode) (int, error) {
	var data unsafe.Pointer
	var nbytes C.size_t
	nbytes = C.ulong(len(b))
	C.pa_stream_begin_write(s.s, &data, &nbytes)
	if nbytes > len(b) {
		nbytes = len(b)
	}

	C.memcpy()

	C.pa_stream_write(
		s.s,
		data,
		nbytes,
		C.onFree,
		0,
		C.pa_seek_mode_t(m.Marshal()),
	)
	return int(nbytes), nil
}

func (s *Stream) OnOverflow(cb OnNotify, userdata interface{}) {
	C.pa_stream_set_overflow_callback(s.s, C.onNotify, nil)
}

func (s *Stream) OnUnderflow(cb OnNotify, userdata interface{}) {
	C.pa_stream_set_underflow_callback(s.s, C.onNotify, nil)
}

func (s *Stream) OnStart(cb OnRequest, userdata interface{}) {
	C.pa_stream_set_started_callback(s.s, C.onNotify, nil)
}

func (s *Stream) OnStateChanged(cb OnRequest, userdata interface{}) {
	C.pa_stream_set_state_callback(s.s, C.onNotify, nil)
}

func (s *Stream) OnSuspended(cb OnRequest, userdata interface{}) {
	C.pa_stream_set_suspended_callback(s.s, C.onNotify, nil)
}
