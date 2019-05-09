package eosws

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

func Auth(apiKey string) (token string, expiration time.Time, err error) {
	return AuthWithURL(apiKey, "https://auth.dfuse.io/v1/auth/issue")
}

func AuthWithURL(apiKey string, authServiceURL string) (token string, expiration time.Time, err error) {
	resp, err := http.Post(authServiceURL, "application/json", bytes.NewBuffer([]byte(fmt.Sprintf(`{"api_key":"%s"}`, apiKey))))
	if err != nil {
		return token, expiration, err
	}
	if resp.StatusCode != 200 {
		return token, expiration, fmt.Errorf("wrong status code from Auth: %d", resp.StatusCode)
	}
	var result apiToJWTResp
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return
	}
	return result.Token, time.Unix(result.ExpiresAt, 0), nil
}

func New(endpoint, token, origin string) (*Client, error) {
	endpoint = fmt.Sprintf("%s?token=%s", endpoint, token)
	reqHeaders := http.Header{"Origin": []string{origin}}
	conn, resp, err := websocket.DefaultDialer.Dial(endpoint, reqHeaders)
	if err != nil {
		if resp != nil {
			return nil, fmt.Errorf("error, returned status=%d: %s: %s", resp.StatusCode, err, resp.Header.Get("X-Websocket-Handshake-Error"))
		}
		return nil, fmt.Errorf("error dialing to endpoint: %s", err)
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
	readError error
	incoming  chan interface{}
	writeLock sync.Mutex
}

//

func (c *Client) Read() (interface{}, error) {
	select {
	case el := <-c.incoming:
		if el == nil {
			return nil, c.readError
		}
		return el, nil
	}
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) readLoop() {
	var err error
	var msgType int
	var cnt []byte

	defer func() {
		c.readError = err
		close(c.incoming)
		c.conn.Close()
	}()

	for {
		_ = c.conn.SetReadDeadline(time.Now().Add(120 * time.Second))
		msgType, cnt, err = c.conn.ReadMessage()
		if err != nil {
			// the `defer` will return the `err` here..
			// LOG error, close write, shutdown client, store the error, whatever
			return
		}
		if msgType != websocket.TextMessage {
			fmt.Println("eosws client: invalid incoming message type", msgType)
			// Server should not send messages other than json-encoded text messages
			continue
		}

		var inspect CommonIn
		err = json.Unmarshal(cnt, &inspect)
		if err != nil {
			fmt.Println("eosws client: error unmarshaling incoming message:", err)
			// LOG THERE WAS AN ERROR IN THE INCOMING JSON:
			continue
		}

		if inspect.Type == "ping" {
			pong := bytes.Replace(cnt, []byte(`"ping"`), []byte(`"pong"`), 1)
			//fmt.Println("eosws client: sending pong", string(pong))
			_ = c.conn.WriteMessage(websocket.TextMessage, pong)
			continue
		}

		objType := IncomingMessageMap[inspect.Type]
		if objType == nil {
			fmt.Printf("eosws client: received unsupported incoming message type %q\n", inspect.Type)
			// LOG: incoming message not supported, do we pass the raw JSON object?
			continue
		}

		obj := reflect.New(objType).Interface()
		err = json.Unmarshal(cnt, &obj)
		if err != nil {

			fmt.Printf("Error unmarshalling :%q :%s\n", inspect.Type, err)
			fmt.Println("Data: ", string(cnt))
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

type apiToJWTResp struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}
