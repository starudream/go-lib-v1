package errx

import (
	"fmt"
	"testing"

	"github.com/starudream/go-lib/codec/json"
)

func Test(t *testing.T) {
	t.Log(ErrParam)
	t.Log(ErrUnAuth)
	t.Log(ErrForbidden)
	t.Log(ErrNotFound)
	t.Log(ErrNoMethod)
	t.Log(ErrConflict)

	e1 := OK.WithMetadata("foo", "e1")
	t.Log(e1)
	e2 := e1.WithMetadata("bar", "e2")
	t.Log(e2)
	t.Log(e1)

	e9 := fmt.Errorf("new error: %w", e1)
	t.Log(e9)
	t.Log(From(e9))

	t.Log(json.MustMarshalString(e2))
}
