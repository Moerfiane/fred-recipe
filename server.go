package main

import (
	"fmt"
	"net/http"
	"os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func getPort() string {
	p := os.Getenv("Port")
	if p != "" {
		return ":" + p
	}
	return ":1234"
}

func main() {
	port := getPort()
	http.HandleFunc("/", sayHello)
	fmt.Println("Listenting on port" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
