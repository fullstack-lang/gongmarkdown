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

This is a **standalone** Angular component using ` + "`" + "ngx-markdown" + "`." + `

## Features
- Renders markdown from string
- Supports remote markdown loading
- Fully configurable

### Code Example
` + "```" + `typescript
const markdown = 'Hello **World**!';
` + "```" + `

- List item 1
- List item 2
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

		// table := gongmarkdown_models.GenerateTableOfDummnies()
		// root.SubElements = append(root.SubElements, table)

		// generic call
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
