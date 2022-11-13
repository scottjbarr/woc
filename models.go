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

type TXRawQuery struct {
	Hashes []string `json:"txids"`
}

type TXRaw struct {
	Blockhash     string `json:"blockhash"`
	Blockheight   int64  `json:"blockheight"`
	Blocktime     int64  `json:"blocktime"`
	Confirmations int64  `json:"confirmations"`
	Txid          string `json:"txid"`
	Hex           string `json:"hex"`
}

type Vin struct {
	Coinbase  string    `json:"coinbase"`
	ScriptSig ScriptSig `json:"scriptSig"`
	Sequence  int64     `json:"sequence"`
	Txid      string    `json:"txid"`
	Vout      int64     `json:"vout"`
}

type ScriptSig struct {
	Asm string `json:"asm"`
	Hex string `json:"hex"`
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

type BlockInfo struct {
	BestBlockHash        string  `json:"bestblockhash"`
	Blocks               int64   `json:"blocks"`
	Chain                string  `json:"chain"`
	Chainwork            string  `json:"chainwork"`
	Difficulty           float64 `json:"difficulty"`
	Headers              int64   `json:"headers"`
	MedianTime           int64   `json:"mediantime"`
	Pruned               bool    `json:"pruned"`
	VerificationProgress float64 `json:"verificationprogress"`
}

type BlockDetail struct {
	Hash              string      `json:"hash"`
	Confirmations     int         `json:"confirmations"`
	Size              int         `json:"size"`
	Height            int         `json:"height"`
	Version           int         `json:"version"`
	VersionHex        string      `json:"versionHex"`
	Merkleroot        string      `json:"merkleroot"`
	Txcount           int         `json:"txcount"`
	Tx                []string    `json:"tx"`
	Time              int         `json:"time"`
	Mediantime        int         `json:"mediantime"`
	Nonce             int64       `json:"nonce"`
	Bits              string      `json:"bits"`
	Difficulty        float64     `json:"difficulty"`
	Chainwork         string      `json:"chainwork"`
	Previousblockhash string      `json:"previousblockhash"`
	Nextblockhash     string      `json:"nextblockhash"`
	CoinbaseTx        CoinbaseTX  `json:"coinbaseTx"`
	TotalFees         float64     `json:"totalFees"`
	Miner             string      `json:"miner"`
	Pages             interface{} `json:"pages"`
}

type CoinbaseTX struct {
	Hex           string `json:"hex"`
	Txid          string `json:"txid"`
	Hash          string `json:"hash"`
	Size          int    `json:"size"`
	Version       int    `json:"version"`
	Locktime      int    `json:"locktime"`
	Vin           []Vin  `json:"vin"`
	Vout          []Vout `json:"vout"`
	Blockhash     string `json:"blockhash"`
	Confirmations int    `json:"confirmations"`
	Time          int    `json:"time"`
	Blocktime     int    `json:"blocktime"`
}
