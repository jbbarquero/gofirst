package concurrency_patterns

import (
  "bytes"
  "io"
  "os"
  "strings"
  "testing"
)

/*
* go test -run=TestBarrier/Correct_endpoints -v
* go test -run=TestBarrier/One_endpoint_incorrect -v
* go test -run=TestBarrier/Very_short_timeout -v
* go test -v
*/

func TestBarrier(t *testing.T)  {
  t.Run("Correct endpoints", func(t *testing.T)  {
    endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/user-agent"}
    result := captureBarrierOutput(endpoints...)
    if !strings.Contains(result, "Accept-Encoding") || !strings.Contains(result, "user-agent") {
      t.Fail()
    }
    t.Log(result)
  })

  t.Run("One endpoint incorrect", func(t *testing.T)  {
    endpoints := []string{"http://malformed-url", "http://httpbin.org/user-agent"}
    result := captureBarrierOutput(endpoints...)
    if !strings.Contains(result, "ERROR") {
      t.Fail()
    }
    t.Log(result)
  })

  t.Run("Very short timeout", func(t *testing.T)  {
    endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/user-agent"}
    timeoutMiliseconds = 1

    result := captureBarrierOutput(endpoints...)
    if !strings.Contains(result, "Timeout") {
      t.Fail()
    }
    t.Log(result)
  })

}

func captureBarrierOutput(endpoints ...string) string {
  reader, writer, _ := os.Pipe()

  os.Stdout = writer

  out := make(chan string)

  go func()  {
    var buf bytes.Buffer
    io.Copy(&buf, reader)
    out <- buf.String()
  }()

  barrier(endpoints...)

  writer.Close()
  temp := <- out

  return temp
}
