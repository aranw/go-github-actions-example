package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, sayHello("World"))
	}))
}

func sayHello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
