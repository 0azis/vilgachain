// The minimal CLI tool is represent as a VilgaChain wallet
// Minor docs:
// 1. Active Wallet (if you don't have yet): ./vilgawallet active
// Return 2 keys: public (also as an address) and private (the last one we don't recommend to tell anyone)
// 2. Send Transaction: ./vilgawallet -send -from=<your address> -to=<recipient's address> -amount=<sum of coins>
// 200 OK
// 3. My Balance: ./vilgawallet -balance
// Return your actual balace of VilgaCoins

package main

import (
	"flag"
	"fmt"
	"os"
	"vilgachain/mempool"
	"vilgachain/p2p/keys"
	"vilgachain/wallet"
)

func main() {

	httpClient := wallet.NewTransport()

	sendTx := flag.NewFlagSet("sendtx", flag.ExitOnError)
	sender := sendTx.String("sender", "", "the address of sender")
	recipient := sendTx.String("recipient", "", "the address of recipient")
	amount := sendTx.Int("amount", 0, "the amount of the tx")
	privKey := sendTx.String("priv_key", "", "the private key for sign")
	

	// activeWallet := flag.NewFlagSet("active", flag.ExitOnError)

	if len(os.Args) == 1 {
		fmt.Println("VilgaWallet - is the most powerfull wallet for VilgaCoin")
		return
	}	

	switch os.Args[1] {
	case "version":
		fmt.Println("v1.0")
	case "sendtx":
		sendTx.Parse(os.Args[2:])
		if *sender == "" || *recipient == "" || *amount == 0 || *privKey == "" {
			fmt.Println("[ERR]: You have to type whole data.")
			return
		}
		newTx := mempool.Tx{
			Sender: *sender,
			Recipient: *recipient,
			Amount: *amount,
			PrivKey: *privKey,
		}
		err := httpClient.SendCoin(newTx)
		if err != nil {
			fmt.Println(err)
			fmt.Println("[ERR]: Something went wrong")
			return
		}
		fmt.Println("Your transaction was successfully sent")
	case "active":
		privKey, pubKey := keys.GenerateKeys()
		fmt.Printf("Address: %s\nPrivate Key: %s \n", pubKey, privKey)
	}

}
