package main

import (
	"log"

	"github.com/boltdb/bolt"
)

type BoltStorage struct {
	DB         *bolt.DB
	writerChan chan [3]interface{} //not so agnostic but enough now
}

func NewBoltStorage(dbPath string) *BoltStorage {
	db, err := bolt.Open(dbPath, 0666, nil)
	writerChan := make(chan [3]interface{})
	boltStorage := &BoltStorage{DB: db, writerChan: writerChan}

	//go boltStorage.writer()
	if err != nil {
		panic(err)
	}
	return boltStorage
}

func boltOpen(dbpath string) *bolt.DB {
	db, err := bolt.Open(dbpath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	return db
}

func boltClose(db *bolt.DB) {
	db.Close()
}
