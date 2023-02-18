package main

import (
	"encoding/json"
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

	d := json.NewDecoder(f)
	var story cyao.Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}
}