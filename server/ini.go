package main

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	httpListenAddr = ``
)

func initConfig() (err error) {

	defaultPort := `21001`

	cfg, err := ini.Load(pwd + `/config.ini`)
	if err != nil {
		jw(`config error:`, err)
		httpListenAddr = `:` + defaultPort
		return
	}

	cfgHTTP := cfg.Section(`http`)

	host := cfgHTTP.Key(`host`).MustString(``)
	port := cfgHTTP.Key(`port`).MustString(defaultPort)

	httpListenAddr = fmt.Sprintf(`%s:%s`, host, port)

	return
}
