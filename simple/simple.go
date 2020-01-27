package simple

// #include <pulse/simple.h>
// #include <pulse/error.h>
// #cgo LDFLAGS: -lpulse-simple -lpulse
import "C"
import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

// Simple is a go-wrapper over pa_simple structure. See
// https://www.freedesktop.org/software/pulseaudio/doxygen/simple.html documetation to get familiar with
// pulseaudio simple API.
type Simple struct {
	simple *C.pa_simple
}

// New simple API client. See Simple API documentation for the details.
func New(server, name string, dir StreamDirection, dev, streamName string, ss *SampleSpec, m map[int]ChannelPosition, attr *BufferAttribute) (*Simple, error) {
	var code C.int
	var s Simple
	var cmap *C.pa_channel_map
	if len(m) != 0 {
		cmap = &C.pa_channel_map{}
		cmap.channels = C.uint8_t(len(m))
		for k, v := range m {
			cmap._map[k] = toChannel(v)
		}
	}

	cserver := C.CString(server)
	if server == "" {
		cserver = nil
	}
	cname := C.CString(name)
	if name == "" {
		cname = nil
	}
	cdev := C.CString(dev)
	if dev == "" {
		cdev = nil
	}
	cstream := C.CString(streamName)
	if streamName == "" {
		cstream = nil
	}
	cattr := toAttr(attr)
	cdir := toDirection(dir)
	css := toSampleSpec(ss)

	s.simple = C.pa_simple_new(cserver, cname, cdir, cdev, cstream, css, cmap, cattr, &code)
	if code != 0 {
		return nil, fmt.Errorf("bad simple ctor code: %v", errorFromCode(code))
	}
	return &s, nil
}

func errorFromCode(code C.int) error {
	return fmt.Errorf(C.GoString(C.pa_strerror(code)))
}

// Flush the buffers.
func (s *Simple) Flush() error {
	var code C.int
	C.pa_simple_flush(s.simple, &code)
	if code != 0 {
		return fmt.Errorf("flushing buffers: %v", errorFromCode(code))
	}

	return nil
}

// Drain the buffers.
func (s *Simple) Drain() error {
	var code C.int
	C.pa_simple_drain(s.simple, &code)
	if code != 0 {
		return fmt.Errorf("draining buffers: %v", errorFromCode(code))
	}

	return nil
}

// Latency returns current stream latency.
func (s *Simple) Latency() (time.Duration, error) {
	var code C.int
	latency := C.pa_simple_get_latency(s.simple, &code)
	if code != 0 {
		return 0, fmt.Errorf("draining buffers: %v", errorFromCode(code))
	}

	return time.Duration(latency) * time.Microsecond, nil
}

// Read8 blocks execution (!) until n bytes of data are read. Data read is encoded in specified on stream creation
// PCM format.
func (s *Simple) Read8(n int) ([]byte, error) {
	b := C.malloc(C.ulong(n))
	var code C.int
	C.pa_simple_read(s.simple, b, C.size_t(n), &code)
	if code != 0 {
		return nil, fmt.Errorf("reading buffer: %v", errorFromCode(code))
	}

	return C.GoBytes(b, C.int(n)), nil
}

// Read16 blocks execution (!) until n bytes of data are read. Data read is encoded in specified on stream creation
// PCM format.
func (s *Simple) Read16(n int) ([]int16, error) {
	b := C.malloc(C.ulong(n))
	var code C.int
	C.pa_simple_read(s.simple, b, C.size_t(n), &code)
	if code != 0 {
		return nil, fmt.Errorf("reading buffer: %v", errorFromCode(code))
	}

	return int16fromByte(C.GoBytes(b, C.int(n))), nil
}

// Read32 blocks execution (!) until n bytes of data are read. Data read is encoded in specified on stream creation
// PCM format.
func (s *Simple) Read32(n int) ([]int32, error) {
	b := C.malloc(C.ulong(n))
	var code C.int
	C.pa_simple_read(s.simple, b, C.size_t(n), &code)
	if code != 0 {
		return nil, fmt.Errorf("reading buffer: %v", errorFromCode(code))
	}

	return int32fromByte(C.GoBytes(b, C.int(n))), nil
}

func int16fromByte(src []byte) []int16 {
	var t []int16
	n := len(src) / 2
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&t))
	header.Data = reflect.ValueOf(src).Pointer()
	header.Len = n
	header.Cap = n
	return *(*[]int16)(unsafe.Pointer(&header))
}

func int32fromByte(src []byte) []int32 {
	var t []int32
	n := len(src) / 4
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&t))
	header.Data = reflect.ValueOf(src).Pointer()
	header.Len = n
	header.Cap = n
	return *(*[]int32)(unsafe.Pointer(&header))
}

func newSlice(t reflect.Type, data unsafe.Pointer, n int) interface{} {
	val := reflect.MakeSlice(t, n, n)
	s := (*reflect.SliceHeader)(unsafe.Pointer(val.Pointer()))
	s.Data = uintptr(data)
	s.Len = n
	s.Cap = n
	return val.Interface()
}

// Write8 blocks execution until b PCM data in the specified during the creation PCM format is being written to the
// write buffer.
func (s *Simple) Write8(b []byte) error {
	var code C.int
	C.pa_simple_write(s.simple, C.CBytes(b[:]), C.size_t(len(b)), &code)
	if code != 0 {
		return fmt.Errorf("writing buffer: %v", errorFromCode(code))
	}

	return nil
}

func int16toByte(src []int16) []byte {
	var t []byte
	n := len(src) * 2
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&t))
	header.Data = reflect.ValueOf(src).Pointer()
	header.Len = n
	header.Cap = n
	return *(*[]byte)(unsafe.Pointer(&header))
}

func int32toByte(src []int32) []byte {
	var t []byte
	n := len(src) * 4
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&t))
	header.Data = reflect.ValueOf(src).Pointer()
	header.Len = n
	header.Cap = n
	return *(*[]byte)(unsafe.Pointer(&header))
}

// Write16 blocks execution until b PCM data in the specified during the creation PCM format is being written to the
// write buffer.
func (s *Simple) Write16(b []int16) error {
	var code C.int
	C.pa_simple_write(s.simple, C.CBytes(int16toByte(b)), C.size_t(len(b)), &code)
	if code != 0 {
		return fmt.Errorf("writing buffer: %v", errorFromCode(code))
	}

	return nil
}

// Write32 blocks execution until b PCM data in the specified during the creation PCM format is being written to the
// write buffer.
func (s *Simple) Write32(b []int32) error {
	var code C.int
	C.pa_simple_write(s.simple, C.CBytes(int32toByte(b)), C.size_t(len(b)), &code)
	if code != 0 {
		return fmt.Errorf("writing buffer: %v", errorFromCode(code))
	}

	return nil
}

// Close cleans up pulseaudio client.
func (s *Simple) Close() {
	C.pa_simple_free(s.simple)
}
