package supervisor

import (
	"testing"
)

func TestNewHeaderToken(t *testing.T) {
	// Arrange
	headerStr := "ver:3.0 server:supervisor serial:21 pool:listener poolserial:10 eventname:PROCESS_COMMUNICATION_STDOUT len:54"

	// Act
	token, err := NewHeaderToken(headerStr)

	// Assert
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, "3.0", token.Version, "")
	assertEqual(t, "supervisor", token.Server, "")
	assertEqual(t, 21, token.Serial, "")
	assertEqual(t, "listener", token.Pool, "")
	assertEqual(t, 10, token.PoolSerial, "")
	assertEqual(t, Event(EVENT_PROCESS_COMMUNICATION_STDOUT).String(), token.EventName.String(), "")
	assertEqual(t, 54, token.Length, "")
}
