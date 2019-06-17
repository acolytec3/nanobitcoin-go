package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
)

type Block struct {
    Index string `json:"index"`
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
        Block{Index: "1"},
        Block{Index: "2"},
    }
    handleRequests()
}