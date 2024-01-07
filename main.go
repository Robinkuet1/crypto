package main

import (
	//"github.com/robinkuet1/crypto-node/node"
	"github.com/robinkuet1/crypto-node/node/blockchain"
	"github.com/robinkuet1/crypto-node/node/miner"
	//"fmt"
)

func main() {
	var b blockchain.Blockchain
	b.InitializeEmpty()

	var control miner.MinerControlData = miner.MinerControlData{Blockchain: &b}

	miner.StartMining(&control)

}
