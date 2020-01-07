package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		if err := sc.Err(); err != nil {
			fmt.Printf("read err : %v\n", err)
			os.Exit(1)
		}
		ip, err := net.ResolveIPAddr("ip", sc.Text())
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("domain: %s, ip: %s\n", sc.Text(), ip.String())
	}
}
