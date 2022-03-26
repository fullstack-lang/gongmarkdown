// insertion point sub template for components imports 
  import { MarkdownContentsTableComponent } from './markdowncontents-table/markdowncontents-table.component'
  import { MarkdownContentSortingComponent } from './markdowncontent-sorting/markdowncontent-sorting.component'
  import { ParagraphsTableComponent } from './paragraphs-table/paragraphs-table.component'
  import { ParagraphSortingComponent } from './paragraph-sorting/paragraph-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfMarkdownContentsComponents: Map<string, any> = new Map([["MarkdownContentsTableComponent", MarkdownContentsTableComponent],])
  export const MapOfMarkdownContentSortingComponents: Map<string, any> = new Map([["MarkdownContentSortingComponent", MarkdownContentSortingComponent],])
  export const MapOfParagraphsComponents: Map<string, any> = new Map([["ParagraphsTableComponent", ParagraphsTableComponent],])
  export const MapOfParagraphSortingComponents: Map<string, any> = new Map([["ParagraphSortingComponent", ParagraphSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["MarkdownContent", MapOfMarkdownContentsComponents],
      ["Paragraph", MapOfParagraphsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["MarkdownContent", MapOfMarkdownContentSortingComponents],
      ["Paragraph", MapOfParagraphSortingComponents],
    ]
  )
