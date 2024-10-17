package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	// get arguments from command line
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("no command given")
		os.Exit(1)
	}

	ctx := context.Background()
	client := getClient()
	LEDGER_KEY := os.Getenv("LEDGER_KEY")
	sheetService := NewSheetsService(ctx, client)
	sheetService.SetSpreadSheetId(LEDGER_KEY)

	switch args[0] {
	case "balance-all", "ba":
		res := big.NewInt(sheetService.GetTotalBalance())
		fmt.Println("total balance: Rp.", res)
	case "balance", "b":
		if len(args) < 2 {
			fmt.Println("wallet name is required")
			os.Exit(1)
		}
		res := big.NewInt(sheetService.GetCurrentBalance(args[1]))
		fmt.Println("current balance: Rp.", res)
	case "mutate", "m":
		if len(args) < 5 {
			fmt.Println("wallet name, kredit, debit, and description are required")
			os.Exit(1)
		}
		kredit, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			fmt.Println("kredit must be a number")
			os.Exit(1)
		}
		debit, err := strconv.ParseFloat(args[3], 64)
		if err != nil {
			fmt.Println("debit must be a number")
			os.Exit(1)
		}
		sheetService.PushMutation(args[1], kredit, debit, args[4])
		fmt.Println("done")
	case "transfer", "t":
		if len(args) < 5 {
			fmt.Println("wallet from, wallet to, balance, and admin are required")
		}
		balance, err := strconv.ParseFloat(args[3], 64)
		if err != nil {
			fmt.Println("balance must be a number")
			os.Exit(1)
		}
		admin, err := strconv.ParseFloat(args[4], 64)
		if err != nil {
			fmt.Println("admin must be a number")
			os.Exit(1)
		}
		sheetService.TransferBalance(args[1], args[2], balance, admin)
		fmt.Println("done")
	case "help", "h":
		fmt.Println("balance-all, ba : get total balance")
		fmt.Println("balance, b      : get current balance")
		fmt.Println("mutate, m       : push mutation <wallet> <kredit> <debit> <description>")
		fmt.Println("transfer, t     : trasfer <from> <to> <balance> <admin>")
		fmt.Println("help, h: show this help")
		os.Exit(0)
	default:
		fmt.Println("command not found")
		os.Exit(1)
	}
}
