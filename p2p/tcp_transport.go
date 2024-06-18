package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node over an established TCP connection.
type TCPPeer struct {
	// conn is the underlying TCP connection.
	conn net.Conn
	// outbound is true if the connection was initiated by the local node.
	outbound bool
}

func newTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

// TCPTransport is a transport implementation using TCP.
type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	shakeHands    HandshakeFunc
	decoder       Decoder

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

// NewTCPTransport creates a new TCP transport.
func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		shakeHands:    NOPHandshakeFunc,
		listenAddress: listenAddress,
	}
}

// ListenAndAccept listens for incoming connections and accepts them.
func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

// startAcceptLoop accepts incoming connections in a loop.
func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}

		go t.handleConn(conn)
	}
}

type Temp struct{}

// HandleConn handles an incoming connection.
func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := newTCPPeer(conn, true)

	if err := t.shakeHands(peer); err != nil {
		panic("hello")
	}

	//read loop
	msg := &Temp{}
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			fmt.Printf("tcp error: %s\n", err)
			continue
		}
	}
}
