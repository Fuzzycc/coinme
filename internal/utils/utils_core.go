package utils

import (
	"bytes"
	"io"
)

// type DataFormat int

// const (
// 	Json DataFormat = iota
// 	Csv
// 	Toml
// 	Bolt
// 	Sql
// 	Default
// )

// var (
// 	toDataFormat = map[string]DataFormat{
// 		"json":    Json,
// 		"csv":     Csv,
// 		"toml":    Toml,
// 		"bolt":    Bolt,
// 		"sql":     Sql,
// 		"Default": Json, // Default set here
// 	}
// )

// func parseDataFormat(s string) DataFormat {
// 	var df DataFormat
// 	var ok bool
// 	s = strings.ToLower(s)
// 	df, ok = toDataFormat[s]
// 	if !ok {
// 		df = Default
// 	}
// 	return df
// 	// what i should do instead with -type flag is set
// 	// determine whom we are checking coin or chain files from a lever bool variable set in main.go by each command
// 	// then compare type of the file with the -type flag's value
// 	// if same, do nothing and return true
// 	// if different convert the entire file to the type specified by -type flag
// 	// then return true
// 	// after returning true, the caller function like AddCoin redirects to addCoinTYPE like addCoinJSON, etc.

// 	// for now, working without implementing this.

// 	// ALL OF THIS SHOULD BE ADDED BEFORE RETURNING DF TO THE CALLER FUNCTION
// }

// func IsFloatZero(f float64, tolerance float64) bool {
// 	if f <= 0 { // negative
// 		return (f * -1) <= tolerance
// 	} else { // positive
// 		return f < tolerance
// 	}
// }

func AddCoinHandler(name string, value int, desc string) {
	switch value {
	case 0:
		LogErr("coinme: add coin: no command specified, running list coin instead")
		ListCoin("", "")
	default:
		AddCoinJsonL(name, value, desc)
	}
	// figure out which format we are adding coin to
	// for now, just a placeholder
}

func AddChainHandler(name string, desc string) {
	AddChainJsonL(name, desc)
}

func ListCoin(by string, term any) {
	switch by {
	case "id":
		ids := term.([]int)
		coins := LoadCoinByIdJsonL(ids)
		PrintOut(coins)
	case "value":
		vs := term.([]int)
		coins := LoadCoinByValueJsonL(vs)
		PrintOut(coins)
	case "name":
		ns := term.([]string)
		// PrintOut(ns)
		coins := LoadCoinByNameJsonL(ns)
		PrintOut(coins)
	case "desc":
		ds := term.([]string)
		coins := LoadCoinByDescJsonL(ds)
		PrintOut(coins)
	default:
		LogErr("coinme: list coin: no command specified, running default behavior")
		coins := LoadCoinJsonL()
		PrintOut(coins)
	}
}

func ListChain(by string, term any) {
	switch by {
	case "id":
		ids := term.([]int)
		chains := LoadChainByIdJsonL(ids)
		PrintOut(chains)
	case "cid":
		cids := term.([]int)
		chains := LoadChainByCoinIdJsonL(cids)
		PrintOut(chains)
	case "name":
		ns := term.([]string)
		chains := LoadChainByNameJsonL(ns)
		PrintOut(chains)
	case "desc":
		ds := term.([]string)
		chains := LoadChainByDescJsonL(ds)
		PrintOut(chains)
	case "rid":
		rids := term.([]int)
		chains := LoadChainByRelativeIdJsonL(rids)
		PrintOut(chains)
	default:
		LogErr("coinme: list chain: no command specified, running default behavior")
		chains := LoadChainJsonL()
		PrintOut(chains)
	}
}

func RemoveCoin(by string, term any) {
	switch by {
	case "id":
		ids := term.([]int)
		coins := CleanCoinByIdJsonL(ids)
		PrintOut(coins)
	default:
		LogErr("coinme: remove coin: no command specified, aborting...")
		PrintOut("incomplete command, aborting...")
	}
}

func RemoveChain(by string, term any) {
	switch by {
	case "id":
		ids := term.([]int)
		chains := CleanChainByIdJsonL(ids)
		PrintOut(chains)
	default:
		LogErr("Coinme: remove chain: no command specified, aborting...")
		PrintOut("incomplete command, aborting...")
	}
}

func EditCoin(id, value int, name, desc string) {
	switch {
	case id > 0:
		coin := EditCoinJsonL(id, name, value, desc)
		PrintOut(coin)
	default:
		LogErr("coinme: edit coin: no id specified, running list coin instead")
		ListCoin("", "")
	}
}

func EditChain(id int, name, desc string){
	switch {
	case id > 0:
		chain := EditChainJsonL(id, name, desc)
		PrintOut(chain)
	default:
		LogErr("coinme: edit chain: no id specified, running list chain instead")
		ListChain("", "")
	}
}

func EditChainCoins(id int, coins []int) {
	switch {
	case id > 0 :
		chain := EditChainCoinsJsonL(id, coins)
		PrintOut(chain)
	default:
		LogErr("coinme: edit chain coins: no coinId specified, running list on chain instead")
		ListChain("id", id)
	}
}

// --- --- --- ---
// Reader
// --- --- --- ---

func IORead(input io.Reader) (string, error) {
	buf := make([]byte, 1024)
	var output bytes.Buffer

	for {
		n, err := input.Read(buf)

		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		output.Write(buf[:n])
	}
	return output.String(), nil
}

func IOReadAll(input io.Reader) (string, error) {
	readData, err := io.ReadAll(input)
	CrashErr(err, "coinme-utils")

	return string(readData), nil
}

// --- --- --- ---
// Writer
// --- --- --- ---

func IOWrite(input string, w io.Writer) (int, error) {
	p := ([]byte)(input)
	n, err := w.Write(p)
	if err != nil {
		return 0, err
	}
	return n, nil
}
