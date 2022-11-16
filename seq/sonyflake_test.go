package seq

import (
	"testing"
	"time"
)

func TestSonyflake(t *testing.T) {
	t.Log(MachineId())

	id1, id2 := NextId(), NextId()
	t.Log(id1, id2)

	t.Log(IdTime(id1).Local().Format(time.RFC3339Nano))

	id3 := ""
	FillId(&id3)
	t.Log(id3)
	FillId(&id3)
	t.Log(id3)
	FillId(&id3, true)
	t.Log(id3)
}
