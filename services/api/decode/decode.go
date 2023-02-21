package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/flashbots/go-boost-utils/types"
)

const path = "~/bids/0x03db0c2ed0db77c483b380fe28014afb75287b369f97102e99f51462de1b2db3.json"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	trials := int(1e3)
	times := make([]int64, trials)
	for i := 0; i < trials; i++ {
		data, err := os.ReadFile(path)
		check(err)
		p := new(types.BuilderSubmitBlockRequest)
		start := time.Now()
		err = json.Unmarshal(data, &p)
		dur := time.Since(start)
		fmt.Printf("trial: %d: timing = %v microseconds\n", i, dur.Microseconds())
		check(err)
		times[i] = dur.Microseconds()
	}
	fmt.Printf("times=%v\n", times)
}
