package timex

import (
	"testing"
)

func TestLocation(t *testing.T) {
	t.Logf("%#v", UTC)
	t.Logf("%#v", GMT)
	t.Logf("%#v", CET)
	t.Logf("%#v", PRC)
}
