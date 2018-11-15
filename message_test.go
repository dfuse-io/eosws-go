package eosws

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetActions(t *testing.T) {

	client := newClient(t)

	ga := &GetActionTraces{}
	ga.ReqID = "1"
	ga.StartBlock = -10
	ga.Listen = true
	ga.Data.Account = "eosio.token"
	ga.Data.ActionName = "transfer"
	ga.Data.WithInlineTraces = true

	client.Send(ga)
	defer client.conn.Close()
	expectMessage(t, client, &Listening{}, nil)
	expectMessage(t, client, &ActionTrace{}, nil)
}

func Test_GetTableRowsFetch(t *testing.T) {

	client := newClient(t)

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
	defer client.conn.Close()

	expectMessage(t, client, &TableSnapshot{}, nil)

}

func Test_GetTableRowsListen(t *testing.T) {

	client := newClient(t)

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
	defer client.conn.Close()

	expectMessage(t, client, &Listening{}, nil)
	expectMessage(t, client, &TableDelta{}, func(t *testing.T, msg interface{}) {
		delta := msg.(*TableDelta)
		assert.Equal(t, "1", delta.ReqID)
		assert.NotEqual(t, "", delta.Data.Step)
	})

}

func newClient(t *testing.T) *Client {
	t.Helper()

	key := os.Getenv("EOSWS_API_KEY")
	require.NotEqual(t, "", key)

	client, err := New("wss://staging-kylin.eos.dfuse.io/v1/stream", key, "https://origin.example.com")
	require.NoError(t, err)

	return client
}

func expectMessage(t *testing.T, client *Client, messageType interface{}, validation func(t *testing.T, msg interface{})) {
	msg, err := client.Read()
	if err != nil {
		fmt.Println("DIED", err)
		t.Fail()
		return
	}

	msgType := reflect.TypeOf(msg).String()
	require.Equal(t, reflect.TypeOf(messageType).String(), msgType)

	if validation != nil {
		validation(t, msg)
	}

}
