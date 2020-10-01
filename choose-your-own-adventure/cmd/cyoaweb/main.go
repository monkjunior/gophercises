package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	_ = f
	// d := json.NewDecoder(f)
	// var story cyoa.Story
}
