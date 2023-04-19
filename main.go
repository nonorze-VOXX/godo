package main

import (
	"fmt"
	"time"
)

type UnitTodo struct {
	title string
	sTime time.Time
	eTime time.Time
}

func main() {

}
func checkErr(err error) {
	if err != nil {
		fmt.Print("1")
		panic(err)
	}
}
