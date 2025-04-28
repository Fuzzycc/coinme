package models

type (
	// Abstraction for the underlying coin Id type
	CoinID uint16
	// Abstraction for the underlying coin Value type
	CoinValue uint32
	// Abstraction for the underlying coin Name type
	CoinName string
)

// A coin has a Value and a Name.
type Coin struct {
	Id    CoinID
	Value CoinValue
	Name  CoinName
}

// Underlying Data structure
// Coin ID, Value, Name
