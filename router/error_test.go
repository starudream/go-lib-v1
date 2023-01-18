package router

import (
	"testing"
)

func TestError(t *testing.T) {
	t.Logf("%s", ErrBadRequest)
	t.Logf("%s", ErrUnauthorized)
	t.Logf("%s", ErrForbidden)
	t.Logf("%s", ErrNotFound)
	t.Logf("%s", ErrMethodNotAllowed)
	t.Logf("%s", ErrConflict)
	t.Logf("%s", ErrInternal.WithMetadata("a", 1, "b", M{"foo": "boo"}, "c", true))
}
