package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/fullstack-lang/gongmarkdown/go/cmd/gongmarkdown/markdown"
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

<img src=/assets/images/star.svg width=50 />

`
		content.Content += markdown.GenerateMarkdownSample()

	}

	stack.Stage.Commit()

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))

	if err != nil {
		log.Fatalln(err.Error())
	}
}
