package binding

type Binding interface {
	Name() string
	Bind([]byte, any) error
}

var (
	JSON Binding = jsonBinding{}
)
