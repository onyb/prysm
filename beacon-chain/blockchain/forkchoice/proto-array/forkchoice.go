package proto_array

import "github.com/prysmaticlabs/prysm/shared/params"

func (f *ForkChoice) New(justifiedEpoch uint64, finalizedEpoch uint64, finalizedRoot [32]byte) {
	f.store = &Store{
		justifiedEpoch: justifiedEpoch,
		finalizedEpoch: finalizedEpoch,
		finalizedRoot:  finalizedRoot,
		nodes:          make([]Node, 0),
		nodeIndices:    make(map[[32]byte]uint64),
	}

	f.store.Insert(finalizedRoot, params.BeaconConfig().ZeroHash, justifiedEpoch, finalizedEpoch)

	f.balances = make([]uint64, 0)
	f.votes = make([]Vote, 0)
}

func (f *ForkChoice) ProcessAttestation(validatorIndex uint64, blockRoot [32]byte, blockEpoch uint64) {
	if blockEpoch > f.votes[validatorIndex].nextEpoch {
		f.votes[validatorIndex].nextEpoch = blockEpoch
		f.votes[validatorIndex].nextRoot = blockRoot
	}
}

func (f *ForkChoice) ProcessBlock(blockRoot [32]byte, parentRoot [32]byte, finalizedEpoch uint64, justifiedEpoch uint64) {
	f.store.Insert(blockRoot, parentRoot, justifiedEpoch, finalizedEpoch)
}
