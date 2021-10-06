package profiling_demo

import "testing"

func BenchmarkInc(b *testing.B) {
	var av AtomicVariable

	b.SetParallelism(16)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			av.Inc()
			av.Get()
		}
	})
}
