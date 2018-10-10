package eosws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

func New(endpoint, token, origin string) (*Client, error) {
	endpoint = fmt.Sprintf("%s?token=%s", endpoint, token)
	reqHeaders := http.Header{"Origin": []string{origin}}
	conn, resp, err := websocket.DefaultDialer.Dial(endpoint, reqHeaders)
	if err != nil {
		return nil, fmt.Errorf("error, returned status=%d: %s: %s", resp.StatusCode, err, resp.Header.Get("X-Websocket-Handshake-Error"))
	}
	c := &Client{
		conn:     conn,
		incoming: make(chan interface{}, 1000),
	}

	go c.readLoop()

	return c, nil
}

type Client struct {
	conn      *websocket.Conn
	incoming  chan interface{}
	writeLock sync.Mutex
}

//

func (c *Client) Read() (interface{}, error) {
	select {
	case el := <-c.incoming:
		if el == nil {
			return nil, fmt.Errorf("connection closed")
		}
		return el, nil
	}
}

func (c *Client) readLoop() (*MsgIn, error) {
	defer func() {
		close(c.incoming)
		c.conn.Close()
	}()

	for {
		_ = c.conn.SetReadDeadline(time.Now().Add(120 * time.Second))
		msgType, cnt, err := c.conn.ReadMessage()
		if err != nil {
			return nil, err
		}
		if msgType != websocket.TextMessage {
			continue
		}

		var inspect CommonIn
		err = json.Unmarshal(cnt, &inspect)
		if err != nil {
			// LOG THERE WAS AN ERROR IN THE INCOMING JSON:
			continue
		}

		objType := IncomingMessageMap[inspect.Type]
		if objType == nil {
			// LOG: incoming message not supported, do we pass the raw JSON object?
			continue
		}

		obj := reflect.New(objType).Interface()
		err = json.Unmarshal(cnt, &obj)
		if err != nil {
			// LOG or push an error: "cannot unmarshal incoming message into our struct
			//emitError(inspect.ReqID, "unmarshal_message", err, wsmsg.M{"type": inspect.Type})
			continue
		}

		c.incoming <- obj
	}
}

// Send to the websocket, one of the messages registered through
// `RegisterOutboundMessage`.
func (c *Client) Send(msg OutgoingMessager) error {
	setType(msg)
	cnt, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshalling message %T (%#v): %s", msg, msg, err)
	}

	c.writeLock.Lock()
	defer c.writeLock.Unlock()
	if err := c.conn.WriteMessage(websocket.TextMessage, cnt); err != nil {
		return fmt.Errorf("writing message %T (%#v) to WS: %s", msg, msg, err)
	}

	return nil
}
