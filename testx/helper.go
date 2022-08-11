package testx

import (
	"testing"

	"github.com/starudream/go-lib/codec/json"
)

func P(t *testing.T, err error, vs ...any) {
	t.Helper()

	if err != nil {
		t.Fatal(err)
		return
	}

	for i := 0; i < len(vs); i++ {
		t.Log(json.MustMarshalString(vs[i]))
	}
}
