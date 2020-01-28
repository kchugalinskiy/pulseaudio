#include "_cgo_export.h"
#include "callback.h"

void wrap_successCb(pa_stream *s, int success, void *userdata) {
    SuccessCb(s, success, userdata);
}
pa_stream_success_cb_t onSuccess = wrap_successCb;

void wrap_requestCb(pa_stream *p, size_t nbytes, void *userdata) {
    RequestCb(p, nbytes, userdata);
}
pa_stream_request_cb_t onRequest = wrap_requestCb;

void wrap_notifyCb(pa_stream *p, void *userdata) {
    NotifyCb(p, userdata);
}
pa_stream_notify_cb_t onNotify = wrap_notifyCb;

void wrap_freeCb(void *p) {
    FreeCb(p);
}
pa_free_cb_t onFree = wrap_freeCb;
