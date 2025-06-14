package types

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

// will sort the Coins
func SortCoins(coins []Coin) {
	sort.Sort(ByCoinId(coins))
}

type ByCoinId []Coin

func (a ByCoinId) Len() int           { return len(a) }
func (a ByCoinId) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCoinId) Less(i, j int) bool { return a[i].Id < a[j].Id }
func (a ByCoinId) Contains(i int) bool {
	for _, c := range a {
		if c.Id == i {
			return true
		}
	}
	return false
}

type Coin struct {
	Id    int
	Name  string
	Value int
	Desc  string
	Cdate time.Time
	Mdate time.Time
}

func NewCoin(id int, name string, desc string, value int) (*Coin, error) {
	var c *Coin
	now := time.Now()

	switch {
	case id < 1:
		return nil, errors.New("invalid Coin.Id < 1")
	case len(name) == 0:
		return nil, errors.New("invalid Coin.Name empty")
	case value < 1:
		return nil, errors.New("invalid Coin.Value < 1")
	default:
		c = &Coin{id, name, value, desc, now, now}
		return c, nil
	}
}

func (c Coin) String() string {
	return fmt.Sprintf("\n{Id:%v, Name:\"%v\", Value:%v, Desc:\"%v\", Cdate:\"%v\", Mdate:\"%v\"}\n", c.Id, c.Name, c.Value, c.Desc, c.Cdate, c.Mdate)
}
