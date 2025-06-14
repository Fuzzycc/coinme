package utils

import (
	"bufio"
	"coinme/internal/types"
	"encoding/json"
	"os"
)

// Appends c to f
func SaveCoinJSON(f *os.File, c types.Coin) error {

	b, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}

	_, err = f.Seek(1, 2) // append to file
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	_, err = w.Write(b)
	if err != nil {
		return err
	}

	err = w.Flush()
	if err != nil {
		return err
	}

	return nil
}

// func SaveCoinJSONM(f *os.File, c types.Coin) error {
// 	b, e
// }

// func ReadCoinJSON(f *os.File) ([]types.Coin) {
// 	dec := json.NewDecoder(f)
// 	data, err := dec.Decode()
// }
