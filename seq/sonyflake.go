package seq

import (
	"strconv"
	"time"

	"github.com/starudream/go-lib/config"
	"github.com/starudream/go-lib/internal/ilog"
	"github.com/starudream/go-lib/internal/sonyflake"
)

var (
	_sf *sonyflake.Sonyflake

	startTime = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	machineId = uint16(0)
)

func init() {
	_startTime, _machineId := config.GetTime("sonyflake.start_time"), config.GetUint("sonyflake.machine_id")
	if _startTime.IsZero() || _startTime.After(time.Now()) {
		_startTime = startTime
	}
	_startTime = _startTime.Truncate(time.Second)

	setting := sonyflake.Settings{StartTime: _startTime}
	if _machineId > 0 {
		setting.MachineID = func() (uint16, error) {
			return uint16(_machineId), nil
		}
	}

	_sf = sonyflake.NewSonyflake(setting)
	if _sf == nil {
		ilog.X.Fatal().Msgf("sonyflake setting error")
	}

	id, _ := _sf.NextID()

	startTime, machineId = _startTime, uint16(sonyflake.MachineID(id))

	ilog.X.Debug().Msgf("sonyflake settings: {\"startTime\":\"%s\",\"machineId\":%d}", startTime.Format(time.RFC3339Nano), machineId)
}

func MachineId() uint16 {
	return machineId
}

func NextId() string {
	id, _ := _sf.NextID()
	return strconv.FormatUint(id, 10)
}

func FillId(id *string, force ...bool) {
	if id == nil {
		return
	}
	if *id == "" || (len(force) > 0 && force[0]) {
		*id = NextId()
	}
}

func IdTime(id string) time.Time {
	ui, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return time.Time{}
	}
	return startTime.Add(sonyflake.ElapsedTime(ui))
}
