package utils

import (
	"coinme/internal/conf"
	"coinme/internal/types"
	"encoding/json"
	"os"
	"strings"

	"slices"
)

// ADD

func AddCoinJsonL(name string, value int, desc string) {
	id, _, err := conf.NextIndexes(true, false)
	CrashErr(err, conf.ErrPrefixUtils+"-AddCoin/NextIndexes")

	c, err := types.NewCoin(id, name, desc, value)
	CrashErr(err, conf.ErrPrefixUtils+"-AddCoin/NewCoin")

	root, err := os.OpenRoot(conf.DataDirPath)
	CrashErr(err, conf.ErrPrefixUtils+"-AddCoin/OpenRoot")
	defer root.Close()

	err = SaveCoinJsonLines(root, conf.DataCoinPathJsonL, *c)
	CrashErr(err, conf.ErrPrefixUtils+"-AddCoin/SaveCoinJsonL")
}

func AddChainJsonL(name string, desc string) {
	_, id, err := conf.NextIndexes(false, true)
	CrashErr(err, conf.ErrPrefixUtils+"-AddChain/NextIndexes")

	c, err := types.NewChain(id, name, desc, []int{})
	CrashErr(err, conf.ErrPrefixUtils+"-AddChain/NewChain")

	r, err := os.OpenRoot(conf.DataDirPath)
	CrashErr(err, conf.ErrPrefixUtils+"-AddCoin/OpenRoot")
	defer r.Close()

	err = SaveChainJsonLines(r, conf.DataChainPathJsonL, *c)
	CrashErr(err, conf.ErrPrefixUtils+"-AddCoin/SaveChainJsonL")
}

func SaveCoinJsonLines(r *os.Root, filename string, c types.Coin) error {
	f, err := r.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(c)
	if err != nil {
		return err
	}
	return nil
}

func SaveChainJsonLines(r *os.Root, filename string, c types.Chain) error {
	f, err := r.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(c)
	if err != nil {
		return err
	}
	return nil
}

// LIST

func LoadCoinJsonL() []types.Coin {
	f, err := os.OpenInRoot(conf.DataDirPath, conf.DataCoinPathJsonL)
	CrashErr(err, conf.ErrPrefixUtils+"-LoadCoinJsonL/OpenInRoot")
	defer f.Close()

	d := json.NewDecoder(f)

	coins := []types.Coin{}

	for d.More() {
		var c types.Coin
		err = d.Decode(&c)
		CrashErr(err, conf.ErrPrefixUtils+"-LoadCoinJsonL/Decode")

		coins = append(coins, c)
	}
	return coins
}

func LoadChainJsonL() []types.Chain {
	f, err := os.OpenInRoot(conf.DataDirPath, conf.DataChainPathJsonL)
	CrashErr(err, conf.ErrPrefixUtils+"-LoadChainJsonL/OpenInRoot")
	defer f.Close()

	d := json.NewDecoder(f)

	chains := []types.Chain{}

	for d.More() {
		var c types.Chain
		err = d.Decode(&c)
		CrashErr(err, conf.ErrPrefixUtils+"-LoadChainJsonL/Decode")

		chains = append(chains, c)
	}
	return chains
}

func LoadCoinByDescJsonL(ds []string) []types.Coin {
	f, err := os.OpenInRoot(conf.DataDirPath, conf.DataCoinPathJsonL)
	CrashErr(err, conf.ErrPrefixUtils+"LoadCoinByIdJsonL/OpenInRoot")
	defer f.Close()

	d := json.NewDecoder(f)

	coins := []types.Coin{}
	wasHit := map[int]bool{} // duplicate hit filter

	for d.More() {
		var c types.Coin
		err = d.Decode(&c)
		CrashErr(err, conf.ErrPrefixUtils+"-LoadCoinByIdJsonL/Decode")

		for i := range ds {
			if strings.Contains(strings.ToLower(c.Desc), strings.ToLower(ds[i])) && !wasHit[c.Id] {
				wasHit[c.Id] = true
				coins = append(coins, c)
			}
		}
	}
	return coins
}

