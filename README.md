eosws Go bindings (from the dfuse API)
--------------------------------------

Websocket consumer for the https://dfuse.io API on EOS networks.

## Connecting

```go
    jwt, exp, err := eosws.Auth("server_1234567....")
    if err != nil {
        log.Fatalf("cannot get auth token: %s", err.Error())
    }
    time.AfterFunc(time.Until(exp), log.Println("JWT is now expired, renew it before reconnecting client")) // make sure that you handle updating your JWT

	client, err := eosws.New("wss://mainnet.eos.dfuse.io/v1/stream", jwt, "https://origin.example.com")
    if err != nil {
        log.Fatalf("cannot connect to dfuse endpoint: %s", err.Error())
    }
```

## Sending requests

```go
	ga := &eosws.GetActionTraces{
		ga := &eosws.GetActionTraces{
			ReqID:      "myreq1",
			StartBlock: -5,
			Listen:     true,
		}
	}
	ga.Data.Accounts = "eosio"
	ga.Data.ActionNames = "onblock"
	err = client.Send(ga)
	if err != nil {
		log.Fatalf("error sending request")
    }
```

## Reading responses

```go
	for {
		msg, err := client.Read()
		errorCheck("reading message", err)

		switch m := msg.(type) {
		case *eosws.ActionTrace:
			fmt.Println(m.Data.Trace)
		default:
			fmt.Println("Unsupported message", m)
			break
		}
	}
```

## Examples

See `examples` folder
