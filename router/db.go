package main

import (
    "log"
    "fmt"
    "github.com/boltdb/bolt"
)

func Open() *bolt.DB{

  db, err := bolt.Open("my.db", 0600, nil)
  if err != nil {
      log.Fatal(err)
  }
  defer db.Close()
  return db
}

func CreateBucket(db *bolt.DB, bucketName string) {
   log.Println(db)
  db.Update(func(tx *bolt.Tx) error {
    b, err := tx.CreateBucket([]byte(bucketName))
    if b != nil {
      log.Println(b)
    }
    if err != nil {
        return fmt.Errorf("create bucket: %s", err)
    }
    return nil
})
}

func SaveBucket(db *bolt.DB, bucketName string, key string, value string) {
  db.Update(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte(bucketName))
    err := b.Put([]byte(key), []byte(value))
    return err
})
}

func ReadBucket(db *bolt.DB, bucketName string, key string) {
db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte(bucketName))
    v := b.Get([]byte(key))
    fmt.Printf("The answer is: %s\n", v)
    return nil
})
}
