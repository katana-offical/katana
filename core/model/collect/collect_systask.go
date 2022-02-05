package collect

import (
	"context"
	"fmt"
	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pkg/pool/goroutine"
	"log"
	"time"
)

type SysTask struct {
	*Task
	*gnet.EventServer
	Addr       string
	Multicore  bool
	Async      bool
	Codec      gnet.ICodec
	WorkerPool *goroutine.Pool
}

func (st *SysTask) Start() (err error) {
	go gnet.Serve(st, st.Addr, gnet.WithMulticore(st.Multicore), gnet.WithCodec(st.Codec), gnet.WithTCPKeepAlive(time.Minute*5))
	return nil
}

func (st *SysTask) Stop(ctx context.Context) (err error) {
	return gnet.Stop(ctx, st.Addr)
}

func (st *SysTask) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Printf("Test config server is listening on %s (multi-cores: %t, loops: %d)\n",
		srv.Addr.String(), srv.Multicore, srv.NumEventLoop)
	return
}

func (st *SysTask) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	fmt.Println("frame:", string(frame))
	if st.Async {
		data := append([]byte{}, frame...)
		_ = st.WorkerPool.Submit(func() {
			c.AsyncWrite(data)
			fmt.Println(string(data))
		})
		return
	}
	out = frame
	return
}
