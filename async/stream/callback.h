#include <stdlib.h>
#include <pulse/stream.h>
#include <pulse/def.h>

extern pa_stream_success_cb_t onSuccess;
extern pa_stream_request_cb_t onRequest;
extern pa_stream_notify_cb_t onNotify;
extern pa_free_cb_t onFree;