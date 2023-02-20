package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now) //2023-02-20 11:12:44.6862657 +0800 CST m=+0.009519701

	fmt.Printf("%s,%T", now, now)
	Y := now.Year()
	m := now.Month()
	d := now.Day()
	H := now.Hour()
	i := now.Minute()
	s := now.Second()
	fmt.Println(Y, m, d, H, i, s)                                 //2023 February 20 11 43 59
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", Y, m, d, H, i, s) //2023-02-20 11:43:59

	fmt.Println(now.Format("2006-01-02 03:04:05")) //2023-02-20 11:43:59
	// 说明：2006表示年份
	// 01 表示月
	// 02 表示日
	// 03 表示小时    15表示24制小时
	// 04 表示 分
	// 05 表示 秒

	fmt.Println(now.Unix())      //时间搓
	fmt.Println(now.UnixMicro()) //毫秒

	// 时间戳转时间
	unix := 1676865130
	unix_time := time.Unix(int64(unix), 100000)
	fmt.Println(unix_time)

	//时间字符串转成 时间
	time_str, err := time.ParseInLocation("2006-01-02 03:04:05", "2023-02-20 11:43:59", time.Local)
	if err != nil {
		fmt.Println("错误")
	}
	fmt.Println(time_str)

	// 时间属性
	fmt.Println(time.Hour, time.Second, time.Nanosecond)
	// 时间方法
	new_time := time_str.Add(time.Hour)
	fmt.Println(new_time)
}
