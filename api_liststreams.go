package multichain

import (
	"fmt"
)

// ListStreams returns informations about streams created on the blockchain.
// Pass a stream name, ref or creation txid in streams to retrieve information about one stream only,
// an array thereof for multiple streams, or * for all streams.
// Use count and start to retrieve part of the list only, with negative start values (like the default) indicating the most recently created streams.
// Extra fields are shown for streams to which this node has subscribed.
func (client *Client) ListStreams(streams string, start, count int, verbose bool) (Response, error) {

	if len(streams) == 0 {
		streams = "*"
	}

	params := []interface{}{
		fmt.Sprintf("%s", streams),
	}

	/*

		params := []interface{}{
			fmt.Sprintf("streams=%s", streams),
		}

		if start > 0 && count > 0 {
			params = append(params, fmt.Sprintf("start=%d", start))
			params = append(params, fmt.Sprintf("count=%d", count))
		}
		if verbose {
			params = append(params, fmt.Sprintf("verbose=%v", verbose))
		}
	*/

	msg := map[string]interface{}{
		"jsonrpc": "1.0",
		"id":      ConstID,
		"method":  "liststreams",
		"params":  params,
	}

	return client.Post(msg)
}
