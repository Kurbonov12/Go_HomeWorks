package main

import (
	"fmt"
	"mygo/internal/task"
	"time"
)

func main() {
	t := task.New("", "", "", time.Now(), false)
	fmt.Println(t.Id)
	t1 := task.New("", "", "", time.Now(), false)
	fmt.Print(t1.Id)
}
