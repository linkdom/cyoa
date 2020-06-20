package main

import (
	"awesomeProject/cyoa"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	file := flag.String("file", "story.json", "the JSON with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *file)

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)
	var story cyoa.Story

	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
