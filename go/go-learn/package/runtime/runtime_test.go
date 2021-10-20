package runtime

import (
	"fmt"
	"testing"
	"time"
)

type Tic struct {
	ticker  *time.Ticker // The channel on which the ticks are delivered.
	// contains filtered or unexported fields
}
func TestTime(t *testing.T) {

	//fmt.Println(time.Now().Minute())
	//start := time.NewTicker(5)
	//
	//tick := &Tic{ticker:start}
	//tick.ticker.Stop()
	//
	//c := time.Tick(1 * time.Minute)
	//for now := range c {
	//	fmt.Println(now)
	//}
	// 定时器
	tr := time.NewTimer(1*60)
	defer tr.Stop()
	for {
		<-tr.C
		fmt.Println("timer running...")
		// 需要重置Reset 使 t 重新开始计时
		tr.Reset(time.Second * 2)
	}

}
