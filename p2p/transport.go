// Package p2p defines the interfaces for the transport layer of the P2P network.
package p2p

// Peer is a remote node in the network.
type Peer interface {
}

// Transport handles communication between peers.
// It can be of different types, e.g. TCP, UDP, etc.
type Transport interface {
	ListenAndAccept() error
}
