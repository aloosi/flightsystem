package main

/*
import (
    "encoding/json"
    "net/http"
)

func createFlightHandler(w http.ResponseWriter, r *http.Request) {
    var flight Flight
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&flight); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    if err := createFlight(flight); err != nil {
        http.Error(w, "Failed to create flight", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}
*/