func LoadCoinByNameJsonL(ns []string) []types.Coin {
	f, err := os.OpenInRoot(conf.DataDirPath, conf.DataCoinPathJsonL)
	CrashErr(err, conf.ErrPrefixUtils+"LoadCoinByIdJsonL/OpenInRoot")
	defer f.Close()

	d := json.NewDecoder(f)

	coins := []types.Coin{}
	wasHit := map[int]bool{} // duplicate hit filter

	for d.More() {
		var c types.Coin
		err = d.Decode(&c)
		CrashErr(err, conf.ErrPrefixUtils+"-LoadCoinByIdJsonL/Decode")

		for i := range ns {
			if strings.Contains(strings.ToLower(c.Name), strings.ToLower(ns[i])) && !wasHit[c.Id] {
				wasHit[c.Id] = true
				coins = append(coins, c)
			}
		}
	}
	return coins
}

func LoadCoinByValueJsonL(vs []int) []types.Coin {
	f, err := os.OpenInRoot(conf.DataDirPath, conf.DataCoinPathJsonL)
	CrashErr(err, conf.ErrPrefixUtils+"LoadCoinByIdJsonL/OpenInRoot")
	defer f.Close()

	d := json.NewDecoder(f)

	coins := []types.Coin{}
	wasHit := map[int]bool{} // duplicate hit filter

	for d.More() {
		var c types.Coin
		err = d.Decode(&c)
		CrashErr(err, conf.ErrPrefixUtils+"-LoadCoinByIdJsonL/Decode")

		for i := range vs {
			if c.Value == vs[i] && !wasHit[c.Id] {
				wasHit[c.Id] = false
				coins = append(coins, c)
			}
		}
	}
	return coins
}

func LoadCoinByIdJsonL(ids []int) []types.Coin {
	f, err := os.OpenInRoot(conf.DataDirPath, conf.DataCoinPathJsonL)
	CrashErr(err, conf.ErrPrefixUtils+"LoadCoinByIdJsonL/OpenInRoot")
	defer f.Close()

	d := json.NewDecoder(f)

	coins := []types.Coin{}
	wasHit := map[int]bool{} // duplicate hit filter

	for d.More() {
		var c types.Coin
		err = d.Decode(&c)
		CrashErr(err, conf.ErrPrefixUtils+"-LoadCoinByIdJsonL/Decode")

		for i := range ids {
			if c.Id == ids[i] && !wasHit[c.Id] {
				wasHit[c.Id] = true
				coins = append(coins, c)
			}
		}
	}
	return coins
}

func LoadChainByDescJsonL(ds []string) []types.Chain {
	f, err := os.OpenInRoot(conf.DataDirPath, conf.DataChainPathJsonL)
	CrashErr(err, conf.ErrPrefixUtils+"LoadChainByDescJsonL/OpenInRoot")
	defer f.Close()

	d := json.NewDecoder(f)

	chains := []types.Chain{}
	wasHit := map[int]bool{} // duplicate hit filter

	for d.More() {
		var c types.Chain
		err = d.Decode(&c)
		CrashErr(err, conf.ErrPrefixUtils+"-LoadChainByDescJsonL/Decode")

		for i := range ds {
			if strings.Contains(strings.ToLower(c.Desc), strings.ToLower(ds[i])) && !wasHit[c.Id] {
				wasHit[c.Id] = true
				chains = append(chains, c)
			}
		}
	}
	return chains
}

func LoadChainByNameJsonL(ns []string) []types.Chain {
	f, err := os.OpenInRoot(conf.DataDirPath, conf.DataChainPathJsonL)
	CrashErr(err, conf.ErrPrefixUtils+"LoadChainByNameJsonL/OpenInRoot")
	defer f.Close()

	d := json.NewDecoder(f)

	chains := []types.Chain{}
	wasHit := map[int]bool{} // duplicate hit filter

	for d.More() {
		var c types.Chain
		err = d.Decode(&c)
		CrashErr(err, conf.ErrPrefixUtils+"-LoadChainByNameJsonL/Decode")

		for i := range ns {
			if strings.Contains(strings.ToLower(c.Name), strings.ToLower(ns[i])) && !wasHit[c.Id] {
				wasHit[c.Id] = true
				chains = append(chains, c)
			}
		}
	}
	return chains
}

func LoadChainByRelativeIdJsonL(rids []int) []types.Chain {
	f, err := os.OpenInRoot(conf.DataDirPath, conf.DataChainPathJsonL)
	CrashErr(err, conf.ErrPrefixUtils+"LoadChainByRelativeIdJsonL/OpenInRoot")
	defer f.Close()

	d := json.NewDecoder(f)

	chains := []types.Chain{}
	wasHit := map[int]bool{} // duplicate hit filter

	for d.More() {
		var c types.Chain
		err = d.Decode(&c)
		CrashErr(err, conf.ErrPrefixUtils+"-LoadChainByRelativeIdJsonL/Decode")

		for i := range rids {
			_, ok := c.Relatives[rids[i]]
			if ok && !wasHit[c.Id] { // God said, "You Shan't be the relative of yourself"
				wasHit[c.Id] = true
				chains = append(chains, c)
			}
		}
	}
	return chains
}

