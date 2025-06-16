package types

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

func SortChains(chains []Chain) {
	sort.Sort(ByChainId(chains))
}

type ByChainId []Chain

func (a ByChainId) Len() int           { return len(a) }
func (a ByChainId) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByChainId) Less(i, j int) bool { return a[i].Id < a[j].Id }

type Chain struct {
	Id        int
	Name      string
	Coins     []int
	Relatives map[int]float64 // maps the relative chain id alongside a conversion factor. Chain ID 601 has relative chain 606 with conversion factor of 1.23, so all coins in 601 are x1.13 the value of coins in chain 606, even if it is the same coin.
	Desc      string
	Cdate     time.Time
	Mdate     time.Time
}

func NewChain(id int, name string, desc string, coins []int) (*Chain, error) {
	var c *Chain
	relatives := map[int]float64{}
	now := time.Now()

	switch {
	case id < 1:
		return nil, errors.New("invalid Chain.Id < 1")
	case len(name) == 0:
		return nil, errors.New("invalid Chain.Name empty")
	default:
		c = &Chain{id, name, coins, relatives, desc, now, now}
		return c, nil
	}
}

func NewMDate() time.Time {
	return time.Now()
}

// func (c Chain) String() string { // ssimplified for debugging
// 	return fmt.Sprintf("\n{Id:%v, Name:\"%v\", Coins:%v, Relatives:%v}\n", c.Id, c.Name, c.Coins, c.Relatives)
// }

func (c Chain) String() string {
	return fmt.Sprintf("\n{Id:%v, Name:\"%v\", Coins:%v, Relatives:%v, Desc:\"%v\", Cdate:\"%v\", Mdate:\"%v\"}\n", c.Id, c.Name, c.Coins, c.Relatives, c.Desc, c.Cdate, c.Mdate)
}
