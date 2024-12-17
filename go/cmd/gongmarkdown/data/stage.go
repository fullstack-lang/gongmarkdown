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

	__AnotherDummyData__000000_another := (&models.AnotherDummyData{}).Stage(stage)

	__DummyData__000000_dummyData1 := (&models.DummyData{}).Stage(stage)
	__DummyData__000001_dummyData2 := (&models.DummyData{}).Stage(stage)

	__Element__000000_Root := (&models.Element{}).Stage(stage)

	__MarkdownContent__000000_Dummy := (&models.MarkdownContent{}).Stage(stage)

	// Setup of values

	__AnotherDummyData__000000_another.Name = `another`

	__DummyData__000000_dummyData1.Name = `dummyData1`
	__DummyData__000000_dummyData1.DummyString = `dummyData1 string`
	__DummyData__000000_dummyData1.DummyInt = 42
	__DummyData__000000_dummyData1.DummyFloat = 1.620000
	__DummyData__000000_dummyData1.DummyBool = true
	__DummyData__000000_dummyData1.DummyEnumString = models.PARAGRAPH
	__DummyData__000000_dummyData1.DummyEnumInt = models.ONE
	__DummyData__000000_dummyData1.DummyTime, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__DummyData__000000_dummyData1.DummyDuration = 3660000000000

	__DummyData__000001_dummyData2.Name = `dummyData2`
	__DummyData__000001_dummyData2.DummyString = `dummyData2 string`
	__DummyData__000001_dummyData2.DummyInt = 43
	__DummyData__000001_dummyData2.DummyFloat = 5.770000
	__DummyData__000001_dummyData2.DummyBool = true
	__DummyData__000001_dummyData2.DummyEnumInt = models.ONE
	__DummyData__000001_dummyData2.DummyTime, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__DummyData__000001_dummyData2.DummyDuration = 12840000000000

	__Element__000000_Root.Name = `Root`
	__Element__000000_Root.Content = `
# Standalone Markdown Component

This is a **standalone** Angular component using `ngx-markdown`.

## Features
- Renders markdown from string
- Supports remote markdown loading
- Fully configurable

### Code Example
```typescript
const markdown = 'Hello **World**!';
```

- List item 1
- List item 2
`
	__Element__000000_Root.Type = models.PARAGRAPH

	__MarkdownContent__000000_Dummy.Name = `Dummy`
	__MarkdownContent__000000_Dummy.Content = `
# Standalone Markdown Component

This is a **standalone** Angular component using `ngx-markdown`.

## Features
- Renders markdown from string
- Supports remote markdown loading
- Fully configurable

### Code Example
```typescript
const markdown = 'Hello **World**!';
```

- List item 1
- List item 2


`

	// Setup of pointers
	__DummyData__000000_dummyData1.DummyPointerToGongStruct = __AnotherDummyData__000000_another
	__MarkdownContent__000000_Dummy.Root = __Element__000000_Root
}
