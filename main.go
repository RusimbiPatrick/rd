package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
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
    http.HandleFunc("/sum", sumHandler)
    log.Println("Server started on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
