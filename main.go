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

type ErrorResponse struct {
    Error string `json:"error"`
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
    var reqBody RequestBody

    err := json.NewDecoder(r.Body).Decode(&reqBody)
    if err != nil {
        http.Error(w, `{"error": "Invalid JSON format"}`, http.StatusBadRequest)
        return
    }

    if reqBody.Numbers == nil {
        errorResponse := ErrorResponse{Error: "Missing 'numbers' field in request body"}
        response, _ := json.Marshal(errorResponse)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        w.Write(response)
        return
    }

    for i, num := range reqBody.Numbers {
        if num < -2147483648 || num > 2147483647 {
            errorResponse := ErrorResponse{Error: fmt.Sprintf("Number at index %d is out of range", i)}
            response, _ := json.Marshal(errorResponse)
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
            w.Write(response)
            return
        }
    }

    sum := int32(0)
    for _, num := range reqBody.Numbers {
        sum += num
    }

    w.Header().Set("Content-Type", "application/json")
    jsonResponse := map[string]int32{"sum": sum}
    response, _ := json.Marshal(jsonResponse)
    w.Write(response)
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" 
    }

    http.HandleFunc("/sum", sumHandler)
    log.Printf("Server started on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
