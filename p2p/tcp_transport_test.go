package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddress := ":4080"
	tr := NewTCPTransport(listenAddress)

	assert.Equal(t, tr.listenAddress, listenAddress)

	assert.Nil(t, tr.ListenAndAccept())

	select {}
}
