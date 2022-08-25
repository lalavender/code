package lib_time

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	fmt.Println(now)
	month := now.Month()
	fmt.Println(month)
}
func TestTimeUnix(t *testing.T) {
	now := time.Now()
	unix := now.Unix()
	fmt.Println("unix time:", unix)

	milli := now.UnixMilli()
	fmt.Println(milli)
	nowTime := time.UnixMilli(milli)
	fmt.Println(nowTime)
}

func TestParseTime(t *testing.T) {
	t1, err := time.Parse("2006-01-02 15:04:05", "2022-08-25 22:13:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(t1)

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)

	}
	fmt.Println(location)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", "2022-08-25 22:13:00", location)

	if err != nil {
		panic(err)
	}
	fmt.Println(t2)

}

func TestFormat(t *testing.T) {
	now := time.Now()
	format := now.Format("2006/01/02 15:04:05")
	fmt.Println(format)

}

func TestDuration(t *testing.T) {
	fiveMinute := 5 * time.Minute
	now := time.Now()
	add := now.Add(fiveMinute)
	fmt.Println(add)
	sub := add.Sub(now)
	fmt.Println(sub)

}

func TestTimeComputer(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Shanghai")
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-08-25 22:13:00", location)
	t2, _ := time.Parse("2006-01-02 15:04:05", "2022-08-25 22:13:00")
	fmt.Println(t1.Equal(t2))
	fmt.Println(t1 == t2)
	t3, _ := time.Parse("2006-01-02 15:04:05", "2022-08-25 21:13:00")
	fmt.Println(t2.After(t3))
	fmt.Println(t2.Before(t3))
}

func TestTimer(t *testing.T) {
	//tick := time.Tick(5 * time.Second)
	//for t1 := range tick {
	//	fmt.Println(t1)
	//}

	//time.AfterFunc(5*time.Second, func() {
	//	fmt.Println("5...")
	//})
	time.Sleep(5 * time.Second)
	fmt.Println(time.Now())
	for {

	}
	//time.NewTicker()
	//time.NewTimer()
}
