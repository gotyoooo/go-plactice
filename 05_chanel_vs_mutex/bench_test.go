package bench_test

import (
  "sync"
  "testing"
)

// =============================================================
// channel を使った排他制御より sync.Mutex や sync.RWMutex を使ったほうが速い
// =============================================================

func BenchmarkExclusiveWithChannel(b *testing.B) {
  c := make(chan struct{}, 1)
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    c <- struct{}{}
    // do something.
    <-c
  }
}

func BenchmarkExclusiveWithMutex(b *testing.B) {
  mu := new(sync.Mutex)
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    mu.Lock()
    // do something.
    mu.Unlock()
  }
}

// =============================================================
// 同期処理も sync.WaitGroup を使ったほうが速い
// =============================================================

func BenchmarkSyncWithChannel(b *testing.B) {
  n := 10
  c := make(chan struct{}, n)
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    for j := 0; j < n; j++ {
      go func() {
        // do something.
        c <- struct{}{}
      }()
    }
    for j := 0; j < n; j++ {
      <-c
    }
  }
}

func BenchmarkSyncWithWaitGroup(b *testing.B) {
  n := 10
  var wg sync.WaitGroup
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    wg.Add(n)
    for j := 0; j < n; j++ {
      go func() {
        // do something.
        wg.Done()
      }()
    }
    wg.Wait()
  }
}
