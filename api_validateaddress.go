package multichain

// ValidateAddress returns information about address,
// or the address corresponding to the specified privkey private key or pubkey public key,
// including whether this node has the address’s private key in its wallet.
func (client *Client) ValidateAddress(address string) (Response, error) {

	msg := client.Command(
		"validateaddress",
		[]interface{}{
			address,
		},
	)

	return client.Post(msg)
}
