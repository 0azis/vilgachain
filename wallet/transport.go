package wallet

import (
	"bytes"
	"encoding/json"
	"fmt"

	"net/http"
	"vilgachain/mempool"
)

type TransportInterface interface {
	SendCoin(sender string, recipient string, amount int) error
	GenerateKeys() ([]byte, []byte, error)
}

type transport struct {
	http.Client
}

func (t *transport) SendCoin(newTx mempool.Tx) error {
	var txBuffer bytes.Buffer
	err := json.NewEncoder(&txBuffer).Encode(newTx)
	if err != nil {
		return err
	}
	_, err = t.Post("http://192.168.0.5:8080/tx", "application/json", &txBuffer)
	fmt.Println(err)
	return err
}



func NewTransport() transport {
	return transport{}
}
