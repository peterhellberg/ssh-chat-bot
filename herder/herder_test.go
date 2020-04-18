package herder

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	h := New(
		User("u"),
		Owner("o"),
		Addr("a"),
		Delay(time.Second),
		Check(time.Minute),
		Verbose(true),
	)

	if err := h.validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
