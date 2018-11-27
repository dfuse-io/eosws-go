package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	eosws "github.com/dfuse-io/eosws-go"
)

func main() {
	flag.Parse()

	COUNT := 1

	for i := 0; i < COUNT; i++ {
		j := i
		go func() {
			// client, err := eosws.New("ws://localhost:8001/v1/stream", os.Getenv("EOSWS_API_KEY"), "https://origin.example.com")
			client, err := eosws.New("wss://mainnet.eos.dfuse.io/v1/stream", os.Getenv("EOSWS_API_KEY"), "https://origin.example.com")
			//client, err := eosws.New("wss://kylin.eos.dfuse.io/v1/stream", os.Getenv("EOSWS_API_KEY"), "https://origin.example.com")
			errorCheck("connecting to endpoint", err)

			ga := &eosws.GetActionTraces{}
			ga.ReqID = "foo GetActions"
			ga.StartBlock = -10
			ga.Listen = true
			ga.WithProgress = 5
			ga.Data.Accounts = "eosio.token"
			ga.Data.ActionNames = "transfer"
			ga.Data.WithInlineTraces = true

			fmt.Println("Sending `get_actions` message")
			err = client.Send(ga)
			errorCheck("sending get_actions", err)

			for {
				msg, err := client.Read()
				if err != nil {
					fmt.Println("DIED", j, err)
					return
				}

				switch m := msg.(type) {
				case *eosws.ActionTrace:
					cnt, _ := json.Marshal(m)
					fmt.Println(string(cnt))
				case *eosws.Progress:
					fmt.Println("Progress", j, m.Data.BlockNum)
				default:
					fmt.Println("Unsupported message", m)
				}
			}
		}()
	}

	time.Sleep(500 * time.Second)
}

func errorCheck(prefix string, err error) {
	if err != nil {
		fmt.Printf("ERROR: %s: %s\n", prefix, err)
		os.Exit(1)
	}
}
