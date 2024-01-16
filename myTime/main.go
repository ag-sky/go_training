package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	createDate := time.Date(2024, time.January, 22, 3, 5, 7, 7, time.UTC)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}
