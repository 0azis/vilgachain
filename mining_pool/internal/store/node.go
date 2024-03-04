package store

import (
	"vilgachain/mining_pool/internal/models"

	"github.com/jmoiron/sqlx"
)

type MinerInterface interface {
	Insert(ipAddr string) error
	Select() ([]models.Miner, error)
	GetIP(ipAddr string) (models.Miner, error)
	DeleteNode(ipAddr string) error
}

type miner struct {
	db *sqlx.DB
}

func (m miner) Insert(ipAddr string) error {
	_, err := m.db.Exec(`insert into nodes (ipaddr) values ($1) on conflict (ipaddr) do nothing`, ipAddr)
	return err
}

func (m miner) Select() ([]models.Miner, error) {
	var miners []models.Miner
	err := m.db.Select(&miners, `select * from nodes`)
	return miners, err
}

func (m miner) GetIP(ipAddr string) (models.Miner, error) {
	var miner models.Miner
	err := m.db.Get(&miner, `select * from nodes where ipaddr = $1`, ipAddr)
	return miner, err
}

func (m miner) DeleteNode(ipAddr string) error {
	_, err := m.db.Exec(`delete from nodes where ipaddr = $1`, ipAddr)
	return err
}

func MinerRepo(db sqlx.DB) miner {
	return miner{
		db: &db,
	}
}
