// go run server.go
// go build

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("go-server go")

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	fmt.Println("服務執行在端口: 8888\n")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
