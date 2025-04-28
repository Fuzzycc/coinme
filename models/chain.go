package models

type (
	// Abstraction for the underlying chain Id type
	ChainID uint16
	// Abstraction for the underlying chain Name type
	ChainName string
)

// A chain is a named collection of coins by their Id value.
type Chain struct {
	Id    ChainID
	Name  ChainName
	Coins []CoinID
}
