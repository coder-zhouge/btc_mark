package main

import (
	"bytes"
	"encoding/hex"
	"log"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

func main() {
	walletAddress := "mzD2vD8SBhiiJhD4FFZzC8U3jWLc65WxqH"
	costTxHash := "8b6f48dc3398f28a666121687813ec6b88c3df9b219376d9054437302d6da20c"
	costTxOutputN := uint32(1)
	costTxOutputNLockScript := "76a914cd06c3112d4bbb072c6583c041a4ff0c6e68c73088ac"
	priKeyWIF := "<replace by prikeyWIF>"

	var balance int64 = 100000

	var fee int64 = 0.0001 * 1e8 // calc fee
	var leftToMe = balance - fee

	outputs := make([]*wire.TxOut, 0, 2)

	addr, err := btcutil.DecodeAddress(walletAddress, &chaincfg.TestNet3Params)
	MustNilErr(err)
	pkScript, err := txscript.PayToAddrScript(addr)
	MustNilErr(err)
	outputs = append(outputs, wire.NewTxOut(leftToMe, pkScript))

	comment := "周哥教你搞程序：加油！"
	pkScript, err = txscript.NullDataScript([]byte(comment))
	MustNilErr(err)
	outputs = append(outputs, wire.NewTxOut(int64(0), pkScript))

	txHash, err := chainhash.NewHashFromStr(costTxHash)
	MustNilErr(err)

	inputs := []*wire.TxIn{
		wire.NewTxIn(wire.NewOutPoint(txHash, costTxOutputN), nil, nil),
	}

	lockScript, err := hex.DecodeString(costTxOutputNLockScript)
	MustNilErr(err)

	lockScripts := [][]byte{
		lockScript,
	}

	tx := &wire.MsgTx{
		Version:  wire.TxVersion,
		TxIn:     inputs,
		TxOut:    outputs,
		LockTime: 0,
	}

	sign(tx, priKeyWIF, lockScripts)

	buf := bytes.NewBuffer(make([]byte, 0, tx.SerializeSize()))
	err = tx.Serialize(buf)
	MustNilErr(err)

	txHex := hex.EncodeToString(buf.Bytes())
	log.Println("raw tx hex:", txHex)
}

func MustNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func sign(tx *wire.MsgTx, priKeyWIF string, lockScripts [][]byte) {
	inputs := tx.TxIn

	wif, err := btcutil.DecodeWIF(priKeyWIF)
	MustNilErr(err)

	for i := range inputs {
		var script []byte
		script, err = txscript.SignatureScript(tx, i, lockScripts[i], txscript.SigHashAll, wif.PrivKey, true)
		MustNilErr(err)

		inputs[i].SignatureScript = script
	}
}
