package miner

import (
	"encoding/binary"
	"sync"
	"time"

	b "github.com/robinkuet1/crypto-node/node/blockchain"
)

type MinerControlData struct {
	Blockchain *b.Blockchain
	Lock       sync.Mutex
}

func int64ToByteArray(value int64) [32]byte {
	var byteArray [32]byte
	binary.BigEndian.PutUint64(byteArray[:], uint64(value))
	return byteArray
}

func StartMining(control *MinerControlData) {
	var i int64 = 0
	for {
		var block = control.Blockchain.LastBlock()
		block.PowData = int64ToByteArray(i)
		var hash = block.Hash()
		if b.GetScore(hash) >= block.Difficulty {
			block.Print()
			control.Blockchain.Blocks = append(control.Blockchain.Blocks, b.Block{
				PrevHash:   hash,
				Timestamp:  time.Now().Unix(),
				Difficulty: block.Difficulty,
				//TODO add transaction for mining reward
			})
			i = 0
		}
		i++
	}
}
