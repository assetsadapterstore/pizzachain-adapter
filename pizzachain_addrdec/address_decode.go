package pizzachain_addrdec

import (
	"fmt"
	"github.com/blocktree/openwallet/v2/openwallet"
	"strings"

	"github.com/blocktree/go-owcdrivers/addressEncoder"
)

var (
	PIZPublicKeyPrefix       = "PUB_"
	PIZPublicKeyK1Prefix     = "PUB_K1_"
	PIZPublicKeyR1Prefix     = "PUB_R1_"
	PIZPublicKeyPrefixCompat = "PIZ"

	//PIZ stuff
	PIZ_mainnetPublic = addressEncoder.AddressType{"piz", addressEncoder.BTCAlphabet, "ripemd160", "", 33, []byte(PIZPublicKeyPrefixCompat), nil}
	// PIZ_mainnetPrivateWIF           = AddressType{"base58", BTCAlphabet, "doubleSHA256", "", 32, []byte{0x80}, nil}
	// PIZ_mainnetPrivateWIFCompressed = AddressType{"base58", BTCAlphabet, "doubleSHA256", "", 32, []byte{0x80}, []byte{0x01}}

	Default = AddressDecoderV2{}
)

//AddressDecoderV2
type AddressDecoderV2 struct {
	openwallet.AddressDecoderV2Base
	IsTestNet bool
}

// AddressDecode decode address
func (dec *AddressDecoderV2) AddressDecode(pubKey string, opts ...interface{}) ([]byte, error) {

	var pubKeyMaterial string
	if strings.HasPrefix(pubKey, PIZPublicKeyR1Prefix) {
		pubKeyMaterial = pubKey[len(PIZPublicKeyR1Prefix):] // strip "PUB_R1_"
	} else if strings.HasPrefix(pubKey, PIZPublicKeyK1Prefix) {
		pubKeyMaterial = pubKey[len(PIZPublicKeyK1Prefix):] // strip "PUB_K1_"
	} else if strings.HasPrefix(pubKey, PIZPublicKeyPrefixCompat) { // "PIZ"
		pubKeyMaterial = pubKey[len(PIZPublicKeyPrefixCompat):] // strip "PIZ"
	} else {
		return nil, fmt.Errorf("public key should start with [%q | %q] (or the old %q)", PIZPublicKeyK1Prefix, PIZPublicKeyR1Prefix, PIZPublicKeyPrefixCompat)
	}

	ret, err := addressEncoder.Base58Decode(pubKeyMaterial, addressEncoder.NewBase58Alphabet(PIZ_mainnetPublic.Alphabet))
	if err != nil {
		return nil, addressEncoder.ErrorInvalidAddress
	}
	if addressEncoder.VerifyChecksum(ret, PIZ_mainnetPublic.ChecksumType) == false {
		return nil, addressEncoder.ErrorInvalidAddress
	}

	return ret[:len(ret)-4], nil
}

// AddressEncode encode address
func (dec *AddressDecoderV2) AddressEncode(hash []byte, opts ...interface{}) (string, error) {
	data := addressEncoder.CatData(hash, addressEncoder.CalcChecksum(hash, PIZ_mainnetPublic.ChecksumType))
	return string(PIZ_mainnetPublic.Prefix) + addressEncoder.EncodeData(data, "base58", PIZ_mainnetPublic.Alphabet), nil
}