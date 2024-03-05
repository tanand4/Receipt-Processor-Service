package api

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "receipt-processor/internal/processor"
    "log"
)

// ProcessReceipt handles the POST request to process receipts.
// func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
//     if r.Method != http.MethodPost {
//         http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//         return
//     }

//     // Parse the JSON body.
//     var receipt processor.Receipt
//     if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
//         http.Error(w, err.Error(), http.StatusBadRequest)
//         return
//     }

//     // Process the receipt and generate an ID.
//     id, err := processor.Process(&receipt)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }

//     // Return the ID in JSON format.
//     w.Header().Set("Content-Type", "application/json")
//     json.NewEncoder(w).Encode(struct{ ID string `json:"id"` }{ID: id})
// }

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
    // Function remains unchanged as it correctly handles a POST request with JSON body.
    log.Println("ProcessReceipt called")
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Parse the JSON body.
    var receipt processor.Receipt
    if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Process the receipt and generate an ID.
    id, err := processor.Process(&receipt)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Return the ID in JSON format.
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(struct{ ID string `json:"id"` }{ID: id})
}

// GetPoints handles the GET request to retrieve points by receipt ID.
// func GetPoints(w http.ResponseWriter, r *http.Request) {
//     if r.Method != http.MethodGet {
//         http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//         return
//     }

//     // Extract ID from URL.
//     id := strings.TrimPrefix(r.URL.Path, "/receipts/{id}/points")
//     points, err := processor.GetPoints(id)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusNotFound)
//         return
//     }

//     // Return the points in JSON format.
//     w.Header().Set("Content-Type", "application/json")
//     json.NewEncoder(w).Encode(struct{ Points int `json:"points"` }{Points: points})
// }

func GetPoints(w http.ResponseWriter, r *http.Request) {
    log.Println("GetPoints called")
    vars := mux.Vars(r)
    id := vars["id"] // Get the id parameter from the URL

    points, err := processor.GetPoints(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(struct{ Points int `json:"points"` }{Points: points})
}