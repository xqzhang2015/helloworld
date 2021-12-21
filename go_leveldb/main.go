package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	fmt.Printf("Hello\n")
	// The returned DB instance is safe for concurrent use. Which mean that all
	// DB's methods may be called concurrently from multiple goroutine.
	db, err := leveldb.OpenFile("path/to/db", nil)
	if err != nil {
		fmt.Printf("failed to open leveldb: %v\n", err)
	}
	defer db.Close()

	// Remember that the contents of the returned slice should not be modified.
	data, err := db.Get([]byte("key: xq"), nil)
	if err != nil {
		fmt.Printf("db.Get: %v\n", err)
		// return
	} else {
		fmt.Printf("data: %s\n", data)
	}

	for i := 0; i < 10; i++ {
		err = db.Put([]byte(fmt.Sprintf("key: xq - %d", i)), []byte("value: xq"), nil)
		fmt.Printf("db.Put: %v\n", err)
	}

	data, err = db.Get([]byte("key: xq"), nil)
	if err != nil {
		fmt.Printf("db.Get: %v\n", err)
	} else {
		fmt.Printf("db.Get data: %s\n", data)
	}

	err = db.Delete([]byte("key"), nil)
	fmt.Printf("db.Delete-1: %v\n", err)

	// err = db.Delete([]byte("key: xq"), nil)
	// fmt.Printf("db.Delete-2: %v\n", err)
}
