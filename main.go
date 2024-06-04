// Package main: This is the main file for the p2p package. It creates a new TCP transport and listens on the specified address.
package main

import (
	"log"

	"whatever.com/fs/p2p"
)

func main() {
	listenAddress := ":4080"
	tr := p2p.NewTCPTransport(listenAddress)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal("ListenAndAccept failed")
	}

	select {}
}
