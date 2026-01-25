package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")
	presentTime:= time.Now();
	fmt.Println("presentTime...", presentTime)

	fmt.Println("FormatedTime...", presentTime.Format("02-01-2006 "));

	customTime:=time.Date(2016, time.April, 12, 12, 30, 30, 2000, time.UTC)

	formatCustomTime:=customTime.Format("02-01-2006 15:04:05 Monday");
	fmt.Println("custom time", customTime);
	fmt.Println("formatCustomTime time", formatCustomTime);

}