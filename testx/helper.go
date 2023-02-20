package testx

import (
	"testing"

	"github.com/starudream/go-lib/codec/json"
	"github.com/starudream/go-lib/codec/spew"
)

func P(t *testing.T, err error, vs ...any) {
	t.Helper()

	if err != nil {
		t.Fatal(err)
		return
	}

	for i := 0; i < len(vs); i++ {
		switch x := vs[i].(type) {
		case string:
			t.Log(x)
		case []byte:
			t.Log(string(x))
		default:
			s, e := json.Marshal(x)
			if e == nil {
				t.Log(string(s))
			} else {
				t.Log(spew.Sdump(x))
			}
		}
	}
}
