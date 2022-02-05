package main

import (
	"fmt"
	"katana/core/model/collect"
	"time"
)

func main() {
	sysTask := &collect.SysTask{Addr: "tcp://:514", Multicore: true, Async: false, Codec: nil}
	c, err := collect.Server(sysTask)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(10000000000)
	c.StopServer()
}
