package bench_test

import (
  "bytes"
  "sync"
  "testing"
)

func BenchmarkGoroutine(b *testing.B) {
  in := "0xDeadBeef"
  buf := &bytes.Buffer{}
  for i := 0; i < b.N; i++ {
    buf.WriteString(in)
  }
  buf.Bytes()
}

var globalBuf = &bytes.Buffer{}
var globalMutex = sync.Mutex{}

func BenchmarkGlobalBuffer(b *testing.B) {
  in := "0xDeadBeef"
  globalMutex.Lock()
  defer globalMutex.Unlock()
  globalBuf.Reset()
  for i := 0; i < b.N; i++ {
    globalBuf.WriteString(in)
  }
  globalBuf.Bytes()
}

var globalPool = sync.Pool{
  New: func() interface{} {
    return &bytes.Buffer{}
  },
}

func BenchmarkGlobalPool(b *testing.B) {
  in := "0xDeadBeef"
  buf := globalPool.Get().(*bytes.Buffer)
  defer func() {
    buf.Reset()
    globalPool.Put(buf)
  }()
  for i := 0; i < b.N; i++ {
    buf.WriteString(in)
  }
  buf.Bytes()
}
