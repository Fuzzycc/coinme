package models

import "fmt"

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

// A short hand for id;name;value
func (c *Coin) String() string {
	return fmt.Sprintf("%d;%s;%d", c.Id, c.Name, c.Value)
}

func (cid *CoinID) String() string {
	return fmt.Sprintf("%d", cid)
}

func (cn *CoinName) String() string {
	return string(*cn)
}

func (cv *CoinValue) String() string {
	return fmt.Sprintf("%d", cv)
}

// Underlying Data structure
// Coin ID, Value, Name
