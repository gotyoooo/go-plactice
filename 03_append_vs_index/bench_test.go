package bench_test

import "testing"

// =============================================================
// 要素数が事前にわかっている場合には append を使わない
// =============================================================

func BenchmarkFillSliceByAppend(b *testing.B) {
  n := 100
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    s := make([]int, 0, n)
    for j := 0; j < n; j++ {
      s = append(s, j) // append で追加
    }
  }
}

func BenchmarkFillSliceByIndex(b *testing.B) {
  n := 100
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    s := make([]int, n)
    for j := 0; j < n; j++ {
      s[j] = j // インデックスを使ってスライスに代入
    }
  }
}
