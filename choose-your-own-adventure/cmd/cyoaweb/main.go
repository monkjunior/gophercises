package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/vungocson998/gophercises/choose-your-own-adventure/cyoa"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)
	var story cyoa.Story //We use map[string]Chapter to host JSON data
	d.Decode(&story)

	fmt.Printf("%+v\n", story)
}
