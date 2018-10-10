eosws Go bindings (from the dfuse API)
--------------------------------------

Websocket consumer for the https://dfuse.io API on EOS networks.

## Sample usage

```go
	client, err := eosws.New("wss://eosws.mainnet.eoscanada.com/v1/stream", "eyJ...nadacom", "https://origin.example.com")
	errorCheck("connecting to endpoint", err)

	ga := &eosws.GetActions{}
	ga.ReqID = "get-accounts-jsons"
	ga.StartBlock = -5000
	ga.Listen = true
	ga.Data.Account = "accountsjson"
	ga.Data.ActionName = "set"

	fmt.Println("Sending `get_actions` message")
	err = client.Send(ga)
	errorCheck("sending get_actions", err)

	for {
		msg, err := client.Read()
		errorCheck("reading message", err)

		switch m := msg.(type) {
		case *eosws.ActionTrace:
			fmt.Println(m.Data.Trace)
		default:
			fmt.Println("Unsupported message", m)
		}
	}
```

where:

```go
func errorCheck(prefix string, err error) {
	if err != nil {
		fmt.Printf("ERROR: %s: %s\n", prefix, err)
		os.Exit(1)
	}
}
```
