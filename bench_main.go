package main

import (
	"fmt"
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Add(1, 2)
	}
}

func main() {
	r := testing.Benchmark(BenchmarkAdd)
	fmt.Println(r.String())
	fmt.Println(r.MemString())
}

//go run bench_main.go
//1000000000               0.2327 ns/op
//       0 B/op          0 allocs/op
//PS C:\Users\tla\GolandProjects\Learning_EM>
