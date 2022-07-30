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

func main() {
	defer ants.Release()
	//runTimes := 999999999
	runTimes := 1650
	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(6, func(i interface{}) {
		Get(i)
		wg.Done()
	})
	defer p.Release()
	// Submit tasks one by one.
	for i := 1600; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int64(i))
	}
	wg.Wait()

}

func Get(i interface{}) {
	n := i.(int64)
	rand.Seed(time.Now().UnixNano())
	count := rand.Intn(30)
	util.Log().Info("%v开始,休眠%v", i, count)
	time.Sleep(time.Second * time.Duration(count))
	if service.GetUid(n) {
		service.GetData(n)
	} else {
		util.Log().Error("出错啦")
	}
}
