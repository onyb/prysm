package p2p

import (
	"sync"

	"github.com/ethereum/go-ethereum/p2p/enr"
	"github.com/libp2p/go-libp2p-core/network"
	peer "github.com/libp2p/go-libp2p-peer"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/prysmaticlabs/prysm/beacon-chain/p2p/peers"
	pb "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
	log "github.com/sirupsen/logrus"
)

// MockPeersProvider implements PeersProvider for p2p.
type MockPeersProvider struct {
	lock  sync.Mutex
	peers *peers.Status
}

// Peers provides access the peer status.
func (m *MockPeersProvider) Peers() *peers.Status {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.peers == nil {
		m.peers = peers.NewStatus(5 /* maxBadResponses */)
		// Pretend we are connected to two peers
		id0, err := peer.IDB58Decode("16Uiu2HAkyWZ4Ni1TpvDS8dPxsozmHY85KaiFjodQuV6Tz5tkHVeR")
		if err != nil {
			log.WithError(err).Debug("Cannot decode")
		}
		ma0, err := ma.NewMultiaddr("/ip4/213.202.254.180/tcp/13000")
		if err != nil {
			log.WithError(err).Debug("Cannot decode")
		}
		m.peers.Add(new(enr.Record), id0, ma0, network.DirInbound)
		m.peers.SetConnectionState(id0, peers.PeerConnected)
		m.peers.SetChainState(id0, &pb.Status{FinalizedEpoch: uint64(10)})
		id1, err := peer.IDB58Decode("16Uiu2HAm4HgJ9N1o222xK61o7LSgToYWoAy1wNTJRkh9gLZapVAy")
		if err != nil {
			log.WithError(err).Debug("Cannot decode")
		}
		ma1, err := ma.NewMultiaddr("/ip4/52.23.23.253/tcp/30000/ipfs/QmfAgkmjiZNZhr2wFN9TwaRgHouMTBT6HELyzE5A3BT2wK/p2p-circuit")
		if err != nil {
			log.WithError(err).Debug("Cannot decode")
		}
		m.peers.Add(new(enr.Record), id1, ma1, network.DirOutbound)
		m.peers.SetConnectionState(id1, peers.PeerConnected)
		m.peers.SetChainState(id1, &pb.Status{FinalizedEpoch: uint64(11)})
	}
	return m.peers
}