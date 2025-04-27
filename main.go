package main

import "fmt"

func main() {
	fmt.Println("Coinme, right now!")
}

// A coin has a positive integer Value and a Name.
type coin struct {
	Id    uint16
	Value uint32
	Name  string
}

// Underlying Data structure
// Coin ID, Value, Name

// A chain is a named collection of coins by their Id value.
type chain struct {
	Id    uint16
	Name  string
	Coins []uint16
}

type coinTable []coin

type chainTable []chain

// Chain methods---------------------------
func (c chain) New(string, ...any) // figure out the ID by yourself

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
