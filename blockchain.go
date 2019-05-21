package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
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
//    fmt.Fprintf(w, "There will be a blockchain here someday")
    json.NewEncoder(w).Encode(Blockchain)
    fmt.Println("/chain found")
}
func handleRequests() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", handler)
    router.HandleFunc("/chain", chain)
//    http.HandleFunc("/", handler)
//    http.HandleFunc("/chain", chain)
    log.Fatal(http.ListenAndServe(":8000", router))

}
func main() {
    Blockchain = []Block{
        Block{Index: "1"},
        Block{Index: "2"},
    }
    handleRequests()
}