package _chan

import (
	"fmt"
	"time"
)

func ChanTest() {
	가 := make(chan string)
	나 := make(chan string)
	다 := make(chan string)
	라 := make(chan string)
	마 := make(chan string)

	시간다됨 := time.After(1 * time.Second)

	go func() {
		가 <- "가"
	}()
	go func() {
		나 <- "나"
	}()
	go func() {
		다 <- "다"
	}()
	go func() {
		라 <- "라"
	}()
	go func() {
		마 <- "마"
	}()

	for {
		select {
		case 가가 := <-가:
			fmt.Println(가가)
		case 나나 := <-나:
			fmt.Println(나나)
		case 다다 := <-다:
			fmt.Println(다다)
		case 라라 := <-라:
			fmt.Println(라라)
		case 마마 := <-마:
			fmt.Println(마마)
		case <-시간다됨:
			fmt.Println("시간다됨")
			return
		}
	}
}
