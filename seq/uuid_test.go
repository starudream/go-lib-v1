package seq

import (
	"testing"
)

func TestUUID(t *testing.T) {
	s := UUID()
	t.Log(s)

	t.Log(ParseUUID(s))
}
