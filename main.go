package main

import (
	"context"
	"fmt"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/otcChain/chord-go/chain/types"
	"github.com/otcChain/chord-go/common"
	chordclient "github.com/otcChain/chord-go/rpc/rpc_client"
	"log"
	"math/big"
)

func main() {

	client, err := chordclient.Dial("http://127.0.0.1:6666")
	if err != nil {
		panic(err)
	}
	var privateKey bls.SecretKey
	err = privateKey.DeserializeHexStr("066c6b1a28955a9089670d1e1386484f7370ef7b4f725876e72d82438de06c9e")
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.GetPublicKey()

	fromAddress := common.PubKeyToAddr(publicKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	fmt.Println(fromAddress)
	if err != nil {
		panic(err)
	}

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units

	toAddress, err := common.HexToAddress("fed1gy3afwa745c88dxsznaw82trul3r2p5vsrhmms")
	var data []byte
	//TODO:: make sure the usage of chainID
	//chainID, err := client.NetworkID(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}

	ltx := types.TxData{
		Nonce: nonce,
		To:    &toAddress,
		Value: value,
		Gas:   gasLimit,
		Data:  data,
		Price: nil,
	}
	tx := types.NewTx(ltx)

	if err := tx.SignTx(&privateKey); err != nil {
		panic(err)
	}

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
