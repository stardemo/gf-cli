package get

import (
	"fmt"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/gcmd"
	"github.com/gogf/gf/g/os/genv"
	"github.com/gogf/gf/g/os/gproc"
	"github.com/gogf/gf/g/os/gtime"
	"math"
	"os"
	"time"
)

const (
	gPROXY_CHECK_TIMEOUT = time.Second
)

var (
	proxies = []string{
		"https://mirrors.aliyun.com/goproxy/",
		"https://goproxy.io/",
	}
	httpClient = ghttp.NewClient()
)

func init() {
	httpClient.SetTimeOut(gPROXY_CHECK_TIMEOUT)
}

func Run() {
	genv.Set("GOPROXY", getProxy())
	fmt.Fprintln(os.Stdout, "cleaning cache...")
	gproc.ShellRun("go clean -modcache")
	if value := gcmd.Value.Get(2); value != "" {
		options := gcmd.Option.Build("-")
		if options == "" {
			options = "-u"
		}
		gproc.ShellRun(fmt.Sprintf(`go get %s %s`, options, value))
	} else {
		fmt.Fprintln(os.Stdout, "downloading the latest version of GF...")
		gproc.ShellRun("go get -u github.com/gogf/gf")
	}
}

// getProxy returns the proper proxy for 'go get'.
func getProxy() string {
	if p := genv.Get("GOPROXY"); p != "" {
		return p
	}
	url := ""
	latency := math.MaxInt32
	for _, proxy := range proxies {
		if n := checkProxyLatency(proxy); n < latency {
			url = proxy
		}
	}
	return url
}

// checkProxyLatency checks the latency for specified url.
func checkProxyLatency(url string) int {
	start := gtime.Millisecond()
	r, err := httpClient.Head(url)
	if err != nil || r.StatusCode != 200 {
		return math.MaxInt32
	}
	defer r.Close()
	return int(gtime.Millisecond() - start)
}