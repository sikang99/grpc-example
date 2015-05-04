package proto

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

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
