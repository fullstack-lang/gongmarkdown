package main

import (
	gongmarkdown_models "github.com/fullstack-lang/gongmarkdown/go/models"
)

func createNestedMarkdownStructure(root *gongmarkdown_models.Element, stage *gongmarkdown_models.StageStruct) {

	// Chapter 1
	chapter1 := new(gongmarkdown_models.Element).Stage(stage)
	chapter1.Name = "Chapter 1"
	chapter1.Type = gongmarkdown_models.TITLE
	chapter1.Content = "Introduction to Our Journey"
	root.SubElements = append(root.SubElements, chapter1)

	// Chapter 1 - Subchapter 1
	subchapter1_1 := new(gongmarkdown_models.Element).Stage(stage)
	subchapter1_1.Name = "Subchapter 1.1"
	subchapter1_1.Type = gongmarkdown_models.TITLE
	subchapter1_1.Content = "The Beginning"
	chapter1.SubElements = append(chapter1.SubElements, subchapter1_1)

	// Paragraphs for Subchapter 1.1
	paragraphNames1_1 := []string{
		"First Steps",
		"Setting the Scene",
		"Initial Challenges",
	}
	paragraphContents1_1 := []string{
		"The first paragraph describes the initial context of our story.",
		"This paragraph provides more detailed background information.",
		"The third paragraph sets up the main narrative.",
	}
	addParagraphs(stage, subchapter1_1, paragraphNames1_1, paragraphContents1_1)

	// Chapter 1 - Subchapter 2
	subchapter1_2 := new(gongmarkdown_models.Element).Stage(stage)
	subchapter1_2.Name = "Subchapter 1.2"
	subchapter1_2.Type = gongmarkdown_models.TITLE
	subchapter1_2.Content = "Developing the Story"
	chapter1.SubElements = append(chapter1.SubElements, subchapter1_2)

	// Paragraphs for Subchapter 1.2
	paragraphNames1_2 := []string{
		"New Perspectives",
		"Deepening Narrative",
		"Emerging Themes",
	}
	paragraphContents1_2 := []string{
		"The first paragraph of the second subchapter explores new themes.",
		"This paragraph delves deeper into the developing narrative.",
		"The third paragraph adds complexity to the story.",
	}
	addParagraphs(stage, subchapter1_2, paragraphNames1_2, paragraphContents1_2)

	// Chapter 2
	chapter2 := new(gongmarkdown_models.Element).Stage(stage)
	chapter2.Name = "Chapter 2"
	chapter2.Type = gongmarkdown_models.TITLE
	chapter2.Content = "The Turning Point"
	root.SubElements = append(root.SubElements, chapter2)

	// Chapter 2 - Subchapter 1
	subchapter2_1 := new(gongmarkdown_models.Element).Stage(stage)
	subchapter2_1.Name = "Subchapter 2.1"
	subchapter2_1.Type = gongmarkdown_models.TITLE
	subchapter2_1.Content = "Unexpected Challenges"
	chapter2.SubElements = append(chapter2.SubElements, subchapter2_1)

	// Paragraphs for Subchapter 2.1
	paragraphNames2_1 := []string{
		"Plot Twist",
		"Immediate Impact",
		"Character Responses",
	}
	paragraphContents2_1 := []string{
		"The first paragraph introduces a major plot twist.",
		"This paragraph explores the immediate consequences.",
		"The third paragraph reveals character responses.",
	}
	addParagraphs(stage, subchapter2_1, paragraphNames2_1, paragraphContents2_1)

	// Chapter 2 - Subchapter 2
	subchapter2_2 := new(gongmarkdown_models.Element).Stage(stage)
	subchapter2_2.Name = "Subchapter 2.2"
	subchapter2_2.Type = gongmarkdown_models.TITLE
	subchapter2_2.Content = "Navigating Complexity"
	chapter2.SubElements = append(chapter2.SubElements, subchapter2_2)

	// Paragraphs for Subchapter 2.2
	paragraphNames2_2 := []string{
		"Adaptation",
		"Internal Conflicts",
		"Future Developments",
	}
	paragraphContents2_2 := []string{
		"The first paragraph shows characters adapting to new circumstances.",
		"This paragraph highlights internal conflicts.",
		"The third paragraph sets up future developments.",
	}
	addParagraphs(stage, subchapter2_2, paragraphNames2_2, paragraphContents2_2)
}

// Helper function to add paragraphs to a parent element
func addParagraphs(stage *gongmarkdown_models.StageStruct, parent *gongmarkdown_models.Element, names []string, contents []string) {
	for i := 0; i < len(names); i++ {
		paragraph := new(gongmarkdown_models.Element).Stage(stage)
		paragraph.Name = names[i]
		paragraph.Type = gongmarkdown_models.PARAGRAPH
		paragraph.Content = contents[i]
		parent.SubElements = append(parent.SubElements, paragraph)
	}
}
