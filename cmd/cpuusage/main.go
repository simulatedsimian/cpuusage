package main

import (
	"fmt"
	"os"
	"time"

	"github.com/0xcafed00d/cpuusage"
)

func main() {
	u := cpuusage.Usage{}

	for {
		err := u.Measure()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("Overall: %d Cores: %v\n", u.Overall, u.Cores)
		time.Sleep(1 * time.Second)
	}
}
