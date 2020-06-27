package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	fmt.Println("Simple port scanner")
	if len(os.Args) < 2 {
		fmt.Println("Missing target parameter")
		os.Exit(0)
	} else {
		fmt.Println("Target :", os.Args[1])
	}
	var wg sync.WaitGroup
	now := time.Now()
	for p := 1; p <= 65535; p++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			address := os.Args[1] + ":" + strconv.Itoa(p)
			conn, err := net.DialTimeout("tcp", address, time.Millisecond*40)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("Port %d openned\n", p)

		}(p)
		time.Sleep(7 * time.Millisecond)
	}
	wg.Wait()
	fmt.Println("Scan finished. Time taken :", time.Since(now))
}
