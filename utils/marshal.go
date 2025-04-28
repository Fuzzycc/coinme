package utils

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Fuzzycc/coinme/models"
)

type version uint8

const (
	Ignore version = iota
	Version1
)

// Serialize c into v-compliant raw, returning empty string on failure
func MarshalCoin(c models.Coin, v version) (raw string) {
	var err error
	switch v {
	case Version1:
		raw, err = marshalCoinV1(c)
	default:
		raw, err = "", errors.New("coinme/utils.MarshalCoin: Invalid Version Number")
	}
	if err != nil {
		// do something ?
	}
	return raw // ultimately ignore returning error to reduce coupling?
}

// Serialize c into Version1-Compliant raw, returning e on failure.
func marshalCoinV1(c models.Coin) (raw string, e error) {
	switch {
	case c.Id <= 0:
		raw, e = "", errors.New("coinme/utils.marshalCoinV1: Invalid Coin ID")
		return
	case c.Value <= 0:
		raw, e = "", errors.New("coinme/utils.marshalCoinV1: Invalid Coin Value")
		return
	default:
		// raw, e = fmt.Sprintf("%d;%d;%s;%d", Version1, c.Id, c.Name, c.Value), nil
		raw, e = (string(Version1) + c.String() + ""), nil
		return
	}
}

// De-serialize raw into coin using underlying, Version-dependentent functions
//
// All of coin fields are zero-valued on failure, while only coin.Name **may** be empty on success.
//
// Checking coin.Id is sufficient to indicate success or failure.
func UnmarshalCoin(raw string) (coin *models.Coin) {
	coin = &models.Coin{Id: 0, Value: 0, Name: ""}
	switch {
	case len(raw) == 0:
		return // == return coin
		// return &models.Coin{}, errors.New("coinme/utils.UnmarshalCoin: Empty string")
	case raw[0] == byte(Ignore):
		return
		// return &models.Coin{}, errors.New("coinme/utils.UnmarshalCoin: Invalid string (Version=Ignore)")
	case raw[0] == byte(Version1):
		// vs := fmt.Sprintf("%d;", Version1)
		vs := string(Version1) + ";"
		return UnmarshalCoinV1(strings.Replace(raw, vs, "", 1))
	default:
		return
		// return &models.Coin{}, errors.New("coinme/utils.UnmarshalCoin: Invalid string")
	}
}

// A V1 Pure string is ID;Name;Value

// De-serialize pure (a version-trimmed V1 string) into coin.
//
// All of coin fields are zero-valued on failure, while only coin.Name **may** be empty on success.
//
// Checking coin.Id is sufficient to indicate success or failure.
func UnmarshalCoinV1(pure string) (coin *models.Coin) {
	coin = &models.Coin{Id: 0, Value: 0, Name: ""}
	// var err error

	if pure[len(pure)-1] == ';' {
		pure = strings.Trim(pure, ";")
	} // in case the string ends or begins in ; actually won't get here if it starts with ;

	s := strings.Split(pure, ";")
	if len(s) != 3 {
		// return coin, errors.New("coinme/utils.UnmarshalCoinV1: Invalid V1 pure string")
		return coin
	}

	id, nm, vl := castCoinId(s[0]), castCoinName(s[1]), castCoinValue(s[2])
	switch {
	case id <= 0 || vl <= 0:
		// err = errors.New("coinme/utils.UnmarshalCoinV1: Invalid ID (1st Parameter)")
		return coin
	// case vl <= 0:
	// 	// err = errors.New("coinme/utils.UnmarshalCoinV1: Invalid Value (3rd Parameter)")
	case nm == "":
		coin.Id = id
		coin.Value = vl
		coin.Name = "-"
		// err = nil
	default:
		coin.Id = id
		coin.Value = vl
		coin.Name = nm
		// err = nil
	}
	return coin
	// return coin, err
}

// Attempts to convert s into coinID.
// Returns 0 on failure.
func castCoinId(s string) (coinID models.CoinID) {
	nb, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
		// 	return 0, fmt.Errorf("Conversion error: %v", err)
	}
	coinID = models.CoinID(nb)
	return
}

// Attempts to convert s into coinName.
// Does not fail, son.
func castCoinName(s string) (coinName models.CoinName) {
	coinName = models.CoinName(s)
	return
}

// Attempts to convert s into a coinValue.
// Returns 0 on failure.
func castCoinValue(s string) (coinValue models.CoinValue) {
	nb, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
		// return 0, fmt.Errorf("Conversion error: %v", err)
	}
	coinValue = models.CoinValue(nb)
	return
}
