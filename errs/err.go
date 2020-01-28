package errs

// #cgo LDFLAGS: -lpulse
// #include <pulse/error.h>
import "C"
import "fmt"

func FromCode(code int) error {
	return fmt.Errorf(C.GoString(C.pa_strerror(C.int(code))))
}
