//写一个异步加法：+123，+456
//安全：counter struct
//每五秒重置一次
package Counter_Golang

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu *sync.RWMutex
	m  map[string]int
} //{m: make(map[string]int)}

func (c *Counter) Incr(s string, a int) {
	c.mu.Lock()
	c.m[s] += a
	c.mu.Unlock()
}

func (c *Counter) Get(s string) int {
	var res int
	c.mu.RLock()
	res = c.m[s]
	c.mu.RUnlock()
	return res
}

func (c *Counter) Init() {
	c.mu = new(sync.RWMutex)
	c.m = make(map[string]int)
}

func (c *Counter) Flush2Broker(ms int, FuncCbFlush func()) {
	go func() {
		//类型转换：int->时间类型（毫秒）
		ticker := time.NewTicker(time.Millisecond * time.Duration(ms))
		//异步

		for range ticker.C {
			//每5秒重置一次计数器，一秒一次
			// fmt.Printf("每%.2fs重置计数器\n", float64(ms/1000))
			// c.Init()
			FuncCbFlush()
		}
	}()
}

func Test() {
	c := &Counter{}
	c.Init()

	FuncCbFlush := func() {
		c.Init()
		// println("Flushing... counter")
	}
	c.Flush2Broker(5000, FuncCbFlush)

	var timeSum = time.Millisecond
	for i := 0; i < 10000; i++ {
		start := time.Now()
		//fmt.Println("start:", start)

		c.Incr("get.called", 123)
		c.Incr("get.called", 456)

		for i := 0; i < 10000; i++ {
			c.Incr("get.called", 1)
		}

		//println(counter.Get("get.called"))
		cost := time.Since(start)
		timeSum += cost
		//fmt.Println("cost:", cost)
		//timeSum += float64(cost)
	}
	fmt.Println("平均运行时间：", (timeSum-time.Millisecond)/10000)

	//c.Incr("get.called", 123)
	//c.Incr("get.called", 456)
	//for i := 0; i < 10000; i++ {
	//	c.Incr("get.called", 1)
	//}
	//
	//println(c.Get("get.called"))
	//
	//time.Sleep(6 * time.Second)
	//println(c.Get("get.called"))
	//
	//time.Sleep(6 * time.Second)

	//counter.m["res"] = 0
	//
	//counter.Lock()
	//counter.m["res"] += 456
	//counter.Unlock()
	//
	//for i := 0; i < 12; i++ {
	//	counter.Lock()
	//	counter.m["res"] += 123
	//	fmt.Println(counter.m["res"])
	//	time.Sleep(1 * time.Second)
	//	counter.Unlock()
	//
	//	counter.Lock()
	//	counter.m["res"] += 456
	//	fmt.Println(counter.m["res"])
	//	time.Sleep(1 * time.Second)
	//	counter.Unlock()

}
