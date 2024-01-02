package blockchain

type Transaction struct {
	Sender Hash
	Reciever Hash
	Signature Hash
	Amount int
}

