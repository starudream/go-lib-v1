package httpx

import (
	"net/http"
	"runtime"
	"time"

	"github.com/go-resty/resty/v2"
	"golang.org/x/net/http/httpproxy"

	"github.com/starudream/go-lib/codec/json"
	"github.com/starudream/go-lib/codec/xml"
	"github.com/starudream/go-lib/config"
	"github.com/starudream/go-lib/log"

	"github.com/starudream/go-lib/internal/ilog"
)

type (
	Request  = resty.Request
	Response = resty.Response
)

var _c *resty.Client

var (
	hdrUserAgentKey   = http.CanonicalHeaderKey("User-Agent")
	hdrUserAgentValue = runtime.Version()
)

func init() {
	_c = resty.New()
	_c.JSONMarshal = json.Marshal
	_c.JSONUnmarshal = json.Unmarshal
	_c.XMLMarshal = xml.Marshal
	_c.XMLUnmarshal = xml.Unmarshal

	_c.SetTimeout(5 * time.Minute)
	_c.SetLogger(&logger{Logger: log.With().Str("span", "http").Logger()})
	_c.SetDisableWarn(true)
	_c.SetDebug(config.GetBool("debug"))

	pc := httpproxy.FromEnvironment()
	if pc != nil && (pc.HTTPProxy != "" || pc.HTTPSProxy != "") {
		ilog.X.Debug().Msgf("proxy: %s", json.MustMarshalString(pc))
	}

	_c.SetHeader(hdrUserAgentKey, hdrUserAgentValue)
}

func Client() *resty.Client {
	return _c
}

func SetTimeout(timeout time.Duration) {
	_c.SetTimeout(timeout)
}

func SetUserAgent(ua string) {
	_c.SetHeader(hdrUserAgentKey, ua)
}

func R() *resty.Request {
	return _c.R()
}
