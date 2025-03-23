package keyutils

import (
	"log"

	"github.com/btcsuite/btcutil/base58"
	"github.com/tyler-smith/go-bip32"
	"golang.org/x/crypto/ripemd160"
)

const derivationPath = "m/44'/194'/0'/0"

func GenerateKeys() (mnemonic string, private string, public string) {

	const mnemonicSeed = "silo feeling crown dumb cake require meadow square hypnotized surf regent quiz"
	masterKey, err := bip32.NewMasterKey([]byte(mnemonicSeed))
	if err != nil {
		log.Fatal(err)
	}

	// Set derivation path master
	hardened44, err := masterKey.NewChildKey(bip32.FirstHardenedChild + 44)
	if err != nil {
		log.Fatal(err)
	}

	// Set derivation second segment
	hardened194, err := hardened44.NewChildKey(bip32.FirstHardenedChild + 194)
	if err != nil {
		log.Fatal(err)
	}

	// Set derivation third segment
	hardened0, err := hardened194.NewChildKey(bip32.FirstHardenedChild + 0)
	if err != nil {
		log.Fatal(err)
	}

	rootKey, err := hardened0.NewChildKey(0)
	if err != nil {
		log.Fatal(err)
	}
	childKey, err := rootKey.NewChildKey(0)
	if err != nil {
		log.Fatal(err)
	}

	privateKeyString := keyToString(childKey.Key, "PVT")
	publicKeyString := keyToString(childKey.PublicKey().Key, "PUB")

	return mnemonicSeed, privateKeyString, publicKeyString

}

func keyToString(keyBytes []byte, prefix string) string {
	hasher := ripemd160.New()
	hasher.Write(keyBytes)
	hasher.Write([]byte("K1"))
	fullChecksum := hasher.Sum(nil)
	checksum := fullChecksum[:4]
	keyWithChecksum := append(keyBytes, checksum...)
	encoded := base58.Encode(keyWithChecksum)
	return prefix + "_K1_" + encoded
}
