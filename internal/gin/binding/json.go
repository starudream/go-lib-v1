package binding

import (
	"github.com/starudream/go-lib/internal/gin/internal/json"
)

type jsonBinding struct{}

func (b jsonBinding) Name() string {
	return "json"
}

func (b jsonBinding) Bind(bs []byte, obj any) error {
	err := json.Unmarshal(bs, obj)
	if err != nil {
		return err
	}
	return validate(obj)
}
