package channel

// #include <pulse/stream.h>
// #cgo LDFLAGS: -lpulse
import "C"
import "unsafe"

type Map map[int]ChannelPosition

func (m *Map) Marshal() unsafe.Pointer {
	var cmap *C.pa_channel_map
	if len(*m) != 0 {
		cmap = &C.pa_channel_map{}
		cmap.channels = C.uint8_t(len(*m))
		for k, v := range *m {
			cmap._map[k] = (C.pa_channel_position_t)(Marshal(v))
		}
	}
	return unsafe.Pointer(cmap)
}
