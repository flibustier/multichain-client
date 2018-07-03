package multichain

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	//
	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
	//
	"github.com/dghubble/sling"
)

const (
	// ConstID is a constant used at some places
	ConstID = "multichain-client"
)

// Response type is returned by the JSON RPC API
type Response map[string]interface{}

// Result return the payload of the response
func (r Response) Result() interface{} {
	return r["result"]
}

// Client is the main structure
type Client struct {
	httpClient  *http.Client
	chain       string
	host        string
	port        int
	credentials string
	debug       bool
}

// NewClient initialize a new Client
func NewClient(chain, username, password string) *Client {

	credentials := username + ":" + password

	return &Client{
		httpClient:  &http.Client{},
		chain:       chain,
		credentials: base64.StdEncoding.EncodeToString([]byte(credentials)),
	}
}

// ViaLocal uses localhost connection with given port to connect to the API
func (client *Client) ViaLocal(port int) *Client {
	return client.ViaNode("localhost", port)
}

// ViaNode initialize the connection to the API via IP and port
func (client *Client) ViaNode(ipv4 string, port int) *Client {
	c := *client
	c.host = fmt.Sprintf(
		"http://%s:%v",
		ipv4,
		port,
	)
	return &c
}

// IsDebugMode returns if the debug mode is on
func (client *Client) IsDebugMode() bool {
	return client.debug
}

// DebugMode returns a Client with the debug mode on
func (client *Client) DebugMode() *Client {
	client.debug = true
	return client
}

func (client *Client) Urlfetch(ctx context.Context, seconds ...int) {

	if len(seconds) > 0 {
		ctx, _ = context.WithDeadline(
			ctx,
			time.Now().Add(time.Duration(1000000000*seconds[0])*time.Second),
		)
	}

	client.httpClient = urlfetch.Client(ctx)
}

func (client *Client) msg(params []interface{}) map[string]interface{} {
	return map[string]interface{}{
		"jsonrpc": "1.0",
		"id":      ConstID,
		"params":  params,
	}
}

func (client *Client) Command(method string, params []interface{}) map[string]interface{} {

	msg := client.msg(params)
	msg["method"] = fmt.Sprintf("%s", method)

	if client.debug {
		fmt.Println(msg)
	}

	return msg
}

func (client *Client) Post(msg interface{}) (Response, error) {

	if client.debug {
		fmt.Println("DEBUG MODE ON...")
		fmt.Println(client)
		b, _ := json.Marshal(msg)
		fmt.Println(string(b))
	}

	request, err := sling.New().Post(client.host).BodyJSON(msg).Request()
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "Basic "+client.credentials)

	resp, err := client.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if client.debug {
		fmt.Println(string(b))
	}

	obj := make(Response)

	err = json.Unmarshal(b, &obj)
	if err != nil {
		return nil, err
	}

	if obj["error"] != nil {
		e := obj["error"].(map[string]interface{})
		var s string
		m, ok := msg.(map[string]interface{})
		if ok {
			s = fmt.Sprintf("multichaind - '%s': %s", m["method"], e["message"].(string))
		} else {
			s = fmt.Sprintf("multichaind - %s", e["message"].(string))
		}
		return nil, errors.New(s)
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("INVALID RESPONSE STATUS CODE: " + strconv.Itoa(resp.StatusCode))
	}

	return obj, nil
}
