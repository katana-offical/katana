package collect

import (
	"context"
	"crypto/rand"
	"fmt"
)

type Collect struct {
	CollectId  string
	ctx        context.Context
	taskHandle TaskHandle
}

type Action int

const (
	None Action = iota
	NewTask
	Start
	Stop
)

func Server(handle TaskHandle) (*Collect, error) {
	id, err := rand.Read(make([]byte, 16))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithCancel(context.Background())
	c := &Collect{
		CollectId:  string(id),
		ctx:        ctx,
		taskHandle: handle,
	}
	fmt.Sprintf("creat new task,taskid:%q", c.CollectId)
	err = c.taskHandle.Start()
	if err != nil {
		return nil, err
	}
	fmt.Sprintf("start task,taskid:%q|action:%q", c.CollectId, Start)
	return c, nil
}

func (c *Collect) StopServer() error {
	err := c.taskHandle.Stop(c.ctx)
	if err != nil {
		return err
	}
	fmt.Sprintf("stop task,taskid:%q|action:%q", c.CollectId, Stop)
	return nil
}

type (
	TaskHandle interface {
		Start() (err error)
		Stop(ctx context.Context) (err error)
	}
	Task struct{}
)

func (t *Task) Start() (err error) {
	return
}

func (t *Task) Stop(ctx context.Context) (err error) {
	return
}
