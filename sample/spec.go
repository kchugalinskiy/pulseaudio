package sample

// #include <pulse/sample.h>
import "C"
import "unsafe"

// Spec stores information about single sample item, like the number of the channels or its format.
// see https://www.freedesktop.org/software/pulseaudio/doxygen/structpa__sample__spec.html
type Spec struct {
	Format   SampleFormat
	Rate     uint32
	Channels uint8
}

func Marshal(ss *Spec) unsafe.Pointer {
	if ss == nil {
		return nil
	}
	var css C.pa_sample_spec
	css.format = toSampleFormat(ss.Format)
	css.rate = C.uint32_t(ss.Rate)
	css.channels = C.uint8_t(ss.Channels)
	return unsafe.Pointer(&css)
}
