package main

import (
	"fmt"
	"github.com/kroikie/durationslice"
	"time"
)

func main() {
	ds, err := durationslice.Process(time.Duration((48*time.Hour)+(10*time.Second)), "d min s")
	if err != nil {
		print(err.Error())
		return
	}
	fmt.Printf("%@", ds)
}
