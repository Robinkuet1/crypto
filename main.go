package main

import (
    //"github.com/robinkuet1/crypto-node/node"
    "github.com/robinkuet1/crypto-node/node/blockchain"
    //"fmt"
)

func main() {
    var b blockchain.Blockchain
    b.InitializeEmpty()
    for {
        b.TryMineBlock()
        b.LastBlock().Print()
    }
}
