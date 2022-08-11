package httpx

import (
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/starudream/go-lib/codec/json"
	"github.com/starudream/go-lib/codec/xml"
	"github.com/starudream/go-lib/config"
	"github.com/starudream/go-lib/internal/httpproxy"
	"github.com/starudream/go-lib/internal/ilog"
	"github.com/starudream/go-lib/log"
)

var _c *resty.Client

func init() {
	_c = resty.New()
	_c.JSONMarshal = json.Marshal
	_c.JSONUnmarshal = json.Unmarshal
	_c.XMLMarshal = xml.Marshal
	_c.XMLUnmarshal = xml.Unmarshal

	_c.SetTimeout(5 * time.Minute)
	_c.SetLogger(&logger{Logger: log.Logger()})
	_c.SetDisableWarn(true)
	_c.SetDebug(config.GetBool("debug"))

	pc := httpproxy.FromEnvironment()
	if pc != nil && (pc.HTTPProxy != "" || pc.HTTPSProxy != "") {
		ilog.X.Debug().Msgf("proxy: %s", json.MustMarshalString(pc))
	}
}

func R() *resty.Request {
	return _c.R()
}
