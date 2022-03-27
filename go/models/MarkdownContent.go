package models

type MarkdownContent struct {
	Name    string
	Content string

	Root *Element
}

// UpdateContent updates Content from the root
func (markdownContent *MarkdownContent) UpdateContent() {

	markdownContent.Content = ""

	// verification that there are no cycles in the graph

	if markdownContent.Root != nil {
		markdownContent.RecursiveUpdateContent(markdownContent.Root, 1)
	}

}

func (markdownContent *MarkdownContent) RecursiveUpdateContent(element *Element, depth int) {

	switch element.GetType() {
	case TITLE:
		for level := 0; level < depth; level++ {
			markdownContent.Content = markdownContent.Content + "#"
		}
		markdownContent.Content = markdownContent.Content + " "
	}

	markdownContent.Content = markdownContent.Content + element.GetContent()
	markdownContent.Content = markdownContent.Content + "\n\n"

	for _, subElement := range element.GetSubElements() {
		markdownContent.RecursiveUpdateContent(subElement, depth+1)
	}

}
