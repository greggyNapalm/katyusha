package main 

import (
//    "encoding/json"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
//	"time"

	"github.com/kr/pretty"
	"github.com/docopt/docopt-go"

    "github.com/greggyNapalm/katyushalib"
)

const version = "0.0.1"
const workers_num = 1000
const tgt_host, tgt_port = "127.0.0.1", 80

type KConfig struct {
    dst_host []string
    dst_port int
    coroutines_cnt int
    max_procs int
}

func pp(data interface{}) {
	fmt.Printf("%# v", pretty.Formatter(data))
}

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
	// Create isingle persistent TCP connection,
	// use it in each interaction.

	//res, _ := client.Do(req)
	//io.Copy(ioutil.Discard, res.Body)
	//res.Body.Close()
	//
	//To ensure http.Client connection reuse be sure to do two things:
	//
	//Read until Response is complete (i.e. ioutil.ReadAll(resp.Body))
	//Call Body.Close()

	for {
		//http.Get("http://127.0.0.1:80")
		_, err := http.Get(dst_addr)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func get_uname() string {
	cmd := exec.Command("uname", "-sr")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Trim(out.String(), "\n")
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
    fmt.Println(arguments)
    fmt.Println(arguments["KCFG_PATH"])

    kcfg := compose_cfg(arguments["KCFG_PATH"])

	cpu_num := runtime.NumCPU()
	runtime.GOMAXPROCS(cpu_num)

	dst := compose_url(tgt_host, tgt_port)

	log.Printf("Runtime: %s / golang %s / cores count %d", get_uname(), runtime.Version(), cpu_num)
	log.Printf("Target addr: %s", dst)
}