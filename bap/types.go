package bap

import "github.com/bsv-blockchain/go-sdk/chainhash"

type Address struct {
	Address   string         `json:"address" bson:"address"`
	Txid      chainhash.Hash `json:"txId" bson:"txId"`
	Block     uint32         `json:"block" bson:"block"`
	Timestamp uint32         `json:"-" bson:"timestamp"`
}

type Identity struct {
	IDKey          string      `json:"idKey" bson:"_id"`
	FirstSeen      uint32      `json:"firstSeen" bson:"firstSeen"`
	RootAddress    string      `json:"rootAddress" bson:"rootAddress"`
	CurrentAddress string      `json:"currentAddress" bson:"currentAddress"`
	Addresses      []Address   `json:"addresses" bson:"addresses"`
	Identity       interface{} `json:"identity" bson:"-"`
}

type Signer struct {
	IDKey     string         `json:"idKey"`
	Address   string         `json:"signingAddress"`
	Sequence  uint64         `json:"sequence"`
	Block     uint32         `json:"block"`
	Txid      chainhash.Hash `json:"txId"`
	Timestamp uint32         `json:"timestamp"`
	Revoked   bool           `json:"revoked"`
}

type Attestation struct {
	Id        string    `json:"hash"`
	Attribute string    `json:"attribute,omitempty"`
	Value     string    `json:"value,omitempty"`
	Nonce     string    `json:"nonce,omitempty"`
	URN       string    `json:"urn,omitempty"`
	Signers   []*Signer `json:"signers"`
}
