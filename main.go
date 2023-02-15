package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "USAGE: %s host port [-revert] [-timeout milliseconds; default=500] [-help]\n", flag.CommandLine.Name())
		os.Exit(3)
	}

	revert := flag.Bool("revert", false, "revert the exit code")
	timeout := flag.Float64("timeout", 500, "timeout in milliseconds")

	flag.Parse()

	if flag.NArg() != 2 {
		flag.Usage()
		return
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", host, port), time.Duration(*timeout)*time.Millisecond)

	if err, ok := err.(*net.OpError); ok && err.Timeout() {
		if *revert {
			os.Exit(0)
			return
		}

		os.Exit(1)
		return
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(2)
		return
	}

	conn.Close()

	if *revert {
		os.Exit(1)
		return
	}

	os.Exit(0)
}
