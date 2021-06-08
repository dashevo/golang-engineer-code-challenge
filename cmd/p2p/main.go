package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.ServeMux{}
	var items []map[string]interface{}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			time.Sleep(10 * time.Millisecond)
			b, _ := json.Marshal(items)
			_, _ = w.Write(b)
			items = []map[string]interface{}{}
		} else {
			time.Sleep(20 * time.Millisecond)
			var d map[string]interface{}
			_ = json.NewDecoder(r.Body).Decode(&d)
			items = append(items, d)
			_, _ = w.Write([]byte("persisted"))
			fmt.Println("persisted", d)
		}
	})
	_ = http.ListenAndServe(":8881", &mux)
}
