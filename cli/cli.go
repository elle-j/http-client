package cli

import (
	"fmt"
	"http-client/http"
)

func Run() {
	// TODO

	fmt.Println("Running CLI..")

	fmt.Println("Requesting fake URL...")
	response := http.Get("https://todo.com")
	fmt.Println(response)
}
