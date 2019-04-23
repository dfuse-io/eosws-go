package main

import (
	"encoding/json"
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

		ga := &eosws.GetTableRows{}
		ga.ReqID = "foo GetTableRows"
		ga.StartBlock = -3600
		ga.Listen = true
		ga.Fetch = true
		ga.WithProgress = 5
		ga.Data.JSON = true
		ga.Data.Code = "eosio.token"
		ga.Data.Scope = "eosio"
		ga.Data.Table = "accounts"

		fmt.Println("Sending `get_table_rows` message")
		err = client.Send(ga)
		errorCheck("sending get_table_rows", err)

		for {
			msg, err := client.Read()
			if err != nil {
				fmt.Println("DIED", err)
				return
			}

			switch m := msg.(type) {
			case *eosws.Progress:
				fmt.Println("Progress", m.Data.BlockNum)
			case *eosws.TableDelta:
				fmt.Printf("%d: %+v\n", m.Data.BlockNum, m.Data.DBOp)
			case *eosws.TableSnapshot:
				cnt, _ := json.Marshal(m)
				fmt.Println("Rows: ", string(cnt))
			case *eosws.Listening:
				fmt.Println("listening...")
			default:
				fmt.Println("Unsupported message", m)
			}
		}
	}()

	timeout := 10 * time.Second
	time.Sleep(timeout)
	log.Println("done after", timeout)
}

func errorCheck(prefix string, err error) {
	if err != nil {
		log.Fatalf("ERROR: %s: %s\n", prefix, err)
	}
}
