package handlers

import (
	"encoding/json"
	"net/http"
	"vilgachain/mining_pool/internal/store"
	"vilgachain/mining_pool/pkg"
)

type MinerHandlers interface {
	ConnectToMiningPool(w http.ResponseWriter, r *http.Request)
	DisconnectFromMiningPool(w http.ResponseWriter, r *http.Request)
	ActiveMiners(w http.ResponseWriter, r *http.Request)
}

type Miner struct {
	MinerRepo store.MinerInterface
}

func (m Miner) ConnectToMiningPool(w http.ResponseWriter, r *http.Request) {
	ipAddr :=pkg.CutIPAddress(r.RemoteAddr)

	err := m.MinerRepo.Insert(ipAddr)

	if err != nil {
		http.Error(w, "An Error while connecting you to the network", http.StatusInternalServerError)
		return
	}

	http.Error(w, "You was connected to the network", http.StatusOK)
}

func (m Miner) DisconnectFromMiningPool(w http.ResponseWriter, r *http.Request) {
	ipAddr := pkg.CutIPAddress(r.RemoteAddr)

	node, err := m.MinerRepo.GetIP(ipAddr)

	if err != nil {
		http.Error(w, "An Error while disconnecting you from the network", http.StatusInternalServerError)
		return
	}

	err = m.MinerRepo.DeleteNode(node.IPAddr)

	if err != nil {
		http.Error(w, "An Error while disconnecting you from the network", http.StatusInternalServerError)
		return 
	}
	
	http.Error(w, "You was disconnected from the network", http.StatusOK)
}


func (m Miner) ActiveMiners(w http.ResponseWriter, r *http.Request) {
	nodes, err := m.MinerRepo.Select()
	if err != nil {
		http.Error(w, "An Error while collecting all acitve users", http.StatusInternalServerError)
		return 
	}
	JSON, _ := json.Marshal(nodes)
	w.Write(JSON)
}

func GetMinerHandlers(miner store.MinerInterface) Miner {
	return Miner{
		MinerRepo: miner,
	}
}
