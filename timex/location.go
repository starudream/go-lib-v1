package timex

import (
	"time"
)

var (
	UTC, _ = time.LoadLocation("UTC")
	GMT, _ = time.LoadLocation("GMT")
	CET, _ = time.LoadLocation("CET")
	PRC, _ = time.LoadLocation("PRC")
)
