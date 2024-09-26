package main

import (
	"testing"
)

func TestPeerDiscovery(t *testing.T) {
	peers := PeerDiscovery()
	_, selectedPeer := AskForPeerSelection(peers) // This will not work in test environment
	t.Logf("Selected peer: %s\n", selectedPeer)
}
