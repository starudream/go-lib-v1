package router

import (
	"github.com/starudream/go-lib/internal/gin"
)

func wrapHF(hfs1 []HandlerFunc) (hfs2 []gin.HandlerFunc) {
	hfs2 = make([]gin.HandlerFunc, len(hfs1))
	for i, hf := range hfs1 {
		hfs2[i] = gin.HandlerFunc(hf)
	}
	return
}
