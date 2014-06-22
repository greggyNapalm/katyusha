package katyushalib

type KConfig struct {
    dst_host string
    dst_port int
    coroutines_cnt int
    max_procs int
}


func compose_cfg(cfg_path string) KConfig {
    cfg := KConfig{
        dst_host: "127.0.0.1",
        dst_port: 8080,
        coroutines_cnt: 1000,
        max_procs: 1000,
    }
    return cfg
}
