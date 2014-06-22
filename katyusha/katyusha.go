package main 

import (
//    "encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/docopt/docopt-go"
    "github.com/greggyNapalm/katyusha/katyushalib"
)

const version = "0.0.1"
var pp = katyushalib.PrettyPrint 

func compose_url(tgt_addr string, tgt_port int) string {
	return fmt.Sprintf("http://%s:%d", tgt_addr, tgt_port)
}

func remote_deal(dst_addr string) {
	// Close TCP connection on each request
	for {
		//http.Get("http://127.0.0.1:80")
		_, err := http.Get(dst_addr)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func remote_deal_reuse(dst_addr string) {
	for {
		//http.Get("http://127.0.0.1:80")
		_, err := http.Get(dst_addr)
		if err != nil {
			log.Fatal(err)
		}
	}
}



func main() {
      usage := `katyusha.

Usage:
  katyusha KCFG_PATH
  katyusha -h | --help
  katyusha --version

Arguments:
  KCFG_PATH     Katyushas config file path.


Options:
  -h --help     Show this screen.
  -v --verbose  Give more verbose output.
  --version     Show version.`

    arguments, _ := docopt.Parse(usage, nil, true, "Katyusha load tool v." + version, false)
    kcfg := katyushalib.ComposeCfg(arguments["KCFG_PATH"].(string))
    runtimeInfo := katyushalib.CollectRuntimeInfo()

	runtime.GOMAXPROCS(kcfg.MaxProcs)
    katyushalib.LogRuntimeiInfo(runtimeInfo, kcfg.MaxProcs, version)
}
