package supervisor

import (
	"fmt"
	"testing"
)

const (
	TEST_HOST    = "localhost"
	TEST_PORT    = 9001
	TEST_SOCKET  = "/var/run/supervisor.sock"
	TEST_USER    = "admin"
	TEST_PASS    = "admin"
	TEST_PROCESS = "myService"
)

func assertEqual(t *testing.T, expected, result interface{}, message string) {
	if expected == result {
		return
	}

	if message == "" {
		message = fmt.Sprintf("Expected: '%+v', Got: '%+v'", expected, result)
	}
	t.Fatal(message)
}

func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	return func(t *testing.T) {
		t.Log("teardown test case")
	}
}
