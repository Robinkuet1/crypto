package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

type Hash struct {
	data [32]byte
} 

func (h *Hash) IsNull() bool {
	for i := 0; i < 32; i++ {
		if (h.data[i] != 0) {
			return false
		}
	}
	return true
}

func (h *Hash) ToString() string {
	//TODO cleaner way from array to slice
	data := append([]byte{}, h.data[0:]...)

	return hex.EncodeToString(data)
}

func (h *Hash) Get() []byte {
	//TODO cleaner way from array to slice
	return h.data[0:]
}

func HashToByteArray(h Hash) []byte {
	//TODO cleaner way from array to slice
	return h.data[0:]
}

func HashData(data []byte) Hash {
	hash := new(Hash)
	hash.data = sha256.Sum256(data)
	return *hash
}