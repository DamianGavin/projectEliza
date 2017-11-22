//Damian Gavin Eliza Project
//14/11/17
//Adapted from https://golang.org/doc/articles/wiki/
//And https://github.com/data-representation/go-ajax/blob/master/webapp.go
package main

import (
	"fmt"
	"net/http"

	"./eliza"
)

func main() {

	dir := http.Dir("./static")//static folder 

	fileServer := http.FileServer(dir)

	// Handle "/" the root resource
	// "/" handles EVERYTHING coming in, unless there is a more specific path eg /ask
	http.Handle("/", fileServer)

	// "/ask" resource
	http.HandleFunc("/ask", HandleAsk)

	// actually starts the web server!
	http.ListenAndServe(":8080", nil)
}
func HandleAsk(writer http.ResponseWriter, request *http.Request) {
	userInput := request.URL.Query().Get("userInput")
	answer := eliza.Ask(userInput) // takes the input we got from the request, and sends it to the Ask function
	fmt.Fprintf(writer, answer)    // writes the result back into the ResponseWriter

}
