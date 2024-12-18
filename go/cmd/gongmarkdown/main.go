package main

import (
	"flag"
	"log"
	"strconv"

	gongmarkdown_models "github.com/fullstack-lang/gongmarkdown/go/models"
	gongmarkdown_stack "github.com/fullstack-lang/gongmarkdown/go/stack"
	gongmarkdown_static "github.com/fullstack-lang/gongmarkdown/go/static"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")

	generateDocument = flag.Bool("generateDocument", false, "generates a markdown document")
)

func main() {

	log.SetPrefix("gongmarkdown: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := gongmarkdown_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	stack := gongmarkdown_stack.NewStack(r, "gongmarkdown", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	if *generateDocument {
		content := (&gongmarkdown_models.Content{Name: "Dummy"}).Stage(stack.Stage)
		content.Name = "Dummy"
		content.Content = `
# Standalone Markdown Component

This is a **standalone** Angular component using "ngx-markdown"

## Features
- Renders markdown from string
- Supports remote markdown loading
- Fully configurable

- List item 1
- List item 2

![Local Angular Logo](assets/images/image.png)

<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="100" height="100" fill="gold">
  <polygon points="12,2 15.09,8.26 22,9.27 17,14.14 18.18,21.02 12,17.77 5.82,21.02 7,14.14 2,9.27 8.91,8.26" />
</svg>

<img src=/assets/images/star.svg width=50 />

`

	}

	stack.Stage.Commit()

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))

	if err != nil {
		log.Fatalln(err.Error())
	}
}
