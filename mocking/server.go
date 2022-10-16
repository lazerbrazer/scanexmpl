package mocking

import (
	"fmt"
	"log"
	"net/http"
)

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
func MockServerRunOn80PortPLS() {
	http.HandleFunc("/", echoString)

	log.Fatal(http.ListenAndServe(":80", nil))

}
