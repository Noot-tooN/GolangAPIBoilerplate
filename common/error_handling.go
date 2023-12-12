package common

import (
	"fmt"
	"os"
	"sync"
)

func CheckWg(err error, wg *sync.WaitGroup) {
	if err != nil {
		fmt.Printf("Server error:\n%v\n", err)
		os.Exit(1)
	}
	// If we dont exit make sure to signal Done so we don't get locked
	if wg != nil {
		wg.Done()
	}
}
