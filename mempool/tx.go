package mempool

import "github.com/jmoiron/sqlx"


type TxStoreInterface interface {
	Insert(tx Tx) error
	Delete(txID string) error
	Select() ([]Tx, error)
}

type txStore struct {
	db *sqlx.DB
}

type Tx struct {
	ID        int
	Sender    string
	Recipient string
	Amount    int
	PrivKey string// the secret info from user. it don't have to be in the mempool after.
}

func (ts txStore) Insert(tx Tx) error {
	_, err := ts.db.Exec(`insert into txs (sender, recipient, amount) values ($1, $2, $3)`, tx.Sender, tx.Recipient, tx.Amount)
	return err
}

func (ts txStore) Delete(txID string) error {
	_, err := ts.db.Exec(`delete from txs where id = $1`, txID)
	return err
}

func(ts txStore) Select() ([]Tx, error) {
	var txs []Tx

	err := ts.db.Select(&txs, `select * from txs`)

	return txs, err
}


func TxRepo(store *sqlx.DB) TxStoreInterface {
	return txStore{
		db: store,
	}
}

