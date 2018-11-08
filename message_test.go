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

	msg, err := client.Read()
	if err != nil {
		fmt.Println("DIED", err)
		t.Fail()
		return
	}

	switch m := msg.(type) {
	case *ActionTrace:
		cnt, _ := json.Marshal(m)
		fmt.Println(string(cnt))
		//todo need assert here...
	case *Progress:
		fmt.Println("Progress:", m.Data.BlockNum)
	default:
		fmt.Println("Unsupported message", m)
	}

}
