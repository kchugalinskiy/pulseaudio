package simple

// #include <pulse/sample.h>
import "C"
import (
	"fmt"
)

// SampleFormat is a basic sample characteristic. Each raw PCM sample may be in one of the supported formats.
// If you're using some codecs, you should encode/decode your audio data to PCM format before passing it to
// the pulseaudio.
// See https://www.freedesktop.org/software/pulseaudio/doxygen/sample_8h.html#a3c622fc51f4fc6ebfdcc7b454ac9c05f
type SampleFormat int

const (
	SampleFormatU8 SampleFormat = iota
	SampleFormatALAW
	SampleFormatULAW
	SampleFormatS16LE
	SampleFormatS16BE
	SampleFormatFloat32LE
	SampleFormatFloat32BE
	SampleFormatS32BE
	SampleFormatS32LE
	SampleFormatS24BE
	SampleFormatS24LE
	SampleFormatS2432BE
	SampleFormatS2432LE
	SampleFormatMax
	SampleFormatInvalid
)

func toSampleFormat(f SampleFormat) C.pa_sample_format_t {
	switch f {
	case SampleFormatU8:
		return C.PA_SAMPLE_U8
	case SampleFormatALAW:
		return C.PA_SAMPLE_ALAW
	case SampleFormatULAW:
		return C.PA_SAMPLE_ULAW
	case SampleFormatS16LE:
		return C.PA_SAMPLE_S16LE
	case SampleFormatS16BE:
		return C.PA_SAMPLE_S16BE
	case SampleFormatFloat32LE:
		return C.PA_SAMPLE_FLOAT32LE
	case SampleFormatFloat32BE:
		return C.PA_SAMPLE_FLOAT32BE
	case SampleFormatS32BE:
		return C.PA_SAMPLE_S32BE
	case SampleFormatS32LE:
		return C.PA_SAMPLE_S32LE
	case SampleFormatS24BE:
		return C.PA_SAMPLE_S24BE
	case SampleFormatS24LE:
		return C.PA_SAMPLE_S24LE
	case SampleFormatS2432BE:
		return C.PA_SAMPLE_S24_32BE
	case SampleFormatS2432LE:
		return C.PA_SAMPLE_S24_32LE
	case SampleFormatMax:
		return C.PA_SAMPLE_MAX
	case SampleFormatInvalid:
		return C.PA_SAMPLE_INVALID
	}
	panic(fmt.Sprintf("unsupported sample format %v", f))
}
