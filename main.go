// main.go

package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WORLD\n")
}
func whoamiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You\n")
}
func newYorkHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "New York")
}
func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprint(w, "Current Time: ", currentTime)
}
func main() {
	http.HandleFunc("/api/hello", helloHandler)
   	http.HandleFunc("/api/whoami", whoamiHandler)
	// Start the server on port 8080
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
	http.HandleFunc("/api/newyork", newYorkHandler)
	fmt.Println("Server listening on :80")
	http.ListenAndServe(":80", nil)
	http.HandleFunc("/api/time", timeHandler)
}
