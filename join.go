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
    //shares := make([][]byte, 0, len(os.Args))
    var shares [][]byte
    for _, arg := range os.Args[1:] {
        //fmt.Println(i)
        //fmt.Println(arg)
        share, _ := hex.DecodeString(arg)
        //shares = append(shares, share)
        shares = append(shares, share)
    }
    //fmt.Println(shares)
    key, err := shamir.Combine(shares)
    fmt.Println(err)
    fmt.Println(hex.EncodeToString(key))
}
