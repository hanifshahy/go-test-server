package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Secure Go server is live!"))
	})

	err := http.ListenAndServeTLS(":443", `C:\Users\DELLL\Documents\projects\test\test2\ssl.cert`, `C:\Users\DELLL\Documents\projects\test\test2\domain.key`, nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS error: ", err)
	}
}
