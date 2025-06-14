package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"

	ut "coinme/internal/utils"
)

const (
	version string = "0.1.5"
)

type funcs struct{}

func (f *funcs) defaultUsage() {
	ut.Print("Coinme\t\t\t(v" + version + ")\n- CRUD coins and coinage chains!")
}

func (f *funcs) listCoin(by string, term any) {
	ut.ListCoin(by, term)
}

func (f *funcs) listChain(by string, term any) {
	ut.ListChain(by, term)
}

func (f *funcs) addCoin(name string, value int, desc string) {
	ut.AddCoinHandler(name, value, desc)
}

func (f *funcs) addChain(name string, desc string) {
	ut.AddChainHandler(name, desc)
}

func (f *funcs) removeCoin(by string, term any) {
	ut.RemoveCoin(by, term)
}

func (f *funcs) removeChain(by string, term any) {
	ut.RemoveChain(by, term)
}

func (f *funcs) editCoin(id, value int, name, desc string) {
	ut.EditCoin(id, value, name, desc)
}

func (f *funcs) editChain(id int, name string, desc string) {
	ut.EditChain(id, name, desc)
}

func (f *funcs) editChainCoins(id int, coins []int) {
	ut.EditChainCoins(id, coins)
}

func (f *funcs) editChainRelatives(id int, chains []int) {
	ut.PrintOut("Run: edit chain relatives")
}

