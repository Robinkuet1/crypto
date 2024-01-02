package blockchain

import (
	"fmt"
	"time"
)

type Block struct {
	PrevHash Hash
	Timestamp int64
	Difficulty byte
	Transactions [64]Transaction
	PowData [32]byte
}

//Helper Functions
func Int64ToByteArray(data int64) []byte {
	//TODO cleaner solution
	var byteArr = make([]byte, 8, 8)

	byteArr[0] = (byte)(data)
	byteArr[1] = (byte)(data>>8)
	byteArr[2] = (byte)(data>>16)
	byteArr[3] = (byte)(data>>24)
	byteArr[4] = (byte)(data>>32)
	byteArr[5] = (byte)(data>>40)
	byteArr[6] = (byte)(data>>48)
	byteArr[7] = (byte)(data>>56)

	return byteArr
}

func ByteArrayToInt64(data []byte) int64 {
	if(len(data) < 8){
		return -1
	}
	

	//TODO cleaner solution
	var i int64 = 0
	i += (int64)(data[0])
	i += (int64)(data[1]) << 8
	i += (int64)(data[2]) << 16
	i += (int64)(data[3]) << 24
	i += (int64)(data[4]) << 32
	i += (int64)(data[5]) << 40
	i += (int64)(data[6]) << 48
	i += (int64)(data[7]) << 56

	return i
}


func (b *Block) getRawData() []byte {
	var rawData []byte
	//add prevHash
	rawData = append(rawData, b.PrevHash.Get()...)
	
	//add timestamp
	rawData = append(rawData, Int64ToByteArray(b.Timestamp)...)

	//add POW Data
	rawData = append(rawData, b.PowData[0:]...)

	return rawData
}

func (b *Block) Hash() Hash {
	return HashData(b.getRawData())
}

func (b *Block) Print() {
	fmt.Printf("/----------------------------------------------------------------------------\\\n")
	
	var t = time.Unix(b.Timestamp, 0)
	fmt.Printf(
		"| Timestamp: %d | %d:%d:%d %d.%d.%d                                |\n",
		b.Timestamp, 
		t.Hour(), t.Minute(), t.Second(),
		t.Day(), t.Month(), t.Year())

	//good formating
	if(b.Difficulty < 10) { 
		fmt.Printf("| Difficulty: %d                                                              |\n", b.Difficulty) 
	} else if(b.Difficulty < 100)  { 
		fmt.Printf("| Difficulty: %d                                                             |\n", b.Difficulty) 
	} else if(b.Difficulty >= 100) { 
		fmt.Printf("| Difficulty: %d                                                            |\n", b.Difficulty) 
	}


	//genesis Block
	if b.PrevHash.IsNull() {
		fmt.Printf("| prevHash: none //genesis Block                                             |\n")
	} else {
		fmt.Printf("| prevHash: %s |\n", b.PrevHash.ToString())
	}

	fmt.Printf("\\----------------------------------------------------------------------------/\n")
}