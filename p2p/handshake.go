package p2p

// HandshakeFunc is a function that performs a handshake with a remote peer.
type HandshakeFunc func(Peer) error

// NOPHandshakeFunc is a no-op handshake function.
func NOPHandshakeFunc(Peer) error {
	return nil
}
