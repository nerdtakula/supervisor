package supervisor

import (
  "fmt"
  "testing"
)

const (
  TEST_HOST    = "localhost"
  TEST_PORT    = 9001
  TEST_USER    = "admin"
  TEST_PASS    = "admin"
  TEsT_PROCESS = "myService"
)

func assertEqual(t *testing.T, expected, result interface{}, message string) {
  if expected == result {
    return
  }
  
  if message == "" {
    message = fmt.Sprintf("%v != %v", expected, result)
  }
  t.Fatal(message)
}
