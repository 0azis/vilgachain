package mempool

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func InitMempool() {
	store, err := sqlx.Connect("sqlite3", "/home/oazis/mempool")
	if err != nil {
		panic(err)
	}

	controllers := TxRoutes(TxRepo(store))

	http.HandleFunc("/tx", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetAllTxs(w, r)
		case http.MethodPost:
			controllers.AddTxToPool(w, r)
		case http.MethodDelete:
			controllers.RemoveTx(w, r)
		}
	})

	err = http.ListenAndServe("192.168.0.5:8080", nil)
	if err != nil {
		panic(err)
	}
}
