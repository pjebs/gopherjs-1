// +build js
// +build go1.15

package time

import "github.com/gopherjs/gopherjs/js"

type runtimeTimer struct {
	i       int32
	when    int64
	period  int64
	f       func(interface{}, uintptr)
	arg     interface{}
	timeout *js.Object
	active  bool
	seq     uintptr
}

func resetTimer(r *runtimeTimer, w int64) bool {
	active := stopTimer(r)
	r.when = w
	startTimer(r)
	return active
}
