package main

import (
	"log"

	"github.com/boltdb/bolt"
)

type BoltStorage struct {
	DB         *bolt.DB
	writerChan chan [3]interface{} // not so agnostic but enough now
}

func (this *BoltStorage) writer() {
	for data := range this.writerChan {
		bucket := data[0].(string)
		key := data[1].(string)
		value := data[2].([]byte)
		err := this.DB.Update(func(tx *bolt.Tx) error {
			sesionBucket, err := tx.CreateBucket([]byte(bucket))
			if err != nil {
				return err
			}
			return sesionBucket.Put([]byte(key), value) // Get, Put, Delete
		})
		if err != nil {
			// TODO: Handle instead of panic
			panic(err)
		}
	}
}

func NewBoltStorage(dbPath string) *BoltStorage {
	db, err := bolt.Open(dbPath, 0666, nil)
	writerChan := make(chan [3]interface{})
	boltStorage := &BoltStorage{DB: db, writerChan: writerChan}

	go boltStorage.writer()
	if err != nil {
		panic(err)
	}
	return boltStorage
}

// simple open and close
func boltOpen(dbpath string) *bolt.DB {
	db, err := bolt.Open(dbpath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func boltClose(db *bolt.DB) {
	db.Close()
}
