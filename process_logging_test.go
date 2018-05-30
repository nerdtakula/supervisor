package supervisor

import (
	"testing"
)

func TestReadProcessStdout(t *testing.T) {
	// Setup requirements
	// setupTeardown := setupTestCase(t)
	// defer setupTeardown(t)

	// Create client, Read STDOUT of process, compare output
	client := New(TEST_HOST, TEST_PORT, TEST_USER, TEST_PASS)
	stdout, err := client.ReadProcessStdout("foobar", 0, 11)
	if err != nil {
		t.Fatal(err)
	}
	if expected := "stdout text"; stdout != expected {
		t.Fatalf("expected '%s', got: '%s'", expected, stdout)
	}
}

func TestReadProcessStderr(t *testing.T)   { t.Skip("Not Implemented") }
func TestTailProcessStdout(t *testing.T)   { t.Skip("Not Implemented") }
func TestTailProcessStderr(t *testing.T)   { t.Skip("Not Implemented") }
func TestClearProcessLogs(t *testing.T)    { t.Skip("Not Implemented") }
func TestClearAllProcessLogs(t *testing.T) { t.Skip("Not Implemented") }
