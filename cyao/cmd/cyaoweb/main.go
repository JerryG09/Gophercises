package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/JerryG09/Gophercises/cyao"
)

func main() {
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
	fmt.Printf("%+v\n", story)
}