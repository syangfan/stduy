package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err    error
	output []byte
}

func main() {
	//执行1个cmd,让它在一个协程里去执行，让它执行2秒；sleep 2;echo echo
	//1秒的时候，杀死cmd
	var (
		ctx        context.Context
		cancalFunc context.CancelFunc
		cmd        *exec.Cmd
		resultChan chan *result
		res *result
	)
	resultChan = make(chan *result, 1000)

	ctx, cancalFunc = context.WithCancel(context.TODO())
	go func() {
		var (
			output []byte
			err    error
		)
		cmd = exec.CommandContext(ctx, "C:\\cygwin64\\bin\\bash.exe", "-c", "sleep 2;echo hello;sleep 2")

		//执行任务,捕捉输出
		output, err = cmd.CombinedOutput()

		//把任务输出结果，传给main协程
		resultChan <- &result{
			err:    err,
			output: output,
		}
	}()

	time.Sleep(3 * time.Second)

	//取消上下文
	cancalFunc()

	res = <-resultChan

	fmt.Println(res.err,string(res.output))

}
