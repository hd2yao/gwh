package response

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func ReturnJSON(w http.ResponseWriter, status int, data any) error {
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(status)
    return json.NewEncoder(w).Encode(data)
}

func ReturnError(w http.ResponseWriter, status int, err error) {
    ReturnJSON(w, status, map[string]string{"error": err.Error()})
}

func ParseJSON(r *http.Request, payload any) error {
    if r.Body == nil {
        return fmt.Errorf("missing request body")
    }
    return json.NewDecoder(r.Body).Decode(payload)
}
