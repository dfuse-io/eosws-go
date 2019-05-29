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

MAINLOOP:
	for {
		log.Println("Beginning of loop")
		var jwt string
		var err error
		for {
			jwt, _, err = eosws.Auth(api_key)
			if err != nil {
				log.Println("cannot get JWT token: %s", err.Error())
				time.Sleep(2 * time.Second)
				continue
			}
			break
		}
		log.Println("Got new JWT:", jwt)

		log.Println("connecting to endpoint " + dfuse_endpoint)
		client, err := eosws.New(dfuse_endpoint, jwt, origin)
		errorCheck("connecting to endpoint"+dfuse_endpoint, err)

		ga := &eosws.GetActionTraces{}
		ga.ReqID = "GetActions"
		ga.StartBlock = 0
		ga.Listen = true
		ga.WithProgress = 10
		ga.IrreversibleOnly = false
		ga.Data.Accounts = "delphioracle"
		ga.Data.ActionNames = "write"
		ga.Data.WithInlineTraces = false

		fmt.Printf("Sending `get_actions` message for accounts: %s and action names: %s\n", ga.Data.Accounts, ga.Data.ActionNames)
		err = client.Send(ga)
		errorCheck("sending get_actions", err)

		for {
			var msg interface{}
			msgCh := make(chan interface{})
			errCh := make(chan error)

			go func() {
				msg, err := client.Read()
				if err != nil {
					errCh <- err
				}
				msgCh <- msg
			}()

			select {
			case msg = <-msgCh:
				break
			case err := <-errCh:
				log.Println("Got error", err)
				time.Sleep(2 * time.Second)
				continue MAINLOOP
			case <-time.After(20 * time.Second):
				log.Println("Got no data for 20 seconds")
				continue MAINLOOP
			}

			switch m := msg.(type) {
			case *eosws.ActionTrace:
				log.Println("Block Num:", m.Data.BlockNum, m.Data.TransactionID)
			case *eosws.Progress:
				log.Println("Progress", m.Data.BlockNum)
			case *eosws.Listening:
				log.Println("listening...")
			default:
				log.Println("Unsupported message", m)
			}
		}
	}

}

func errorCheck(prefix string, err error) {
	if err != nil {
		log.Fatalf("ERROR: %s: %s\n", prefix, err)
	}
}
