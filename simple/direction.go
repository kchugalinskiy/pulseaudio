package simple

// #include <pulse/def.h>
import "C"
import "fmt"

// StreamDirection is a basic flow of the stream. A stream may be a playback (write-only), record (read-only) or
// upload (this is similar to a playback, but has some caching mechanism)
// See https://www.freedesktop.org/software/pulseaudio/doxygen/def_8h.html#a7311932553b3f7962a092906576bc347
type StreamDirection int

const (
	StreamDirectionNo StreamDirection = iota
	StreamDirectionPlayback
	StreamDirectionRecord
	StreamDirectionUpload
)

func toDirection(d StreamDirection) C.pa_stream_direction_t {
	switch d {
	case StreamDirectionNo:
		return C.PA_STREAM_NODIRECTION
	case StreamDirectionPlayback:
		return C.PA_STREAM_PLAYBACK
	case StreamDirectionRecord:
		return C.PA_STREAM_RECORD
	case StreamDirectionUpload:
		return C.PA_STREAM_UPLOAD
	}
	panic(fmt.Sprintf("bad stream direction: %v", d))
}
