package p2p

// HandshakeFunc is a function that performs a handshake with a remote peer.
type HandshakeFunc func(any) error

// NOPHandshakeFunc is a no-op handshake function.
func NOPHandshakeFunc(any) error {
	return nil
}
