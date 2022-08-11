package httpproxy

import (
	"os"
)

// Config holds configuration for HTTP proxy settings. See
// FromEnvironment for details.
type Config struct {
	// HTTPProxy represents the value of the HTTP_PROXY or
	// http_proxy environment variable. It will be used as the proxy
	// URL for HTTP requests unless overridden by NoProxy.
	HTTPProxy string `json:"http_proxy"`

	// HTTPSProxy represents the HTTPS_PROXY or https_proxy
	// environment variable. It will be used as the proxy URL for
	// HTTPS requests unless overridden by NoProxy.
	HTTPSProxy string `json:"https_proxy"`

	// NoProxy represents the NO_PROXY or no_proxy environment
	// variable. It specifies a string that contains comma-separated values
	// specifying hosts that should be excluded from proxying. Each value is
	// represented by an IP address prefix (1.2.3.4), an IP address prefix in
	// CIDR notation (1.2.3.4/8), a domain name, or a special DNS label (*).
	// An IP address prefix and domain name can also include a literal port
	// number (1.2.3.4:80).
	// A domain name matches that name and all subdomains. A domain name with
	// a leading "." matches subdomains only. For example "foo.com" matches
	// "foo.com" and "bar.foo.com"; ".y.com" matches "x.y.com" but not "y.com".
	// A single asterisk (*) indicates that no proxying should be done.
	// A best effort is made to parse the string and errors are
	// ignored.
	NoProxy string `json:"no_proxy"`

	// CGI holds whether the current process is running
	// as a CGI handler (FromEnvironment infers this from the
	// presence of a REQUEST_METHOD environment variable).
	// When this is set, ProxyForURL will return an error
	// when HTTPProxy applies, because a client could be
	// setting HTTP_PROXY maliciously. See https://golang.org/s/cgihttpproxy.
	CGI bool `json:"cgi"`
}

// FromEnvironment returns a Config instance populated from the
// environment variables HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or the
// lowercase versions thereof). HTTPS_PROXY takes precedence over
// HTTP_PROXY for https requests.
//
// The environment values may be either a complete URL or a
// "host[:port]", in which case the "http" scheme is assumed. An error
// is returned if the value is a different form.
func FromEnvironment() *Config {
	return &Config{
		HTTPProxy:  getEnvAny("HTTP_PROXY", "http_proxy"),
		HTTPSProxy: getEnvAny("HTTPS_PROXY", "https_proxy"),
		NoProxy:    getEnvAny("NO_PROXY", "no_proxy"),
		CGI:        os.Getenv("REQUEST_METHOD") != "",
	}
}

func getEnvAny(names ...string) string {
	for _, n := range names {
		if val := os.Getenv(n); val != "" {
			return val
		}
	}
	return ""
}