func main() {
	var (
		addCoinName  string
		addCoinValue int
		addCoinDesc  string

		addChainName string
		addChainDesc string

		// listCoinField string
		// listCoinBy   string

		listCoinNum  []int
		listCoinTerm []string

		listChainNum  []int
		listChainTerm []string

		removeCoinNum  []int
		removeChainNum []int

		editCoinId    int
		editCoinName  string
		editCoinValue int
		editCoinDesc  string

		editChainId        int
		editChainName      string
		editChainDesc      string
		editChainCoins     []int
		editChainRelatives []int
	)
	f := new(funcs)

	listCommands := []*cli.Command{
		{
			Name:    "coin",
			Aliases: []string{"c"},
			Usage:   "List coin",
			Commands: []*cli.Command{
				{
					Name:     "id",
					Aliases:  []string{"i"},
					Usage:    "List coins by matching id(s)",
					Category: "coin",
					Arguments: []cli.Argument{
						&cli.IntArgs{
							Name:        "integer",
							Min:         0,
							Max:         -1,
							Value:       0,
							Destination: &listCoinNum,
						},
					},
					Action: func(ctx context.Context, cmd *cli.Command) error {
						f.listCoin("id", listCoinNum)
						return nil
					},
				},
				{
					Name:     "value",
					Aliases:  []string{"v"},
					Usage:    "List coins by matching value(s)",
					Category: "coin",
					Arguments: []cli.Argument{
						&cli.IntArgs{
							Name:        "integer",
							Min:         0,
							Max:         -1,
							Value:       0,
							Destination: &listCoinNum,
						},
					},
					Action: func(ctx context.Context, cmd *cli.Command) error {
						f.listCoin("value", listCoinNum)
						return nil
					},
				},
				{
					Name:     "name",
					Aliases:  []string{"n"},
					Usage:    "List coins by matching name(s)",
					Category: "coin",
					Arguments: []cli.Argument{
						&cli.StringArgs{
							Name:        "term",
							Min:         0,
							Max:         -1,
							Value:       "",
							Destination: &listCoinTerm,
							Config:      cli.StringConfig{TrimSpace: true},
						},
					},
					Action: func(ctx context.Context, cmd *cli.Command) error {
						f.listCoin("name", listCoinTerm)
						return nil
					},
				},
				{
					Name:     "desc",
					Aliases:  []string{"d"},
					Usage:    "List coins by matching description(s)",
					Category: "coin",
					Arguments: []cli.Argument{
						&cli.StringArgs{
							Name:        "term",
							Min:         0,
							Max:         -1,
							Value:       "",
							Destination: &listCoinTerm,
							Config:      cli.StringConfig{TrimSpace: true},
						},
					},
					Action: func(ctx context.Context, cmd *cli.Command) error {
						f.listCoin("desc", listCoinTerm)
						return nil
					},
				},
			},
			Category: "list",
			Action: func(ctx context.Context, cmd *cli.Command) error {
				f.listCoin("", "")
				return nil
			},
		},
		{
			Name:    "chain",
			Aliases: []string{"s"},
			Usage:   "List chain",
			Commands: []*cli.Command{
				{
					Name:     "id",
					Aliases:  []string{"i"},
					Usage:    "List chain by matching id(s)",
					Category: "chain",
					Arguments: []cli.Argument{
						&cli.IntArgs{
							Name:        "integer",
							Min:         0,
							Max:         -1,
							Value:       0,
							Destination: &listChainNum,
						},
					},
					Action: func(ctx context.Context, cmd *cli.Command) error {
						f.listChain("id", listChainNum)
						return nil
					},
				},
				{
					Name:     "coinID",
					Aliases:  []string{"c"},
					Usage:    "List chain by matching coin id(s)",
					Category: "chain",
					Arguments: []cli.Argument{
						&cli.IntArgs{
							Name:        "integer",
							Min:         0,
							Max:         -1,
							Value:       0,
							Destination: &listChainNum,
						},
					},
					Action: func(ctx context.Context, cmd *cli.Command) error {
						f.listChain("cid", listChainNum)
						return nil
					},
				},
				{
					Name:     "relativeID",
					Aliases:  []string{"r"},
					Usage:    "List chain by matching relative chain id(s)",
					Category: "chain",
					Arguments: []cli.Argument{
						&cli.IntArgs{
							Name:        "integer",
							Min:         0,
							Max:         -1,
							Value:       0,
							Destination: &listChainNum,
						},
					},
					Action: func(ctx context.Context, c *cli.Command) error {
						f.listChain("rid", listChainNum)
						return nil
					},
				},
				{
					Name:     "name",
					Aliases:  []string{"n"},
					Usage:    "List chain by matching name(s)",
					Category: "chain",
					Arguments: []cli.Argument{
						&cli.StringArgs{
							Name:        "term",
							Min:         0,
							Max:         -1,
							Value:       "",
							Destination: &listChainTerm,
							Config:      cli.StringConfig{TrimSpace: true},
						},
					},
					Action: func(ctx context.Context, cmd *cli.Command) error {
						f.listChain("name", listChainTerm)
						return nil
					},
				},
				{
					Name:     "desc",
					Aliases:  []string{"d"},
					Usage:    "List chain by matching description(s)",
					Category: "chain",
					Arguments: []cli.Argument{
						&cli.StringArgs{
							Name:        "term",
							Min:         0,
							Max:         -1,
							Value:       "",
							Destination: &listChainTerm,
							Config:      cli.StringConfig{TrimSpace: true},
						},
					},
					Action: func(ctx context.Context, cmd *cli.Command) error {
						f.listChain("desc", listChainTerm)
						return nil
					},
				},
			},
			Category: "list",
			Action: func(ctx context.Context, cmd *cli.Command) error {
				f.listChain("", "")
				return nil
			},
		},
	}

	addCommands := []*cli.Command{
		{
			Name:    "coin",
			Aliases: []string{"c"},
			Usage:   "Add a new coin",
			Arguments: []cli.Argument{
				&cli.StringArg{
					Name:        "name ",
					Destination: &addCoinName,
				},
				&cli.IntArg{
					Name:        "value ",
					Destination: &addCoinValue,
				},
				&cli.StringArg{
					Name:        "desc",
					Destination: &addCoinDesc,
				},
			},
			Action: func(ctx context.Context, cmd *cli.Command) error {
				f.addCoin(addCoinName, addCoinValue, addCoinDesc)
				return nil
			},
		},
		{
			Name:    "chain",
			Aliases: []string{"s"},
			Usage:   "Add a new coinage chain",
			Arguments: []cli.Argument{
				&cli.StringArg{
					Name:        "name ",
					Destination: &addChainName,
				},
				&cli.StringArg{
					Name:        "desc",
					Destination: &addChainDesc,
				},
			},
			Action: func(ctx context.Context, cmd *cli.Command) error {
				f.addChain(addChainName, addChainDesc)
				return nil
			},
		},
	}

	removeCommands := []*cli.Command{
		{
			Name:      "coin",
			Aliases:   []string{"c"},
			Usage:     "Remove coin by matching id(s)",
			UsageText: "Coinme remove coin [integer ...]",
			Arguments: []cli.Argument{
				&cli.IntArgs{
					Name:        "integer",
					Min:         0,
					Max:         -1,
					Destination: &removeCoinNum,
				},
			},
			Action: func(ctx context.Context, cmd *cli.Command) error {
				f.removeCoin("id", removeCoinNum)
				return nil
			},
		},
		{
			Name:      "chain",
			Aliases:   []string{"s"},
			Usage:     "Remove chain by matching id(s)",
			UsageText: "Coinme remove coin [integer ...]",
			Arguments: []cli.Argument{
				&cli.IntArgs{
					Name:        "Integer",
					Min:         0,
					Max:         -1,
					Destination: &removeChainNum,
				},
			},
			Action: func(ctx context.Context, cmd *cli.Command) error {
				f.removeChain("id", removeChainNum)
				return nil
			},
		},
	}

	editCommands := []*cli.Command{
		{
			Name:      "coin",
			Aliases:   []string{"c"},
			Usage:     "Edit coin by matching id",
			UsageText: "Coinme edit coin {id} [name] [value] [desc]",
			Arguments: []cli.Argument{
				&cli.IntArg{
					Name:        "id",
					Value:       0,
					Destination: &editCoinId,
				},
				&cli.IntArg{
					Name:        "value",
					Value:       0,
					Destination: &editCoinValue,
				},
				&cli.StringArg{
					Name:        "name",
					Value:       "x",
					Destination: &editCoinName,
					Config:      cli.StringConfig{TrimSpace: true},
				},
				&cli.StringArg{
					Name:        "desc",
					Value:       "x",
					Destination: &editCoinDesc,
					Config:      cli.StringConfig{TrimSpace: true},
				},
			},
			Action: func(ctx context.Context, c *cli.Command) error {
				f.editCoin(editCoinId, editCoinValue, editCoinName, editCoinDesc)
				return nil
			},
		},
		{
			Name:      "chain",
			Aliases:   []string{"s"},
			Usage:     "Edit chain by matching id",
			UsageText: "Coinme edit chain {id} [name] [desc]\nCoinme edit chain coins [coinId...]\nCoinme edit chain relatives [chainId...]",
			Arguments: []cli.Argument{
				&cli.IntArg{
					Name:        "id",
					Value:       0,
					Destination: &editChainId,
				},
				&cli.StringArg{
					Name:        "name",
					Value:       "x",
					Destination: &editChainName,
				},
				&cli.StringArg{
					Name:        "desc",
					Value:       "x",
					Destination: &editChainDesc,
				},
			},
			Action: func(ctx context.Context, c *cli.Command) error {
				f.editChain(editChainId, editChainName, editChainDesc)
				return nil
			},
			Commands: []*cli.Command{
				{
					Name:      "coins",
					Aliases:   []string{"c"},
					Usage:     "Edit chain coins by ids (Overwrites)",
					UsageText: "Coinme edit chain coins id [coinId...]",
					Arguments: []cli.Argument{
						&cli.IntArg{
							Name:        "id",
							Value:       0,
							Destination: &editChainId,
						},
						&cli.IntArgs{
							Name:        "coinId",
							Min:         0,
							Max:         -1,
							Destination: &editChainCoins,
						},
					},
					Action: func(ctx context.Context, c *cli.Command) error {
						f.editChainCoins(editChainId, editChainCoins)
						return nil
					},
				},
				{
					Name:      "relatives",
					Aliases:   []string{"r"},
					Usage:     "Edit chain relative chains by ids (Overwrites)",
					UsageText: "Coinme edit chain relatives [chainId...]",
					Arguments: []cli.Argument{
						&cli.IntArg{
							Name:        "id",
							Value:       0,
							Destination: &editChainId,
						},
						&cli.IntArgs{
							Name:        "chainId",
							Min:         0,
							Max:         -1,
							Destination: &editChainRelatives,
						},
					},
					Action: func(ctx context.Context, c *cli.Command) error {
						// TODO
						f.editChainRelatives(editChainId, editChainRelatives)
						return nil
					},
				},
			},
		},
	}

	commands := []*cli.Command{
		{ // --- --- --- --- --- --- --- --- LIST
			Name:     "list",
			Aliases:  []string{"l"},
			Usage:    "List coins (d) or chains",
			Commands: listCommands,
			Action: func(ctx context.Context, cmd *cli.Command) error {
				f.listCoin("", "")
				return nil
			},
		},
		{ // --- --- --- --- --- --- --- --- ADD
			Name:     "add",
			Aliases:  []string{"a"},
			Usage:    "Add a coin (d) or chain",
			Commands: addCommands,
			Action: func(ctx context.Context, cmd *cli.Command) error {
				f.addCoin("", 0, "")
				return nil
			},
		},
		{ // --- --- --- --- --- --- --- --- REMOVE
			Name:     "remove",
			Aliases:  []string{"r"},
			Usage:    "Remove a choin (d) or chain",
			Commands: removeCommands,
			Action: func(ctx context.Context, cmd *cli.Command) error {
				f.removeCoin("", "")
				return nil
			},
		},
		{ // --- --- --- --- --- --- --- --- EDIT
			Name:     "edit",
			Aliases:  []string{"e"},
			Usage:    "Edit choin (d) or chain",
			Commands: editCommands,
			Action: func(ctx context.Context, cmd *cli.Command) error {
				f.editCoin(0, 0, "x", "x")
				return nil
			},
		},
	}

	cmd := &cli.Command{
		Name:     "Coinme",
		Usage:    "CRUD coin exchange rates",
		Commands: commands,
		// Flags:    flags,
		Action: func(context.Context, *cli.Command) error {
			f.defaultUsage()
			return nil
		},
		Version: version,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
