package routes

import (
	"net/http"
	"vilgachain/mining_pool/internal/handlers"
	"vilgachain/mining_pool/internal/store"

	"github.com/jmoiron/sqlx"
)

func NodesRoutes(db *sqlx.DB) {
	controllers := handlers.GetMinerHandlers(store.MinerRepo(*db))
	http.HandleFunc("/miner", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.ActiveMiners(w, r)
		case http.MethodPost:
			controllers.ConnectToMiningPool(w, r)
		case http.MethodDelete:
			controllers.DisconnectFromMiningPool(w, r)
		}
	})
}
