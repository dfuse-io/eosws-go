package main

import (
	"fmt"
	"log"
	"os"
	"time"

	eosws "github.com/dfuse-io/eosws-go"
)

var dfuse_endpoint = "wss://mainnet.eos.dfuse.io/v1/stream"
var origin = "https://origin.example.com"

func main() {

	api_key := os.Getenv("EOSWS_API_KEY")
	if api_key == "" {
		log.Fatalf("please set your API key to environment variable EOSWS_API_KEY")
	}

	jwt, _, err := eosws.Auth(api_key)
	if err != nil {
		log.Fatalf("cannot get JWT token: %s", err.Error())
	}

	client, err := eosws.New(dfuse_endpoint, jwt, origin)
	errorCheck("connecting to endpoint"+dfuse_endpoint, err)

	go func() {

		ga := &eosws.GetActionTraces{}
		ga.ReqID = "foo GetActions"
		ga.StartBlock = -350
		ga.Listen = true
		ga.WithProgress = 5
		ga.IrreversibleOnly = true
		ga.Data.Accounts = "eosio.token"
		ga.Data.ActionNames = "transfer"
		ga.Data.WithInlineTraces = true

		fmt.Printf("Sending `get_actions` message for accounts: %s and action names: %s", ga.Data.Accounts, ga.Data.ActionNames)
		err = client.Send(ga)
		errorCheck("sending get_actions", err)

		for {
			msg, err := client.Read()
			if err != nil {
				fmt.Println("DIED", err)
				return
			}

			switch m := msg.(type) {
			case *eosws.ActionTrace:
				fmt.Println("Block Num:", m.Data.BlockNum, m.Data.TransactionID)
			case *eosws.Progress:
				fmt.Println("Progress", m.Data.BlockNum)
			case *eosws.Listening:
				fmt.Println("listening...")
			default:
				fmt.Println("Unsupported message", m)
			}
		}
	}()

	time.Sleep(8 * time.Second)
}

func errorCheck(prefix string, err error) {
	if err != nil {
		log.Fatalf("ERROR: %s: %s\n", prefix, err)
	}
}
