package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

func main() {
	file := "testnet3_wallet.dat"
	_, err := os.Lstat(file)
	if !os.IsNotExist(err) {
		log.Println("wallet file exists, skip")

		return
	}

	priKey, err := btcec.NewPrivateKey()
	if err != nil {
		log.Fatalln(err)
	}

	priKeyWif, err := btcutil.NewWIF(priKey, &chaincfg.TestNet3Params, true)
	if err != nil {
		log.Fatalln(err)
	}

	address, err := btcutil.NewAddressPubKey(priKey.PubKey().SerializeUncompressed(), &chaincfg.TestNet3Params)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("prikey: %s\n", priKeyWif)
	log.Printf("address: %s\n", address.EncodeAddress())

	_ = ioutil.WriteFile(file, []byte(fmt.Sprintf(
		"prikey: %s\naddress: %s\n", priKeyWif.String(), address.EncodeAddress())), os.ModePerm)
}
