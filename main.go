package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/Fuzzycc/coinme/ledger"
)

/*
In utility, make a Coin Marshal to BaseLedger, which is a record []string in the V1 form
Define the V1 form and abstract the code? Nah, just write it to work and refine it later.
or should the marsha function be in the ledger?
The ledger would become too big later on if I have multipel forms, each with their own marshal unmarshal

so, to write a coin... r.Write(utils.MarshalV1SCSV(coin)) V1 denonates the rules on the field.
So there can be V1SCSV for the base, V1JSON, V1XML, V1SQl, etc.

For base.
New coins are appended only, sorting is separate
So, Cleaning needs to just discard any occurence beyond the first
Because, manually editing the file is the only way to get ID duplication
Because the CLI does not allow ID insertion
*/

func main() {
	// still in development, this is verification code
	fmt.Println("Coinme, right now!")

	// Open sample file
	file, err := os.OpenInRoot("./data/", "coin.sample.txt")
	if err != nil {
		panic(err)
	}
	// defer file.Close()

	// create BaseReader, an enconding/csv Reader wrapper
	r := ledger.NewBaseReader(file)
	// l.Comma = ';'
	// l.Comment = '0'

	// Read all entries. Read also exist
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	// close the file
	file.Close()

	// open the file for writing
	root, _ := os.OpenRoot("./data/")
	f, err := root.OpenFile("coin.sample.scsv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0o644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// create tmp file as a data store
	// f, err := os.CreateTemp("", "coinme-tmpfile-*.txt")
	// defer f.Close()
	// defer os.Remove(f.Name())
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// write to the file
	w := ledger.NewBaseWriter(f)
	w.SetComma(';')
	w.WriteAll(records)

	f.Seek(0, 0)
	r = ledger.NewBaseReader(f)

	// Read from the file
	records, err = r.ReadAll()
	if err != nil {
		panic(err)
	}

	// display results
	display := func() {
		w := tabwriter.NewWriter(os.Stdout, 1, 1, 4, ' ', 0)
		defer w.Flush() //<<< printing is here
		for _, record := range records {
			fmt.Fprintf(w, "%+v\t\"%+v\"\t%+v\n", record[1], record[2], record[3])
		}
		fmt.Println(f.Name())
	}
	display()
}

// type (
// 	// Abstraction for the underlying coin Id type
// 	CoinID uint16
// 	// Abstraction for the underlying coin Value type
// 	CoinValue uint32
// 	// Abstraction for the underlying coin Name type
// 	CoinName string
// 	// Abstraction for the underlying chain Id type
// 	ChainID uint16
// 	// Abstraction for the underlying chain Name type
// 	ChainName string
// )

// // A coin has a Value and a Name.
// type coin struct {
// 	Id    CoinID
// 	Value CoinValue
// 	Name  CoinName
// }

// // Underlying Data structure
// // Coin ID, Value, Name

// // A chain is a named collection of coins by their Id value.
// type chain struct {
// 	Id    ChainID
// 	Name  ChainName
// 	Coins []CoinID
// }

// type coinTable []coin

// type chainTable []chain

// // Init methods---------------------------
// func NewCoin(s CoinName, v CoinValue) coin
// func NewChain(s ChainName, coins ...any) chain

// // Chain methods---------------------------

// func (c chain) Add(...any) error // for each coin, if switch type coin then add coin.id, if uint then add the id

// func (c chain) Remove(any) error

// func (c chain) Modify(any, uint32, string) error

// // No read methods, fields are exported and public

// func (c chain) Convert(any, any) (int, error) // accepts coin, coin.id, or Coins[] index

// func (c chain) String() string

// // Coin methods----------------------------

// func (c coin) New(string, uint32)

// func (c coin) Modify(string, uint32)

// // coinTable methods-------------------------
// func (ct coinTable) Last() uint16

// func (ct coinTable) First() uint16

// func (ct coinTable) Exist(any) bool

// // coinChain methods-------------------------
// func (ct chainTable) Last() uint16

// func (ct chainTable) First() uint16

// func (ct chainTable) Exist(any) bool
