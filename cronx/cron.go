package cronx

import (
	"time"

	"github.com/robfig/cron/v3"

	"github.com/starudream/go-lib/log"
)

type (
	Entry   = cron.Entry
	EntryID = cron.EntryID
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

type Job interface {
	Name() string
	Run()
}

func AddJob(spec string, job Job) (EntryID, error) {
	return _c.AddJob(spec, job)
}

func Run() {
	_c.Run()
}

func Entries() []Entry {
	return _c.Entries()
}

func Remove(id EntryID) {
	_c.Remove(id)
}
