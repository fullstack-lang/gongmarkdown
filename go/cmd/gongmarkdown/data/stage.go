package main

import (
	"time"

	"github.com/fullstack-lang/gongmarkdown/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Content__000000_Dummy := (&models.Content{}).Stage(stage)

	// Setup of values

	__Content__000000_Dummy.Name = `Dummy`
	__Content__000000_Dummy.Content = `
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

	// Setup of pointers
}
