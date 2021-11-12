package woc

import "fmt"

type TXDetail struct {
	Blockhash     string `json:"blockhash"`
	Blockheight   int64  `json:"blockheight"`
	Blocktime     int64  `json:"blocktime"`
	Confirmations int64  `json:"confirmations"`
	Hash          string `json:"hash"`
	Locktime      int64  `json:"locktime"`
	Size          int64  `json:"size"`
	Time          int64  `json:"time"`
	Txid          string `json:"txid"`
	Version       int64  `json:"version"`
	Vin           []Vin  `json:"vin"`
	Vout          []Vout `json:"vout"`
}

type Vin struct {
	Coinbase  string `json:"coinbase"`
	ScriptSig struct {
		Asm string `json:"asm"`
		Hex string `json:"hex"`
	} `json:"scriptSig"`
	Sequence int64  `json:"sequence"`
	Txid     string `json:"txid"`
	Vout     int64  `json:"vout"`
}

type Float64 float64

func NewFloat64(f float64) Float64 {
	return Float64(f)
}

func (f Float64) String() string {
	return fmt.Sprintf("%0.8f", f)
}

type Vout struct {
	N            int64        `json:"n"`
	ScriptPubKey ScriptPubKey `json:"scriptPubKey"`
	Value        Float64      `json:"value"`
}

type ScriptPubKey struct {
	Addresses   []string    `json:"addresses"`
	Asm         string      `json:"asm"`
	Hex         string      `json:"hex"`
	IsTruncated bool        `json:"isTruncated"`
	OpReturn    interface{} `json:"opReturn"`
	ReqSigs     int64       `json:"reqSigs"`
	Type        string      `json:"type"`
}

type HistoryTX struct {
	TxHash string `json:"tx_hash"`
	Height int64  `json:"height"`
}

type BroadcastRequest struct {
	TXHex string `json:"txhex"`
}
