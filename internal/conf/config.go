package conf

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

const (
	DataDirPath        string = "./data"
	DataCoinPathJsonL   string = "coin.jsonl"
	DataChainPathJsonL  string = "chain.jsonl"
	DataExternalConfig string = "external.json"

	ConfigDir  string = "./internal/conf"
	ConfigPath string = "config.json"

	ErrPrefix      string = "coinme"
	ErrPrefixUtils string = "coinme-utils"
	ErrPrefixCli   string = "coinme-cli"
	ErrPrefixTypes string = "coinme-types"

	IgnoreCoinName string = "x"
	IgnoreCoinDesc string = "x"

	IgnoreChainName string = "x"
	IgnoreChainDesc string = "x"
	IgnoreChainCoins int = 0
)

type ExternalConfig struct {
	NextCoinIndex  int `json:"NextCoinIndex"`
	NextChainIndex int `json:"NextChainIndex"`
}

// Reads NextCoinIndex and NextChainIndex from external.json.
// Increments them by 1 depending on ci and si respectively BEFORE returning.
//
// If err != nil, returned ints are 0, 0
//
// Example: If you need the next Coin Index whe adding a coin, use NextIndexes(true,false).
func NextIndexes(ci bool, si bool) (nextCoinId, nextChainId int, err error) {
	r, err := os.OpenRoot(DataDirPath)
	if err != nil {
		return 0, 0, err
	}
	defer r.Close()
	f, err := r.OpenFile(DataExternalConfig, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()

	var out bytes.Buffer
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, 0, err
		}
		_, err = out.Write(buf[:n])
		if err != nil {
			return 0, 0, err
		}
	}
	data := out.Bytes()

	Cfg := ExternalConfig{}
	err = json.Unmarshal(data, &Cfg)
	if err != nil {
		return 0, 0, err
	}

	if ci {
		Cfg.NextCoinIndex += 1
	}
	if si {
		Cfg.NextChainIndex += 1
	}

	data, err = json.MarshalIndent(Cfg, "", "\t")
	if err != nil {
		return 0, 0, err
	}

	f.Truncate(0)
	f.Seek(0, 0)
	f.Write(data)

	return Cfg.NextCoinIndex, Cfg.NextChainIndex, nil
}
