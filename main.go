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

	dir := http.Dir("./static")

	fileServer := http.FileServer(dir)

	// Handle "/" the root resource
	// "/" handles EVERYTHING coming in, unless there is a more specific path eg /ask
	http.Handle("/", fileServer)

	// "/ask" resource
	http.HandleFunc("/ask", handleAsk)

	// actually starts the web server!
	http.ListenAndServe(":8080", nil)
}
func handleAsk(writer http.ResponseWriter, request *http.Request) {
	// fmt.Fprintln(writer, "hi there! You went to /ask")
	// fmt.Fprintln(writer, "<h1> this is a web site!</h1>") go sends a plain string -> browser knows what to do with it as it is just html
	// ask eliza a question
	// write the answer to the writer
	// "user-input"
	// example URL
	//
	userInput := request.URL.Query().Get("user-input")
	answer := eliza.Ask(userInput) // takes the input we got from the request, and sends it to the Ask function
	fmt.Fprintf(writer, answer)    // writes the result back into the ResponseWriter

}
