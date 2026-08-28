package exectime

import (
	"fmt"
	"time"
)

func Benchmark(f func()) {
	t1 := time.Now()
	f()
	execTime := time.Since(t1)
	fmt.Printf("exec time : %v\n", execTime)
}
