package ex

import (
	"testing"

	"github.com/starudream/go-lib/router/dx"
)

func TestError(t *testing.T) {
	t.Logf("%s", BadRequest)
	t.Logf("%s", Unauthorized)
	t.Logf("%s", Forbidden)
	t.Logf("%s", NotFound)
	t.Logf("%s", MethodNotAllowed)
	t.Logf("%s", Conflict)
	t.Logf("%s", Internal.WithMetadata("a", 1, "b", dx.M{"foo": "boo"}, "c", true))
}
