package katyushalib

import (
    "fmt"
//    "reflect"
    "encoding/json"
    "io/ioutil"

	"github.com/kr/pretty"
)

//type KConfig struct {
//    dst_host string `json:"dst_host"`
//    dst_port int `json:"dst_port"`
//    coroutines_cnt int `json:"coroutines_cnt"`
//    max_procs int `json:"max_procs"`
//}

type KConfig struct {
    dst_host string
    dst_port int
    coroutines_cnt int
    max_procs int
}
func (cfg *KConfig) SetDstHost(host string) {
    cfg.dst_host = host
}

func PrettyPrint(data interface{}) {
	fmt.Printf("\n%# v\n", pretty.Formatter(data))
}

func ComposeCfg(cfg_path string) KConfig {
    content, err := ioutil.ReadFile(cfg_path)
    if err != nil {
        fmt.Print("Failed to read config file:", err)
    }
    //fmt.Print(string(content[:]))

    var cfg KConfig
    //var cfg map[string]interface{}
    //content1 := []byte(`{"dst_host":"111","dst_port":"222"}`)
    //fmt.Print(content)
    err = json.Unmarshal(content, &cfg)
    if err != nil {
        fmt.Print("Failed to parse config file:", err)
    }

    //xt := reflect.TypeOf(content).Kind()

    //cfg.SetDstHost("1.2.3.4")
    //fmt.Print(cfg)
    //fmt.Printf("%T: %s\n", xt, xt)

    //cfg1 := KConfig{
    //    dst_host: "127.0.0.1",
    //    dst_port: 8080,
    //    coroutines_cnt: 1000,
    //    max_procs: 1000,
    //}
    return cfg
}
