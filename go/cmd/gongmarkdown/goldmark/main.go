package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
)

// printAST recursively prints the AST nodes
func printAST(node ast.Node, depth int) {
	indent := ""
	for i := 0; i < depth; i++ {
		indent += "  "
	}

	fmt.Printf("%sType: %T\n", indent, node)

	// Print kind and text for some node types
	switch n := node.(type) {
	case *ast.Text:
		fmt.Printf("%s  Text: %s\n", indent, string(n.Text(nil)))
	case *ast.Link:
		fmt.Printf("%s  Destination: %s\n", indent, string(n.Destination))
	}

	// Recursively print child nodes
	child := node.FirstChild()
	for child != nil {
		printAST(child, depth+1)
		child = child.NextSibling()
	}
}

func main() {
	// Sample markdown content
	markdownContent := `# Hello Markdown!

This is a **sample** markdown document with:
- A list item
- Another list item

[Link to Goldmark](https://github.com/yuin/goldmark)

> A blockquote example

Inline code: 

## Code Block
` + "```go\n" + `func example() {
    fmt.Println("Hello, World!")
}
` + "```"

	// Create a new markdown parser
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,      // GitHub Flavored Markdown
			extension.Footnote, // Footnote support
		),
		goldmark.WithRendererOptions(
			html.WithXHTML(),
			html.WithUnsafe(), // Allow raw HTML
		),
	)

	// Create a reader from the markdown content
	reader := text.NewReader([]byte(markdownContent))

	// Parse the markdown and create the AST
	document := md.Parser().Parse(reader)

	// Print the AST
	fmt.Println("Abstract Syntax Tree:")
	printAST(document, 0)

	// Render to string
	var stringBuffer bytes.Buffer
	stringRenderer := goldmark.New()
	err := stringRenderer.Renderer().Render(&stringBuffer, []byte(markdownContent), document)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nMarkdown as Plain Text:")
	fmt.Println(stringBuffer.String())

	// Render to HTML
	var htmlBuffer bytes.Buffer
	err = md.Renderer().Render(&htmlBuffer, []byte(markdownContent), document)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nMarkdown as HTML:")
	fmt.Println(htmlBuffer.String())

	// Optionally write HTML to a file
	err = os.WriteFile("output.html", htmlBuffer.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
