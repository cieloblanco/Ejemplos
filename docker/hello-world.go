package main

import(
        "fmt"
        "log"
        "net/http"
        "os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)	
	log.Fatal(http.ListenAndServe(":8080", mux))	
}

func hello(w http.ResponseWriter, r *http.Request){
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	fmt.Fprintf(w, "Hello, world!\n")
	fmt.Fprintf(w, "Hostname: %s\n", host)
}
