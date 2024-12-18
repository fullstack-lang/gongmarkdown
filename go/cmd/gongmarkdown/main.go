package main

import (
	"flag"
	"log"
	"strconv"
	"time"

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
		markdownContent := (&gongmarkdown_models.MarkdownContent{Name: "Dummy"}).Stage(stack.Stage)
		markdownContent.Name = "Dummy"

		root := (&gongmarkdown_models.Element{Name: "Root"}).Stage(stack.Stage)
		root.Type = gongmarkdown_models.PARAGRAPH
		root.Content = `
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
		markdownContent.Root = root

		// table := (&gongmarkdown_models.Element{Name: "Table"}).Stage(stack.Stage)
		// table.Content = "Synthetic table"
		// table.Type = gongmarkdown_models.TABLE

		dummyData1 := (&gongmarkdown_models.DummyData{}).Stage(stack.Stage)
		dummyData1.Name = "dummyData1"
		dummyData1.DummyString = "dummyData1 string"
		dummyData1.DummyBool = true
		dummyData1.DummyEnumInt = gongmarkdown_models.ONE
		dummyData1.DummyEnumString = gongmarkdown_models.PARAGRAPH
		dummyData1.DummyDuration = time.Hour + time.Minute
		dummyData1.DummyInt = 42
		dummyData1.DummyFloat = 1.62

		dummyData2 := (&gongmarkdown_models.DummyData{}).Stage(stack.Stage)
		dummyData2.Name = "dummyData2"
		dummyData2.DummyString = "dummyData2 string"
		dummyData2.DummyBool = true
		dummyData2.DummyDuration = 3*time.Hour + 34*time.Minute
		dummyData1.DummyEnumInt = gongmarkdown_models.ONE

		dummyData2.DummyInt = 43
		dummyData2.DummyFloat = 5.77
		dummyData2.DummyEnumInt = gongmarkdown_models.ONE

		another := (&gongmarkdown_models.AnotherDummyData{Name: "another"}).Stage(stack.Stage)
		dummyData1.DummyPointerToGongStruct = another

		table := gongmarkdown_models.GenerateTableOfDummnies(stack.Stage)
		root.SubElements = append(root.SubElements, table)

		// // add a chapter1
		// chapter1 := new(gongmarkdown_models.Element).Stage(stack.Stage)
		// chapter1.Name = "Chapter 1"
		// chapter1.Type = gongmarkdown_models.TITLE
		// chapter1.Content = "This is the title of chapter 1"
		// root.SubElements = append(root.SubElements, chapter1)

		// paragraph := new(gongmarkdown_models.Element).Stage(stack.Stage)
		// paragraph.Name = "paragraph 1"
		// paragraph.Type = gongmarkdown_models.PARAGRAPH
		// paragraph.Content = "This is the content of paragraph 1"
		// chapter1.SubElements = append(chapter1.SubElements, paragraph)

		createNestedMarkdownStructure(root, stack.Stage)
	}

	// fetch the document singloton
	var singloton *gongmarkdown_models.MarkdownContent
	for s := range stack.Stage.MarkdownContents {
		singloton = s
	}
	if singloton == nil {
		singloton = new(gongmarkdown_models.MarkdownContent).Stage(stack.Stage)
	}
	singloton.GenerateContent()

	stack.Stage.Commit()

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))

	if err != nil {
		log.Fatalln(err.Error())
	}
}
