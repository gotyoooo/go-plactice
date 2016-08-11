package bench_test

import "testing"

// =============================================================
// メモリのアロケーション回数を減らすとパフォーマンスが大幅に改善する
// =============================================================

func BenchmarkMemAllocOndemand(b *testing.B) {
    n := 10
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        s := make([]string, 0) // メモリには割り当てていない
        for j := 0; j < n; j++ {
            s = append(s, "alice") // ここで追加される毎にメモリアロケーションが動く
        }
    }
}

func BenchmarkMemAllocAllBeforeUsing(b *testing.B) {
    n := 10
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        s := make([]string, 0, n) // 先に必要な分メモリに割り当てている
        for j := 0; j < n; j++ {
            s = append(s, "alice") // 割り当てたものに突っ込んでいる
        }
    }
}
