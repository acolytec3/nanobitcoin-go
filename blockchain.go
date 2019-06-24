package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "time"
)

type Transaction struct {
    Sender string `json:"sender"`
    Recipient string `json:"recipient"`
    Amount float32 `json:"amount"` 
}

type Block struct {
    Index string `json:"index"`
    Timestamp time.Time `json:"timestamp"`
    Transactions []Transaction `json:"transactions"`
    Previous_hash string `json:"previous_hash"`
}

var Blockchain []Block

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, nanobitcoin-go is coming!")
    fmt.Println("Found our node")
}

func chain(w http.ResponseWriter, r *http.Request){
    json.NewEncoder(w).Encode(Blockchain)
    fmt.Println("/chain found")
}
func handleRequests() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/chain", chain)
    log.Fatal(http.ListenAndServe(":8000", nil))

}
func main() {
    Blockchain = []Block{
        Block{Index: "1", Timestamp: time.Now()},
        Block{Index: "2"},
    }
    handleRequests()
}