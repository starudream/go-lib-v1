package cronx

import (
	"time"

	"github.com/robfig/cron/v3"

	"github.com/starudream/go-lib/log"
)

var _c *cron.Cron

func init() {
	l := &logger{Logger: log.With().Str("span", "cron").Logger()}

	_c = cron.New(
		cron.WithLocation(time.Local),
		cron.WithSeconds(),
		cron.WithLogger(l),
		cron.WithChain(cron.Recover(l)),
	)
}

var (
	Start = _c.Start
	Stop  = _c.Stop
	Run   = _c.Run

	Entries = _c.Entries
	Entry   = _c.Entry
	Remove  = _c.Remove
)

type Job interface {
	Name() string
	Run()
}

func AddJob(spec string, job Job) (cron.EntryID, error) {
	return _c.AddJob(spec, job)
}
