package main

import (
	"fmt"
	"time"
)

var ch = make(chan int, 1)

func main() {

	go Work("goroutine1")
	<-ch
	go Work("goroutine2")
	<-ch
	go Work("goroutine3")
	<-ch
	fmt.Println("successful")
}

func Work(workName string) {

	time.Sleep(time.Second)
	ch <- 1
	fmt.Println(workName)

}
