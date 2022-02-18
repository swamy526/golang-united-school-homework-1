package solution

import "testing"

func TestGetMessage(t *testing.T) {
	if GetMessage() != "Hello :world_map:" {
		t.Error("Expected: Hello :world_map:")
	}
}