package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	// "html"
	"os"

	"github.com/JerryG09/Gophercises/cyao"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the server on")
	fileName := flag.String("file", "gopher.json", "the JSON file with the story")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *fileName)

	f, err:= os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	story, err := cyao.JsonStory(f)

	if err != nil {
		panic(err)
	}
	
	
	h := cyao.NewHandler(story, nil)
	fmt.Printf("Starting the server at: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
	
}
