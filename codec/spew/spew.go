package spew

import (
	"github.com/davecgh/go-spew/spew"
)

var (
	Config = &spew.Config

	Dump  = spew.Dump
	Fdump = spew.Fdump
	Sdump = spew.Sdump
)

func init() {
	Config.Indent = "  "
}
