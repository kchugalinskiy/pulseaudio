package stream

// #include <pulse/stream.h>
// #include "callback.h"
import "C"
import "unsafe"

//int cb(const void *inputBuffer, void *outputBuffer, unsigned long frames, const PaStreamCallbackTimeInfo *timeInfo, PaStreamCallbackFlags statusFlags, void *userData) {
//type streamCallback(inputBuffer, outputBuffer unsafe.Pointer, frames C.ulong, timeInfo *C.PaStreamCallbackTimeInfo, statusFlags C.PaStreamCallbackFlags, userData unsafe.Pointer) {

//type onEvent func() (p *C.pa_stream, name *C.char, pl *C.pa_proplist, userdata unsafe.Pointer)
//type OnEvent func() (s *Stream, name string, pl *C.pa_proplist, userdata interface{})

//type onNotify func(p *C.pa_stream, userdata unsafe.Pointer)
type OnNotify func(s *Stream, userdata interface{})

//func (cb OnNotify) marshal() onNotify {
//	return func(p *C.pa_stream, userdata unsafe.Pointer) {
//		cb(&Stream{s: p}, (interface{})(uintptr(userdata)))
//	}
//}

//export NotifyCb
func NotifyCb(p *C.pa_stream, userdata unsafe.Pointer) {

}

//type onRequest func(p *C.pa_stream, nbytes C.size_t, userdata unsafe.Pointer)
type OnRequest func(s *Stream, nbytes int, userdata interface{})

//func (cb OnRequest) marshal() onRequest {
//	return func(p *C.pa_stream, nbytes C.size_t, userdata unsafe.Pointer) {
//		cb(&Stream{s: p}, int(nbytes), (interface{})(uintptr(userdata)))
//	}
//}

//export RequestCb
func RequestCb(p *C.pa_stream, nbytes C.size_t, userdata unsafe.Pointer) {

}

//type onSuccess func(p *C.pa_stream, succ C.int, userdata unsafe.Pointer)
type OnSuccess func(s *Stream, succ int, userdata interface{})

//func (cb OnSuccess) marshal() onSuccess {
//	return func(p *C.pa_stream, succ C.int, userdata unsafe.Pointer) {
//		cb(&Stream{s: p}, int(succ), (interface{})(uintptr(userdata)))
//	}
//}

//export SuccessCb
func SuccessCb(p *C.pa_stream, succ C.int, userdata unsafe.Pointer) {

}

//type onFree func(p unsafe.Pointer)
type OnFree func(interface{})

//
//func (cb OnFree) marshal() onFree {
//	return func(p unsafe.Pointer) {
//		cb((interface{})(uintptr(p)))
//	}
//}

//export FreeCb
func FreeCb(p unsafe.Pointer) {
	C.free(p)
}
