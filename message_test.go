package eosws

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetActions(t *testing.T) {

	key := os.Getenv("EOSWS_API_KEY")
	require.NotEqual(t, "", key)

	client, err := New("wss://staging-kylin.eos.dfuse.io/v1/stream", key, "https://origin.example.com")
	require.NoError(t, err)

	ga := &GetActionTraces{}
	ga.ReqID = "1"
	ga.StartBlock = -10
	ga.Listen = true
	ga.Data.Account = "eosio.token"
	ga.Data.ActionName = "transfer"
	ga.Data.WithInlineTraces = true

	client.Send(ga)

	for {
		msg, err := client.Read()
		if err != nil {
			fmt.Println("DIED", err)
			t.Fail()
			return
		}

		switch m := msg.(type) {
		case *ActionTrace:
			cnt, _ := json.Marshal(m)
			fmt.Println("cnt: ", string(cnt))
			return
		case *Listening:
			break
		default:
			cnt, _ := json.Marshal(m)
			fmt.Println("cnt: ", string(cnt))
			t.Errorf("Invalid message: %T", m)
		}

	}
}

func Test_GetTableRowsFetch(t *testing.T) {

	key := os.Getenv("EOSWS_API_KEY")
	require.NotEqual(t, "", key)

	client, err := New("wss://staging-kylin.eos.dfuse.io/v1/stream", key, "https://origin.example.com")
	require.NoError(t, err)

	ga := &GetTableRows{}
	ga.ReqID = "1"
	ga.StartBlock = -6000
	ga.Listen = false
	ga.Fetch = true
	ga.WithProgress = 5
	ga.Data.JSON = true
	ga.Data.Code = "eosio.token"
	ga.Data.Scope = "eosio"
	ga.Data.Table = "accounts"
	client.Send(ga)

	for {
		msg, err := client.Read()
		if err != nil {
			fmt.Println("DIED", err)
			t.Fail()
			return
		}

		switch m := msg.(type) {
		case *TableSnapshot:
			cnt, _ := json.Marshal(m)
			fmt.Println(string(cnt))
			return
		default:
			t.Errorf("Invalid message: %T", m)
		}
	}
}

func Test_GetTableRowsListen(t *testing.T) {

	key := os.Getenv("EOSWS_API_KEY")
	require.NotEqual(t, "", key)

	client, err := New("wss://staging-kylin.eos.dfuse.io/v1/stream", key, "https://origin.example.com")
	require.NoError(t, err)

	ga := &GetTableRows{}
	ga.ReqID = "1"
	ga.StartBlock = -3600
	ga.Listen = true
	ga.Fetch = false
	ga.WithProgress = 5
	ga.Data.JSON = true
	ga.Data.Code = "eosio"
	ga.Data.Scope = "eosio"
	ga.Data.Table = "global"
	client.Send(ga)

	for {

		msg, err := client.Read()
		if err != nil {
			fmt.Println("DIED", err)
			t.Fail()
			return
		}

		switch m := msg.(type) {
		case *TableDelta:
			cnt, _ := json.Marshal(m)
			fmt.Println(string(cnt))
			return
		case *Listening:
			break
		default:
			t.Errorf("Invalid message: %T", m)
		}
	}
}
