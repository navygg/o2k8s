package main

import (
	"fmt"
	"log"
	"net/http"
)

// HiHandle response /hi
func HiHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := r.Form["name"][0]

	w.Write([]byte(fmt.Sprintf("Hi %s\n", user)))
}

func main() {
	http.HandleFunc("/hi", HiHandle)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
