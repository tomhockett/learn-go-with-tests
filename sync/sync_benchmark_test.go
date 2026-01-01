package sync

import (
	"testing"
)

func BenchmarkCounterAtomic(b *testing.B) {
	counter := Counter{}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			counter.Inc()
		}
	})
}

func BenchmarkCounterAtomicRead(b *testing.B) {
	counter := Counter{}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			counter.Value()
		}
	})
}

func BenchmarkCounterAtomicMixed(b *testing.B) {
	counter := Counter{}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if b.N%10 == 0 {
				counter.Value() // 10% reads
			} else {
				counter.Inc() // 90% writes
			}
		}
	})
}

func BenchmarkCounterSequential(b *testing.B) {
	counter := Counter{}

	for i := 0; i < b.N; i++ {
		counter.Inc()
	}
}
