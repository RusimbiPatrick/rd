package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
)

type RequestBody struct {
    Numbers []int32 `json:"numbers"`
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
    var reqBody RequestBody

    err := json.NewDecoder(r.Body).Decode(&reqBody)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    sum := int32(0)
    for _, num := range reqBody.Numbers {
        sum += num
    }

    fmt.Fprintf(w, "%d", sum)
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // default port
    }

    http.HandleFunc("/sum", sumHandler)
    log.Printf("Server started on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
