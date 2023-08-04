package bench

import "testing"

func BenchmarkSelect(b *testing.B) {
	InitServerIndex()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Select()
	}
}

// parallel 平行的，同时发生的
func BenchmarkSelectParallel(b *testing.B) {
	InitServerIndex()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Select()
		}
	})
}
