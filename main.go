package main

import (
	"github.com/panjf2000/ants/v2"
	"github.com/techoc/bili/model"
	"github.com/techoc/bili/service"
	"github.com/techoc/bili/util"
	"math/rand"
	"sync"
	"time"
)

func init() {
	model.Init()
}

var wg sync.WaitGroup

var isContinue bool

func main() {
	isContinue = true
	defer ants.Release()
	//runTimes := 999999999
	//k := 1000
	w := 10000
	runTimes := 9 * w
	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(6, func(i interface{}) {
		Get(i)
		wg.Done()
	})
	defer p.Release()
	// Submit tasks one by one.
	for i := 8 * w; i < runTimes; i++ {
		if !isContinue {
			util.Log().Error("正在结束任务")
			break
		}
		wg.Add(1)
		_ = p.Invoke(int64(i))
	}
	wg.Wait()
}

func Get(i interface{}) {
	n := i.(int64)
	status := service.GetUid(n)
	if status == 1 {
		//id 存在
		service.GetData(n)
	} else if status == 0 {
		util.Log().Error("出错啦,ID不存在")
	} else if status == -1 {
		util.Log().Error("请求被拦截，结束任务")
		isContinue = false
		return
	}

	{
		rand.Seed(time.Now().UnixNano())
		count := 10
		time.Sleep(time.Second * time.Duration(count))
		util.Log().Info("休眠10s")
	}
}
