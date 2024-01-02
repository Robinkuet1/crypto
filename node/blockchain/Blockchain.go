package blockchain

import (
	"math"
	"time"
)

type Blockchain struct {
	Blocks []Block
}

func (b *Blockchain) InitializeEmpty() {
	//TODO cleaner way
	var genesisBlock Block
	//genesisBlock.PrevHash = nil
	genesisBlock.Timestamp = time.Now().Unix()
	genesisBlock.Difficulty = 10

	var blocks []Block = []Block{}
	blocks = append(blocks, genesisBlock)

	b.Blocks = blocks
}

func (b *Blockchain) Print() {
	for i := 0; i < len(b.Blocks); i++ {
		b.Blocks[i].Print()
	}
}

func GetScore(b Block) byte {
	var hash = HashToByteArray(b.Hash())

	var i byte = 0
	for i < 255 {
		var b = hash[int(math.Floor(float64(i)/8))]
		if (b & byte(math.Pow(2, float64(int(i)%8)))) != 0 {
			return i
		}
		i++
	}
	return 0
}

func (b *Blockchain) TryMineBlock() {
	last := b.LastBlock()

	var i int64 = 0
	for score := GetScore(*last); score < last.Difficulty; score = GetScore(*last) {
		//powData = random
		last.PowData[0] = (byte)(i)
		last.PowData[1] = (byte)(i >> 8)
		last.PowData[2] = (byte)(i >> 16)
		last.PowData[3] = (byte)(i >> 24)

		i++
	}

	newBlock := new(Block)
	newBlock.PrevHash = b.LastBlock().Hash()
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Difficulty = b.CalculateDifficulty()

	b.Blocks = append(b.Blocks, *newBlock)

}

func (b *Blockchain) CalculateDifficulty() byte {
	if len(b.Blocks) < 2 {
		return 16
	}
	diff := b.Blocks[len(b.Blocks)-1].Timestamp - b.Blocks[len(b.Blocks)-2].Timestamp
	if diff < 5 {
		return b.Blocks[len(b.Blocks)-1].Difficulty + 1
	} else {
		return b.Blocks[len(b.Blocks)-1].Difficulty - 1
	}
}

func (b *Blockchain) LastBlock() *Block {
	return &b.Blocks[len(b.Blocks)-1]
}

// TODO
func (b *Blockchain) Valid() bool {
	return true
}
