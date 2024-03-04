package sender

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Sender struct {
	db *sqlx.DB
}

func (s *Sender) SendBlockchain() {
	var ipAddrs []string
	s.db.Select(&ipAddrs, `select * from nodes`)

	for addr := range ipAddrs {
		// curl network testing
		resp, _ := http.Get(ipAddrs[addr])
		resp.Body.Close()
	}
}

