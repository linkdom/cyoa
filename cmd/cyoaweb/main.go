package main

import (
	"awesomeProject/cyoa"
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

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