func LoadChainByCoinIdJsonL(cids []int) []types.Chain {
	f, err := os.OpenInRoot(conf.DataDirPath, conf.DataChainPathJsonL)
	CrashErr(err, conf.ErrPrefixUtils+"LoadChainByCoinIdJsonL/OpenInRoot")
	defer f.Close()

	d := json.NewDecoder(f)

	chains := []types.Chain{}
	wasHit := map[int]bool{} // duplicate hit filter

	for d.More() {
		var c types.Chain
		err = d.Decode(&c)
		CrashErr(err, conf.ErrPrefixUtils+"-LoadChainByCoinIdJsonL/Decode")

		for i := range cids {
			for j := range c.Coins {
				if c.Coins[j] == cids[i] && !wasHit[c.Id] {
					wasHit[c.Id] = true
					chains = append(chains, c)
				}
			}
		}
	}
	return chains
}

func LoadChainByIdJsonL(ids []int) []types.Chain {
	f, err := os.OpenInRoot(conf.DataDirPath, conf.DataChainPathJsonL)
	CrashErr(err, conf.ErrPrefixUtils+"LoadChainByIdJsonL/OpenInRoot")
	defer f.Close()

	d := json.NewDecoder(f)

	chains := []types.Chain{}
	wasHit := map[int]bool{} // duplicate hit filter

	for d.More() {
		var c types.Chain
		err = d.Decode(&c)
		CrashErr(err, conf.ErrPrefixUtils+"-LoadChainByIdJsonL/Decode")

		for i := range ids {
			if c.Id == ids[i] && !wasHit[c.Id] {
				wasHit[c.Id] = true
				chains = append(chains, c)
			}
		}
	}
	return chains
}

func CleanCoinByIdJsonL(ids []int) []types.Coin {
	slices.Sort(ids)
	r, err := os.OpenRoot(conf.DataDirPath)
	CrashErr(err, conf.ErrPrefixUtils+"-CleanCoinByIdJsonL/OpenRoot")
	defer r.Close()

	f, err := r.OpenFile(conf.DataCoinPathJsonL, os.O_RDWR, 0644)
	CrashErr(err, conf.ErrPrefixUtils+"-CleanCoinByIdJsonL/OpenFile")
	defer f.Close()

	d := json.NewDecoder(f)

	coins := []types.Coin{}
	removed := []types.Coin{}
	// wasHit := map[int]bool{}
	// wasRemoved := map[int]bool{}

	id := 1
	once := true

	for d.More() {
		var c types.Coin
		err = d.Decode(&c)
		CrashErr(err, conf.ErrPrefixUtils+"-CleanCoinByIdJsonL/Decode")

		if once {
			id = c.Id
			once = false
		}
		match := slices.Contains(ids, c.Id)
		c.Id = id
		if match {
			removed = append(removed, c)
		} else {
			coins = append(coins, c)
			id += 1
		}
	}
	// rewrite the file with the new coins without the target one
	err = f.Truncate(0)
	f.Seek(0, 0)
	CrashErr(err, conf.ErrPrefixUtils+"-CleanCoinByIdJsonL/Truncate")

	e := json.NewEncoder(f)

	for i := range coins {
		e.Encode(coins[i])
	}

	return removed
}

func CleanChainByIdJsonL(ids []int) []types.Chain {
	slices.Sort(ids)
	r, err := os.OpenRoot(conf.DataDirPath)
	CrashErr(err, conf.ErrPrefixUtils+"-CleanChainByIdJsonL/OpenRoot")
	defer r.Close()

	f, err := r.OpenFile(conf.DataChainPathJsonL, os.O_RDWR, 0644)
	CrashErr(err, conf.ErrPrefixUtils+"-CleanChainByIdJsonL/OpenFile")
	defer f.Close()

	d := json.NewDecoder(f)

	chains := []types.Chain{}
	removed := []types.Chain{}

	id := 0

	for d.More() {
		var c types.Chain
		err = d.Decode(&c)
		CrashErr(err, conf.ErrPrefixUtils+"-CleanChainByIdJsonL/Decode")

		if id < 1 { // exec once
			id = c.Id
		}

		match := slices.Contains(ids, c.Id)
		c.Id = id
		if match {
			removed = append(removed, c)
		} else {
			chains = append(chains, c)
			id += 1
		}
	}
	// rewrite the file with the new chains without the target one
	err = f.Truncate(0)
	f.Seek(0, 0)
	CrashErr(err, conf.ErrPrefixUtils+"-CleanChainByIdJsonL/Truncate")

	e := json.NewEncoder(f)

	for i := range chains {
		e.Encode(chains[i])
	}

	return removed
}

