package app

import (
	"log"
	"net/http"
	"time"
)

func main() {
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/status", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("{\"status\": \"ok\"}"))
	})
	log.Fatal(s.ListenAndServe())

}
