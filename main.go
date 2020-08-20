package main

import (
	"fmt"

	"github.com/brianium/mnemonic"
)

func main() {

    // generate a random Mnemonic in English with 256 bits of entropy
    m, _ := mnemonic.NewRandom(256, mnemonic.English)

    // print the Mnemonic as a sentence
    fmt.Println(m.Sentence())

    // inspect underlying words
    fmt.Println(m.Words)

    // generate a seed from the Mnemonic
    seed := m.GenerateSeed("passphrase")

    // print the seed as a hex encoded string
    fmt.Println(seed)

}