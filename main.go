package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.String("p", "8080", "listen port")
	flag.Parse()

	if *port == "" {
		if p := os.Getenv("PORT"); p != "" {
			*port = p
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		r.Body.Close()

		fmt.Println("\n====== HTTP Request ======")
		fmt.Printf("%s %s %s\n", r.Method, r.URL.String(), r.Proto)
		fmt.Println("Host:", r.Host)
		fmt.Println("RemoteAddr:", r.RemoteAddr)
		fmt.Println("--- Headers ---")
		for k, v := range r.Header {
			fmt.Printf("%s: %v\n", k, v)
		}
		if len(body) > 0 {
			fmt.Println("--- Body ---")
			fmt.Println(string(body))
		}
		fmt.Println("==========================")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Printf("Server running on http://0.0.0.0:%s", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
