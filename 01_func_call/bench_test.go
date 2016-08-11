package bench_test

import "testing"

// =============================================================
// goは関数やメソッドの呼び出しが遅い
// =============================================================

func BenchmarkUseFunction(b *testing.B) {
  sum := 0
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    sum = add(sum, 100)
  }
}

func BenchmarkNotUseFunction(b *testing.B) {
  sum := 0
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    sum = sum + 100
  }
}

func add(x int, y int) int {
  return x + y
}
