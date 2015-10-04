/*
* @Author: Narayanaperumal Gurusamy
* @Date:   2015-10-04 22:51:19
* @Last Modified by:   Narayanaperumal Gurusamy
* @Last Modified time: 2015-10-04 23:01:18
 */
package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	port = flag.String("p", "12124", "Port number to listen on")
)

func e(err error) {
	if err != nil {
		fmt.Printf("Erroring out: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	// Bind to the port
	bind, err := net.ResolveUDPAddr("0.0.0.0:" + *port)
	e(err)

	// Create listener
	listener, err := net.ListenUDP("udp", bind)
	e(err)

	fmt.Printf("Listening to port %s...\n", *port)
	for {
		// read into a new buffer
		buffer := make([]byte, 1024)
		_, _, err := listener.ReadFrom(buffer)
		e(err)

		// log to standard output
		fmt.Println(string(buffer))
	}
}
