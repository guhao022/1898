package utils

import (

	"testing"

	"time"
	"fmt"
)

func Test_RandomCreateBytes(t *testing.T) {
	start := time.Now()
	for i:=0;i<10000;i++ {
		RandomCreateBytes(6)
	}

	fmt.Println(time.Since(start))
}

func Benchmark_RandomCreateBytes(b *testing.B) {
	RandomCreateBytes(100000000)
}
