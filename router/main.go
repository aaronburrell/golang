package main

import (
    "log"
    "net/http"
)

func main() {

    router := NewRouter()
  //  db := Open()
  //  CreateBucket(db, "MyBucket")
  //  SaveBucket(db,"MyBucket", "test", "I did it!")
  //  ReadBucket(db,"MyBucket","test")
    log.Fatal(http.ListenAndServe(":8080", router))
}
