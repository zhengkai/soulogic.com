package main

import (
	"fmt"

	"github.com/zhengkai/zj"
)

var (
	logTunnel = 1000

	jInfo = zj.NewPure(&zj.Config{
		Caller: zj.CallerShorter,
		Tunnel: logTunnel,
		Echo:   true,
		ErrorFn: func(o *zj.Logger) {
			fmt.Println(`info log fail:`, o.Error)
		},
		Filename:   `log/log.txt`,
		TimeFormat: zj.TimeMS,
	})

	jWarn = zj.NewPure(&zj.Config{
		Caller: zj.CallerShorter,
		Tunnel: logTunnel,
		Echo:   true,
		ErrorFn: func(o *zj.Logger) {
			fmt.Println(`warn log fail:`, o.Error)
		},
		Filename:   `log/warn.txt`,
		TimeFormat: zj.TimeMS,
	})

	jDump = zj.NewPure(&zj.Config{
		// Echo:   true,
		Tunnel: logTunnel,
		ErrorFn: func(o *zj.Logger) {
			fmt.Println(`dump log fail:`, o.Error)
		},
		Filename: `log/dump.txt`,
	})

	jGateway = zj.NewPure(&zj.Config{
		// Echo:   true,
		Tunnel: logTunnel,
		ErrorFn: func(o *zj.Logger) {
			fmt.Println(`gateway log fail:`, o.Error)
		},
		Filename:   `log/gateway.txt`,
		TimeFormat: zj.TimeMS,
	})

	j   = jInfo.Log
	jf  = jInfo.Logf
	jw  = jWarn.Log
	jwf = jWarn.Logf
)

func initLog() {
	jDump.Color(`38;2;200;255;100`)
	jWarn.Color(`38;2;255;130;130`)
}

func closeLog() {
	jInfo.Close()
	jWarn.Close()
	jDump.Close()
	jGateway.Close()
}
