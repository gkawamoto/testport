package main

import (
	"fmt"
	"net"
	"time"
	"os"
)

func main() {
	if len(os.Args) < 3 || len(os.Args) > 4 {
		usage()
		return
	}
	if len(os.Args) == 4 {
		if os.Args[3] != "-revert" {
			usage()
			return
		}
	}
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", os.Args[1], os.Args[2]), 500 * time.Millisecond)

	if err, ok := err.(*net.OpError); ok && err.Timeout() {
		os.Exit(1)
		return
	}

	if err != nil {
		// Log or report the error here
		fmt.Printf("error: %s\n", err)
		os.Exit(2)
		return
	}
	conn.Close()
	os.Exit(0)
}

func usage() {
	fmt.Fprintf(os.Stderr, "USAGE: %s host port [-revert]\n", os.Args[0])
	os.Exit(1)
}
