package seq

import (
	"testing"
)

func TestUUID(t *testing.T) {
	s1 := UUID()
	t.Log(s1)

	s2 := UUIDShort()
	t.Log(s2)

	t.Log(ParseUUID(s1))
}
