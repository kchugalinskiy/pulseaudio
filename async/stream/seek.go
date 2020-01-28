package stream

// #include <pulse/def.h>
import "C"
import "fmt"

type SeekMode int

const (
	SeekModeRelative SeekMode = iota
	SeekModeAbsolute
	SeekModeRelativeOnRead
	SeekModeRelativeEnd
)

func (s SeekMode) Marshal() int {
	switch s {
	case SeekModeRelative:
		return C.PA_SEEK_RELATIVE
	case SeekModeAbsolute:
		return C.PA_SEEK_ABSOLUTE
	case SeekModeRelativeOnRead:
		return C.PA_SEEK_RELATIVE_ON_READ
	case SeekModeRelativeEnd:
		return C.PA_SEEK_RELATIVE_END
	}
	panic(fmt.Sprintf("unsupported seek mode: %v", s))
}
