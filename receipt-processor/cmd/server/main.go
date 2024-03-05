package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "receipt-processor/internal/api"
)


func main() {
    r := mux.NewRouter()
    // Assuming api.ProcessReceipt and api.GetPoints are correctly defined in your project.
    r.HandleFunc("/receipts/process", api.ProcessReceipt).Methods("POST")
    r.HandleFunc("/receipts/{id}/points", api.GetPoints).Methods("GET")

    log.Println("Server listening on :8080")
    http.ListenAndServe(":8080", r)
}
