package utils

import (
	"coinme/models"
	"fmt"
)

type version uint8

const (
	Version1 version = iota + 1
)

func MarshalCoin(c models.Coin, v version) string {
	switch v {
	case Version1:
		return marshalCoinV1(c)
	default:
		return ""
	}
}

func marshalCoinV1(c models.Coin) string {
	return fmt.Sprintf("%s;%s;%s", c.Id, c.Name, c.Value)
}
