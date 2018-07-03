package multichain

import (
	"flag"
	"fmt"
)

var client *Client

func init() {

	chain := flag.String("chain", "", "is the name of the chain")
	host := flag.String("host", "localhost", "is a string for the hostname")
	port := flag.Int("port", 80, "is a string for the host port")
	username := flag.String("username", "multichainrpc", "is a string for the username")
	password := flag.String("password", "12345678", "is a string for the password")

	flag.Parse()

	client = NewClient(
		*chain,
		*username,
		*password,
	).ViaNode(
		*host,
		*port,
	).DebugMode()

	fmt.Println(client.IsDebugMode())
}
