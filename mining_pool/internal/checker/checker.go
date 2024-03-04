package checker

import (
	"net/http"
	"github.com/jmoiron/sqlx"
)

type Checker struct {
	db *sqlx.DB
}

func (c *Checker) AnalyzeDNS() {
	var ipAddrs []string
	c.db.Select(&ipAddrs, `select * from nodes`)

	for addr := range ipAddrs {
		// curl network testing
		resp, err := http.Get(ipAddrs[addr])
		if err != nil {
			c.db.Query(`delete from nodes where ipaddr = $1`, addr)
		}
		defer resp.Body.Close()
	}

}
