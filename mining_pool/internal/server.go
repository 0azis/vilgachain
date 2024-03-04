package internal

import (
	"fmt"
	"net/http"
	"vilgachain/mining_pool/internal/routes"
	"vilgachain/mining_pool/internal/store"
)

// vilgachain will be working in my local network

func InitServer() {
	db, err := store.InitStore()
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	
	routes.NodesRoutes(db)
	http.ListenAndServe("192.168.0.5:8080", nil)
}
