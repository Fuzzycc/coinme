package main

import (
	"fmt"
)

func main() {
	fmt.Println("Coinme, right now!")
}

type (
	// Abstraction for the underlying coin Id type
	CoinID uint16
	// Abstraction for the underlying coin Value type
	CoinValue uint32
	// Abstraction for the underlying coin Name type
	CoinName string
	// Abstraction for the underlying chain Id type
	ChainID uint16
	// Abstraction for the underlying chain Name type
	ChainName string
)

// A coin has a Value and a Name.
type coin struct {
	Id    CoinID
	Value CoinValue
	Name  CoinName
}
// Underlying Data structure
// Coin ID, Value, Name

// A chain is a named collection of coins by their Id value.
type chain struct {
	Id    ChainID
	Name  ChainName
	Coins []CoinID
}

type coinTable []coin

type chainTable []chain

// Init methods---------------------------
func NewCoin(s CoinName, v CoinValue) (coin)
func NewChain(s ChainName,coins ...any) (chain)


// Chain methods---------------------------

func (c chain) Add(...any) error // for each coin, if switch type coin then add coin.id, if uint then add the id

func (c chain) Remove(any) error

func (c chain) Modify(any, uint32, string) error

// No read methods, fields are exported and public

func (c chain) Convert(any, any) (int, error) // accepts coin, coin.id, or Coins[] index

func (c chain) String() string

// Coin methods----------------------------

func (c coin) New(string, uint32)

func (c coin) Modify(string, uint32)

// coinTable methods-------------------------
func (ct coinTable) Last() uint16

func (ct coinTable) First() uint16

func (ct coinTable) Exist(any) bool

// coinChain methods-------------------------
func (ct chainTable) Last() uint16

func (ct chainTable) First() uint16

func (ct chainTable) Exist(any) bool
