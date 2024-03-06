package mempool

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"vilgachain/p2p/keys"
)

func Decoder(body io.ReadCloser) (Tx, error) {
	// decoded transaction
	var decodedTx Tx
	err := json.NewDecoder(body).Decode(&decodedTx)

	return decodedTx, err
}

type TxControllers struct {
	txRepo TxStoreInterface
}

func (tx TxControllers) AddTxToPool(w http.ResponseWriter, r *http.Request) {
	body, err := Decoder(r.Body)
	if err != nil {
		http.Error(w, "An error while decode JSON", http.StatusBadRequest)
		return
	}

	status := keys.Verify(body.PrivKey, body.Sender)
	fmt.Println(status)
	if !status {
		fmt.Println("UNVALID SIGN")
		http.Error(w, "Unvalid sign", http.StatusUnauthorized)
		return
	}

	err = tx.txRepo.Insert(body)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "An error while insert your transaction in pool", http.StatusInternalServerError)
		return
	}

	http.Error(w, "Transaction was successfully added to the pool", http.StatusCreated)
}

func (tx TxControllers) GetAllTxs(w http.ResponseWriter, r *http.Request) {
	txs, err := tx.txRepo.Select()
	if err != nil {
		http.Error(w, "An error while get transactions", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(txs)
}

func (tx TxControllers) RemoveTx(w http.ResponseWriter, r *http.Request) {
	queryID := r.URL.Query().Get("id")

	err := tx.txRepo.Delete(queryID)

	if err != nil {
		http.Error(w, "An error while deleting transaction", http.StatusInternalServerError)
		return
	}

	http.Error(w, "Transaction was successfully deleted from the pool", http.StatusOK)
}

func TxRoutes(repo TxStoreInterface) TxControllers {
	return TxControllers{
		txRepo: repo,
	}
}