// EDIT

func EditCoinJsonL(id int, name string, value int, desc string) types.Coin {
	// read all coins
	coins := LoadCoinJsonL()
	var coin types.Coin
	var target int

	// modify target coin, or return original coin, unmodified on error
	for i, c := range coins {
		if c.Id == id {
			if value < 1 {
				value = c.Value
			}
			if name == conf.IgnoreCoinName {
				name = c.Name
			}
			if desc == conf.IgnoreCoinDesc {
				desc = c.Desc
			}
			cptr, err := types.NewCoin(id, name, desc, value)
			if DashErr(err, "-EditCoin/NewCoin") {
				coin = c
				target = i
				break
				// no modification took place
			}
			coin = *cptr
			coin.Cdate = c.Cdate
			target = i
			break
		}
	}
	coins[target] = coin

	// write everything back
	r, err := os.OpenRoot(conf.DataDirPath)
	CrashErr(err, "-EditCoin/OpenRoot")
	defer r.Close()

	f, err := r.OpenFile(conf.DataCoinPathJsonL, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	CrashErr(err, "-EditCoin/OpenFile")
	defer f.Close()

	e := json.NewEncoder(f)

	for i := range coins {
		e.Encode(coins[i])
	}

	// return modified coin
	return coin
}

func EditChainJsonL(id int, name, desc string) types.Chain {
	// read all coins
	chains := LoadChainJsonL()
	var chain types.Chain
	var target int

	// modify target coin, or return original coin, unmodified on error
	for i, c := range chains {
		if c.Id == id {
			if name == conf.IgnoreChainName {
				name = c.Name
			}
			if desc == conf.IgnoreChainDesc {
				desc = c.Desc
			}
			cptr, err := types.NewChain(id, name, desc, c.Coins)
			if DashErr(err, "-EditChain/NewChain") {
				chain = c
				target = i
				break
				// no modification took place
			}
			chain = *cptr
			chain.Cdate = c.Cdate
			chain.Relatives = c.Relatives
			target = i
			break
		}
	}
	chains[target] = chain

	// write everything back
	r, err := os.OpenRoot(conf.DataDirPath)
	CrashErr(err, "-EditChain/OpenRoot")
	defer r.Close()

	f, err := r.OpenFile(conf.DataChainPathJsonL, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	CrashErr(err, "-EditChain/OpenFile")
	defer f.Close()

	e := json.NewEncoder(f)

	for i := range chains {
		e.Encode(chains[i])
	}

	// return modified coin
	return chain
}

func EditChainCoinsJsonL(id int, coins []int) types.Chain {
	// read all coins
	chains := LoadChainJsonL()
	var chain types.Chain
	var target int

	// modify target coin, or return original coin, unmodified on error
	for i, c := range chains {
		if c.Id == id {
			if len(coins) < 1 || coins[0] == conf.IgnoreChainCoins {
				coins = c.Coins
			}
			cptr, err := types.NewChain(id, c.Name, c.Desc, coins)
			if DashErr(err, "-EditChainCoins/NewChain") {
				chain = c
				target = i
				break
				// no modification took place
			}
			chain = *cptr
			chain.Cdate = c.Cdate
			chain.Relatives = c.Relatives
			target = i
			break
		}
	}
	chains[target] = chain

	// write everything back
	r, err := os.OpenRoot(conf.DataDirPath)
	CrashErr(err, "-EditChainCoins/OpenRoot")
	defer r.Close()

	f, err := r.OpenFile(conf.DataChainPathJsonL, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	CrashErr(err, "-EditChainCoins/OpenFile")
	defer f.Close()

	e := json.NewEncoder(f)

	for i := range chains {
		e.Encode(chains[i])
	}

	// return modified coin
	return chain
}
