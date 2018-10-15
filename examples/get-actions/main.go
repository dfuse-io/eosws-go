package main

import (
	"fmt"
	"os"

	eosws "github.com/eoscanada/eosws-go"
)

func main() {
	apiKey := os.Getenv("EOSWS_API_KEY")
	if apiKey == "" {
		errorCheck("missing api key", fmt.Errorf("An API key must be specified through the 'EOSWS_API_KEY' environment variable"))
	}

	client, err := eosws.New("wss://mainnet.eos.dfuse.io/v1/stream", apiKey, "https://example.com")
	errorCheck("connecting to endpoint", err)

	ga := &eosws.GetActions{}
	ga.ReqID = "get-accounts-jsons"
	ga.StartBlock = -5000
	ga.Listen = true
	ga.Data.Account = "eosio.token"
	ga.Data.ActionName = "transfer"

	fmt.Println("Sending `get_actions` message")
	err = client.Send(ga)
	errorCheck("sending get_actions", err)

	for {
		msg, err := client.Read()
		errorCheck("reading message", err)

		switch m := msg.(type) {
		case *eosws.ActionTrace:
			fmt.Println(m.Data.Trace)
		}
	}
}

func errorCheck(prefix string, err error) {
	if err != nil {
		fmt.Printf("ERROR: %s: %s\n", prefix, err)
		os.Exit(1)
	}
}
