package solution

import "testing"

// TestGetMessage calls solution.GetMessage, checking
// for a valid return value.
func TestGetMessage(t *testing.T) {
	actual_msg := "Hello :world_map:"
	msg :=  GetMessage()
	if msg != actual_msg {
		t.Error("Expected: Hello :world_map:")
	}
}