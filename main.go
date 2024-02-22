package main

import (
	"fmt"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, sayHello("World"))
	})); err != nil {
		panic(err)
	}
}

func sayHello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
