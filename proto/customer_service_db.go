package proto

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

/*
	Binary and JSON conversion
*/
func (d *Person) GobEncode() ([]byte, error) {
	w := new(bytes.Buffer)
	encoder := gob.NewEncoder(w)
	err := encoder.Encode(d.Id)
	if err != nil {
		return nil, err
	}
	err = encoder.Encode(d.Name)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func (d *Person) GobDecode(buf []byte) error {
	r := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(r)
	err := decoder.Decode(&d.Id)
	if err != nil {
		return err
	}
	err = decoder.Decode(&d.Name)
	if err != nil {
		return err
	}
	return nil
}

func ExampleConvertUse() {
	d := Person{Id: 7, Name: "stoney"}

	// writing
	buffer := new(bytes.Buffer)
	enc := gob.NewEncoder(buffer)
	err := enc.Encode(d)
	if err != nil {
		log.Fatal("encode error:", err)
	}

	// reading
	p := new(Person)
	buffer = bytes.NewBuffer(buffer.Bytes())
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(p)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Println(p, err)
}

/*
	BoltDB utility functions
*/
func BoltOpen(dbpath string, dbopt *bolt.Options) (*bolt.DB, error) {
	return bolt.Open(dbpath, 0644, dbopt)
}

func BoltClose(db *bolt.DB) {
	db.Close()
}

func BoltMonitor(db *bolt.DB, ts time.Duration) {
	// Grab the initial stats.
	prev := db.Stats()

	for {
		// Wait for 10s.
		time.Sleep(ts)

		// Grab the current stats and diff them.
		stats := db.Stats()
		diff := stats.Sub(&prev)

		// Encode stats to JSON and print to STDERR.
		json.NewEncoder(os.Stderr).Encode(diff)

		// Save stats for the next loop.
		prev = stats
	}
}

func BoltState(db *bolt.DB) {
	stat := db.Stats()
	json.NewEncoder(os.Stderr).Encode(stat)
}

func BoltListAll(db *bolt.DB) error {
	err := db.View(func(tx *bolt.Tx) error {
		err := tx.ForEach(func(name []byte, b *bolt.Bucket) error {
			fmt.Printf("bk:%q\n", name)
			BoltListBucket(db, name)
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatalf("%v %T\n", err, err)
	}

	return nil
}

func BoltListBucket(db *bolt.DB, bucket []byte) error {
	err := db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(bucket).Cursor()
		for k, v := c.Last(); k != nil; k, v = c.Prev() {
			fmt.Printf("\tk:%q, v:%q\n", k, v)
		}
		return nil
	})

	if err != nil {
		log.Fatalf("%v %T\n", err, err)
	}

	return nil
}

func BoltDeleteBucket(db *bolt.DB, bucket []byte) error {
	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket(bucket)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatalf("%v %T\n", err, err)
	}

	return nil
}

func BoltGetItem(db *bolt.DB, bucket, key []byte) ([]byte, error) {
	var value []byte

	err := db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket(bucket)
		if bk == nil {
			return fmt.Errorf("Bucket %q not found!", bk)
		}

		value = bk.Get(key)
		if value == nil {
			log.Printf("Key %q not found\n", key)
		}
		fmt.Println(string(value))

		return nil
	})

	if err != nil {
		log.Fatalf("%v %T\n", err, err)
	}

	return value, err
}

func BoltPutItem(db *bolt.DB, bucket, key, value []byte) error {
	err := db.Update(func(tx *bolt.Tx) error {
		bk, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}

		err = bk.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatalf("%v %T\n", err, err)
	}

	return nil
}

func BoltDeleteItem(db *bolt.DB, bucket, key []byte) error {
	err := db.Update(func(tx *bolt.Tx) error {
		bk := tx.Bucket(bucket)
		if bk == nil {
			return fmt.Errorf("Bucket %q not found!", bk)
		}

		err := bk.Delete(key)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatalf("%v %T\n", err, err)
	}

	return nil
}
