package _map

import "testing"

// BenchmarkLookup measures the performance of looking up keys in a map
func BenchmarkLookup(b *testing.B) {
	// Create a large map
	m := make(map[int]int, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		m[i] = i
	}

	// Measure lookup performance
	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 0; i < b.N; i++ {
		_ = m[i%1_000_000]
	}
}

// BenchmarkInsertion measures the performance of inserting keys into a map
func BenchmarkInsertion(b *testing.B) {
	// Create a large map
	m := make(map[int]int, 1_000_000)

	// Measure insertion performance
	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 1_000_000; i < 2_000_000; i++ {
		m[i] = i
	}
}

// BenchmarkDeletion measures the performance of deleting keys from a map
func BenchmarkDeletion(b *testing.B) {
	// Create a large map
	m := make(map[int]int, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		m[i] = i
	}

	// Measure deletion performance
	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 0; i < b.N; i++ {
		delete(m, i)
	}
}
