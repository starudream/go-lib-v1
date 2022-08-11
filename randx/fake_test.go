package randx

import (
	"testing"
)

func Test(t *testing.T) {
	t.Log(F().LetterN(16))
	t.Log(F().DigitN(16))

	t.Log(F().Name())
	t.Log(F().FirstName())
	t.Log(F().LastName())
}
