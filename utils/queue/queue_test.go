package queue

import (
	"testing"
	"time"
)

func Test_Queue(t *testing.T) {
	tt := time.Now().Unix()

	aa := time.Unix(tt, 0)

	bb := aa.Format("20060102030405")

	println(bb)
}
