package katyushalib

import (
    "fmt"
	"bytes"
	"log"
	"strings"
    //"reflect"
	"os/exec"
    "runtime"
    "encoding/json"
    "io/ioutil"

	"github.com/kr/pretty"
)

const DefaultCoroutinesCnt int = 1000

type KConfig struct {
    // main app config
    DstHost string `json:"dst_host"`
    DstPort int `json:"dst_port"`
    CoroutinesCnt int `json:"coroutines_cnt"`
    MaxProcs int `json:"max_procs"`
}

type RuntimeInfo struct {
    Uname string
    GolangVer string
    AvailableCores int
}


func (cfg *KConfig) Fulfil() {
    //xt := reflect.TypeOf(cfg.MaxProcs).Kind()
    //fmt.Printf("%T: %s\n", xt, xt)
    if (cfg.MaxProcs <= 0){
	    cfg.MaxProcs = runtime.NumCPU()
    }
    if (cfg.CoroutinesCnt <= 0){
	    cfg.CoroutinesCnt = DefaultCoroutinesCnt
    }
}

func PrettyPrint(data interface{}) {
	fmt.Printf("\n%# v\n", pretty.Formatter(data))
}

func ComposeCfg(cfg_path string) KConfig {
    content, err := ioutil.ReadFile(cfg_path)
    if err != nil {
        fmt.Print("Failed to read config file:", err)
    }

    var cfg KConfig
    err = json.Unmarshal(content, &cfg)
    if err != nil {
        fmt.Print("Failed to parse config file:", err)
    }
    cfg.Fulfil()
    return cfg
}

func get_uname() string {
	cmd := exec.Command("uname", "-sr")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
        fmt.Print(err)
	}
	return strings.Trim(out.String(), "\n")
}

func CollectRuntimeInfo() RuntimeInfo {
    info := RuntimeInfo{
        Uname: get_uname(),
        GolangVer: runtime.Version(),
        AvailableCores: runtime.NumCPU(),
    }
    return info
}

func LogRuntimeiInfo(info RuntimeInfo, usedCores int, appVersion string) {
  	log.Printf("Runtime: katyusha v.%s / golang %s / cores used %d of %d", appVersion, RuntimeInfo.GolangVer, usedCores, RuntimeInfo.AvailableCores)
}
