package woc

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	WOCHost = "https://api.whatsonchain.com"

	ContentTypeApplicationJSON = "application/json"

	PathHistoryForAddress = "/v1/bsv/%s/address/%s/history"

	PathBroadcast = "/v1/bsv/%s/broadcast"

	PathBulkTXs = "/v1/bsv/%s/txs"

	NetworkMain = "main"
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

func (w *Client) BulkTXs(ctx context.Context, hashes []string) ([]TXDetail, error) {
	if len(hashes) > 20 {
		// split the hashes into chunks of 20
		chunks := ChunkSlice(hashes, 20)

		//txs := []BSVJsonAddy{}
		txs := []TXDetail{}

		for _, set := range chunks {
			res, err := w.BulkTXs(ctx, set)
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

func (w *Client) Broadcast(tx []byte) error {
	s := fmt.Sprintf("%x", tx)

	req := BroadcastRequest{
		TXHex: s,
	}

	// e.g. POST https://api.whatsonchain.com/v1/bsv/main/broadcast
	path := fmt.Sprintf(PathBroadcast, w.Network)

	return w.post(path, req, nil)
}

func (w *Client) get(path string, out interface{}) error {
	res, err := w.Client.Get(w.buildURL(path))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP %v %s", res.StatusCode, http.StatusText(res.StatusCode))
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(out)
}

func (w *Client) post(path string, postBody, out interface{}) error {
	payload, err := NewJSONReader(postBody)
	if err != nil {
		return err
	}

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
