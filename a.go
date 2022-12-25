package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type user struct {
	Username string    `json:"username"`
	Nickname string    `json:"nickname"`
	Sex      uint8     `json:"sex"`
	Birthday time.Time `json:"birthday"`
}

// 结构体变量成员首字母未大写
func main() {
	u := user{
		Username: "坤坤",
		Nickname: "阿坤",
		Sex:      20,
		Birthday: time.Now(),
	}
	bs, err := json.Marshal(&u)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(bs))
}
