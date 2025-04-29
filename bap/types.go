package bap

import "github.com/bsv-blockchain/go-sdk/chainhash"

var IdentityKey = "identity"
var ProfileKey = "profile"
var AddressIdentityKey = "add:id"
var AttestationKey = "attestation"

type Address struct {
	Address   string         `json:"address"`
	Txid      chainhash.Hash `json:"txId"`
	Block     uint32         `json:"block"`
	Timestamp uint32         `json:"-"`
}

type Identity struct {
	IDKey          string    `json:"idKey"`
	FirstSeen      uint32    `json:"firstSeen"`
	RootAddress    string    `json:"rootAddress"`
	CurrentAddress string    `json:"currentAddress"`
	Addresses      []Address `json:"addresses"`
	Identity       any       `json:"identity,omitempty"`
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
	Id      string    `json:"hash"`
	Signers []*Signer `json:"signers"`
	// Attribute string    `json:"attribute,omitempty"`
	// Value     string    `json:"value,omitempty"`
	// Nonce     string    `json:"nonce,omitempty"`
	// URN       string    `json:"urn,omitempty"`
}
