package another

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	//线程安全：并发时只读
	M *sync.Map
}

func (c *Counter) Incr(s interface{}, n int) {
	//成功读取则累加
	if addend, loadSuccess := c.M.LoadOrStore(s, n); loadSuccess {
		c.M.Store(s, addend.(int)+n)
	}
}

func (c *Counter) Init() {
	*c = Counter{new(sync.Map)}
}

//获取数据，如果成功则返回累加结果
func (c *Counter) Get(s interface{}) int {
	addend, loadSuccess := c.M.Load(s)
	if loadSuccess {
		return addend.(int)
	} else {
		return 0
	}
}

func (c *Counter) Flush2Broker(ms int, FuncCbFlush func()) {
	go func() {
		//类型转换：int->时间类型（毫秒）
		ticker := time.NewTicker(time.Millisecond * time.Duration(ms))
		//异步
		go func() {
			for range ticker.C {
				//每5秒重置一次计数器，一秒一次
				// fmt.Printf("每%.2fs重置计数器\n", float64(ms/1000))
				FuncCbFlush()
			}
		}()
	}()
}

// 重置计数器
func (c *Counter) FuncCbFlush() {
	c.Init()
}

func Test() {
	counter := &Counter{}
	counter.Init()
	counter.Flush2Broker(5000, counter.FuncCbFlush)
	//counter.Incr("get.called", 123)
	//counter.Incr("get.called", 456)

	//for i := 0; i < 12; i++ {
	//	counter.Incr("get.called", 1)
	//	fmt.Println(counter.Get("get.called"))
	//	time.Sleep(1 * time.Second)
	//}

	//var a = time.Millisecond
	//fmt.Println(a)
	var timeSum = time.Millisecond
	for i := 0; i < 10000; i++ {
		start := time.Now()
		//fmt.Println("start:", start)

		counter.Incr("get.called", 123)
		counter.Incr("get.called", 456)
		for i := 0; i < 10000; i++ {
			counter.Incr("get.called", 1)
		}

		//println(counter.Get("get.called"))
		cost := time.Since(start)
		timeSum += cost
		//fmt.Println("cost:", cost)
		//timeSum += float64(cost)
	}
	fmt.Println("平均运行时间：",(timeSum - time.Millisecond) / 10000)

	//time.Sleep(6 * time.Second)
	//println(counter.Get("get.called"))
	//
	//time.Sleep(6 * time.Second)

}
