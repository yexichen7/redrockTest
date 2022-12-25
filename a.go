package main

import "fmt"

//延迟处理的语句按 defer 的逆序进行执行,"3"语句在return的后面，未执行。若要输出可如下执行
func main() {
	var a = true
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("3")
	}()
	if a {
		fmt.Println("2")
		return
	}

}
