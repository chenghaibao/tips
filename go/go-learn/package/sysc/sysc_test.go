package sysc

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

type lock struct {
	mux *sync.Mutex
}

var rw sync.RWMutex
var count int
var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

// 有很多人正在读取数据，你给我站一边去，等它们读（读解锁）完你再来写（写锁定
func TestSyscRW(t *testing.T) {

	ch := make(chan struct{}, 10)
	for i := 0; i < 5; i++ {
		go read(i, ch)
	}
	for i := 0; i < 5; i++ {
		go write(i, ch)
	}

	for i := 0; i < 10; i++ {
		<-ch
	}
}

func read(n int, ch chan struct{}) {
	rw.RLock()
	fmt.Printf("goroutine %d 进入读操作...\n", n)
	v := count
	fmt.Printf("goroutine %d 读取结束，值为：%d\n", n, v)
	rw.RUnlock()
	ch <- struct{}{}
}

func write(n int, ch chan struct{}) {
	rw.Lock()
	fmt.Printf("goroutine %d 进入写操作...\n", n)
	v := rand.Intn(1000)
	count = v
	fmt.Printf("goroutine %d 写入结束，新值为：%d\n", n, v)
	rw.Unlock()
	ch <- struct{}{}
}

// 锁
func TestSysc(t *testing.T) {
	var l sync.Mutex
	var a = 0
	// 启动 100 个协程，需要足够大
	// var lock sync.Mutex
	for i := 0; i < 100; i++ {
		go func(id int) {
			l.Lock()
			defer l.Unlock()
			a += 1
			fmt.Printf("goroutine %d, a=%d\n", id, a)
		}(i)
	}

	// 等待 1s 结束主程序
	// 确保所有协程执行完
	time.Sleep(time.Second)
}

func TestCh(t *testing.T) {
	ch := make(chan struct{}, 2)

	var l sync.Mutex
	go func() {
		l.Lock()
		defer l.Unlock()
		fmt.Println("goroutine1: 我会锁定大概 2s")
		time.Sleep(time.Second * 2)
		fmt.Println("goroutine1: 我解锁了，你们去抢吧")
		ch <- struct{}{}
	}()

	go func() {
		fmt.Println("groutine2: 等待解锁")
		l.Lock()
		defer l.Unlock()
		fmt.Println("goroutine2: 哈哈，我锁定了")
		ch <- struct{}{}
	}()

	// 等待 goroutine 执行结束
	for i := 0; i < 2; i++ {
		<-ch
	}
}

func TestOnce(t *testing.T) {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

func TestWait(t *testing.T) {
	var wg sync.WaitGroup
	seconds := []int{1, 2}
	for i, s := range seconds {
		// 计数加 1
		wg.Add(1)
		go func(i, s int) {
			// 计数减 1
			defer wg.Done()
			fmt.Printf("goroutine%d 结束\n", i)
		}(i, s)
	}

	// 等待执行结束
	wg.Wait()
	fmt.Println("所有 goroutine 执行结束")
}

// 减少重复对象的使用
func TestPool(t *testing.T) {
	p := &sync.Pool{
		New: func() interface{} {
			return 222
		},
	}

	a := p.Get().(int)
	p.Put(1)
	p.Put(2)
	b := p.Get().(int)
	c := p.Get().(int)
	fmt.Println(a, b, c)
}

// 协成等待下发通知在执行
func TestCond(t *testing.T) {
	for i := 0; i < 5; i++ {
		go func(x int){
			cond.L.Lock() //获取锁
			fmt.Println("aaa: ", x)
			cond.Wait()//等待通知  暂时阻塞
			fmt.Println("bbb: ", x)
			time.Sleep(time.Second * 2)
			cond.L.Unlock()//释放锁
		}(i)
	}
	fmt.Println("start all")
	time.Sleep(time.Second * 1)
	fmt.Println("broadcast")
	cond.Signal()   // 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 1)
	cond.Signal()// 3秒之后 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 1)
	cond.Broadcast()//3秒之后 下发广播给所有等待的goroutine
	time.Sleep(time.Second * 10)
	fmt.Println("finish all")
}
