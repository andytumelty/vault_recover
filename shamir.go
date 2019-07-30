package main

import (
	// "crypto/aes"
	// "crypto/cipher"
	// "crypto/rand"
	// "crypto/subtle"
	// "encoding/binary"

	//"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/hashicorp/vault/shamir"
)

func main() {
	// enc := "c6932b0b230154b05c28bae8abf1c513a9ba8df61cd76c2133b39254c1d53510"
	enc := os.Args[1]
	key, _ := hex.DecodeString(enc)
	shares, _ := shamir.Split(key, 5, 3)
	keys := make([]string, 0, len(shares))
	// keysB64 := make([]string, 0, len(shares))
	for _, k := range shares {
		keys = append(keys, hex.EncodeToString(k))
		// keysB64 = append(keysB64, base64.StdEncoding.EncodeToString(k))
		fmt.Println(hex.EncodeToString(k))
		// fmt.Println(base64.StdEncoding.EncodeToString(k))
	}
}
