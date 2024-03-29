package woc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	WOCHost = "https://api.whatsonchain.com"

	ContentTypeApplicationJSON = "application/json"

	PathHistoryForAddress = "/v1/bsv/%s/address/%s/history"

	// PathBroadcast is the path to POST tx messages to.
	PathBroadcast = "/v1/bsv/%s/tx/raw"

	PathBulkTXs = "/v1/bsv/%s/txs"

	PathRawTX = "/v1/bsv/%s/txs/hex"

	// PathChainInfo is that path that returns information about the chain in general.
	//
	// GET https://api.whatsonchain.com/v1/bsv/<network>/chain/info
	PathChainInfo = "/v1/bsv/%s/chain/info"

	PathGetByHeight = "/v1/bsv/%s/block/height/%v"

	// NetworkMain indicates that the client should use the "main" network. This is the default.
	NetworkMain = "main"

	// NetworkTest indicates that the client should use the "test" network.
	NetworkTest = "test"

	// NetworkSTN indicates that the client should use the "stn" network.
	NetworkSTN = "stn"

	// PathTXByHash is the path to get a TX by hash.
	//
	// e.g. GET /v1/bsv/main/tx/hash/c1d32f28baa27a376ba977f6a8de6ce0a87041157cef0274b20bfda2b0d8df96
	PathTXByHash = "/v1/bsv/%s/tx/hash/%s"
)

type TXQuery struct {
	IDs []string `json:"txids"`
}

type Client struct {
	Host    string
	Network string
	Client  *http.Client
	TXCache map[string]string
}

func New() *Client {
	return &Client{
		Host:    WOCHost,
		Network: NetworkMain,
		Client:  newDefaultHTTPClient(),
		TXCache: map[string]string{},
	}
}

func (w *Client) RawTXs(hashes []string) ([]TXRaw, error) {
	path := fmt.Sprintf(PathRawTX, w.Network)

	txs := []TXRaw{}

	// {"txids" : ["294cd1ebd5689fdee03509f92c32184c0f52f037d4046af250229b97e0c8f1aa","91f68c2c598bc73812dd32d60ab67005eac498bef5f0c45b822b3c9468ba3258"]}"

	q := TXRawQuery{
		Hashes: hashes,
	}

	if err := w.post(path, q, &txs); err != nil {
		return nil, err
	}

	return txs, nil
}

func (w *Client) BulkTXs(hashes []string) ([]TXDetail, error) {
	if len(hashes) > 20 {
		// split the hashes into chunks of 20
		chunks := ChunkSlice(hashes, 20)

		//txs := []BSVJsonAddy{}
		txs := []TXDetail{}

		for _, set := range chunks {
			res, err := w.BulkTXs(set)
			if err != nil {
				return nil, err
			}

			txs = append(txs, res...)
		}

		return txs, nil
	}

	// TODO check the cache before querying....
	// ...

	q := TXQuery{
		IDs: hashes,
	}

	addresses := []TXDetail{}

	path := fmt.Sprintf(PathBulkTXs, w.Network)

	if err := w.post(path, q, &addresses); err != nil {
		return nil, err
	}

	return addresses, nil
}

func (w *Client) HistoryForAddress(address string) ([]HistoryTX, error) {
	path := fmt.Sprintf(PathHistoryForAddress, w.Network, address)

	txs := []HistoryTX{}

	if err := w.get(path, &txs); err != nil {
		return nil, err
	}

	return txs, nil
}

func (w *Client) LatestBlockInfo() (*BlockInfo, error) {
	path := fmt.Sprintf(PathChainInfo, w.Network)

	out := BlockInfo{}

	if err := w.get(path, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (w *Client) GetByHeight(height int64) (*BlockDetail, error) {
	// GET https://api.whatsonchain.com/v1/bsv/<network>/block/height/<height>
	path := fmt.Sprintf(PathGetByHeight, w.Network, height)

	out := BlockDetail{}

	if err := w.get(path, &out); err != nil {
		return nil, err
	}

	return &out, nil

}

func (w *Client) Broadcast(b []byte) error {
	s := fmt.Sprintf("%x", b)

	req := BroadcastRequest{
		TXHex: s,
	}

	// e.g. POST https://api.whatsonchain.com/v1/bsv/main/broadcast
	path := fmt.Sprintf(PathBroadcast, w.Network)

	return w.post(path, req, nil)
}

func (w *Client) GetTXByHash(hash string) (*TXDetail, error) {
	path := fmt.Sprintf(PathTXByHash, w.Network, hash)

	tx := TXDetail{}

	if err := w.get(path, &tx); err != nil {
		return nil, err
	}

	return &tx, nil
}

func (w *Client) get(path string, out interface{}) error {
	res, err := w.Client.Get(w.buildURL(path))
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode >= 300 || res.StatusCode < 200 {
		return fmt.Errorf("statuscode=%v : message=%s", res.StatusCode, body)
	}

	if err := json.Unmarshal(body, out); err != nil {
		return err
	}

	return nil
}

func (w *Client) post(path string, postBody, out interface{}) error {
	payload, err := NewJSONReader(postBody)
	if err != nil {
		return err
	}

	// fmt.Printf("payload = %s\n", payload)
	res, err := w.Client.Post(w.buildURL(path), ContentTypeApplicationJSON, payload)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// fmt.Printf("body = %s\n", body)

	if res.StatusCode >= 300 || res.StatusCode < 200 {
		return fmt.Errorf("statuscode=%v : message=%s", res.StatusCode, body)
	}

	if out != nil {
		if err := json.Unmarshal(body, out); err != nil {
			return err
		}
	}

	return nil
}

func (w *Client) buildURL(path string) string {
	return fmt.Sprintf("%s/%s", w.Host, path)
}
