package main

import (
	"awesomeProject/cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
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

	h := cyoa.NewHandler(story)
		fmt.Printf("Starting the server on port: %d\n", *port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
	}
