package simple

// #include <pulse/def.h>
import "C"

// see https://www.freedesktop.org/software/pulseaudio/doxygen/structpa__buffer__attr.html
type BufferAttribute struct {
	MaxLength    uint32
	TargetLength uint32
	PreBuffer    uint32
	MinReq       uint32
	FragSize     uint32
}

func toAttr(attr *BufferAttribute) *C.pa_buffer_attr {
	if attr == nil {
		return nil
	}
	var cattr C.pa_buffer_attr
	cattr.maxlength = C.uint32_t(attr.MaxLength)
	cattr.tlength = C.uint32_t(attr.TargetLength)
	cattr.prebuf = C.uint32_t(attr.PreBuffer)
	cattr.minreq = C.uint32_t(attr.MinReq)
	cattr.fragsize = C.uint32_t(attr.FragSize)
	return &cattr
}
