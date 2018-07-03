package multichain

import "fmt"

// Prepares an unspent transaction output (useful for building atomic exchange transactions) containing qty units of asset, where asset is an asset name, ref or issuance txid. Multiple items can be specified within the first parameter to include several assets within the output. The output will be locked against automatic selection for spending unless the optional lock parameter is set to false. Returns the txid and vout of the prepared output.
func (client *Client) PrepareLockUnspent(asset string, quantity float64, lock bool) (Response, error) {

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id":      ConstID,
		"method":  "preparelockunspent",
		"params": []interface{}{
			map[string]float64{
				asset: quantity,
			},
			fmt.Sprintf("lock=%v", lock),
		},
	}

	return client.Post(msg)
}
