package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeDemo(t *testing.T) {
	now := time.Now() //获取当前时间
	// current time:2020-12-01 22:24:30.85736 +0800 CST m=+0.000096031
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	fmt.Printf("%T,%T,%T,%T,%T,%T,%T\n", now, year, month, day, hour, minute, second)
	// time.Time,int,time.Month,int,int,int,int

	now1 := time.Now()
	// 当前时间戳 TimeStamp type:int64, TimeStamp:1606832965
	fmt.Printf("TimeStamp type:%T, TimeStamp:%v", now1.Unix(), now1.Unix())

	now2 := time.Now()
	// 纳秒级时间戳TimeStamp type:int64, TimeStamp:1606833059999670000
	fmt.Printf("TimeStamp type:%T, TimeStamp:%v\n", now2.UnixNano(), now2.UnixNano())

	timestamp := time.Now().Unix()
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	tm2, _ := time.Parse("01/02/2006", "02/08/2015")
	fmt.Println(tm2.Unix())
	fmt.Println(timeObj.Format("2006-01-02 03:04:05"))
	fmt.Println(timeObj.Format("2006-01-02 03:04:05 PM"))
	fmt.Println(timeObj.Format("02/01/2006 15:04:05 PM"))
	fmt.Println(timeObj)
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
