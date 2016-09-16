package main

import (
	"fmt"
	"strconv"
	"time"
)

type Err struct {
	Code int
	Msg  string
}

func (e Err) String() string {
	return time.Now().String() + ". Code: " + strconv.Itoa(e.Code) + ". Msg: " + e.Msg
}

func main() {

	e := Err{
		Code: 101,
		Msg:  "Ugly error msg body.",
	}

	fmt.Println(e)
}
