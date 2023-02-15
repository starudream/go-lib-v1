package spew

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	data := map[string]any{
		"a": 1,
		"b": "b",
		"c": true,
		"d": 1.23,
		"e": []string{"1", "2", "3"},
		"f": map[string]any{
			"f1": []int{1, 2, 3},
		},
	}
	fmt.Println(Sdump(data))
}
