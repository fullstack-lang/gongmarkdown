// insertion point sub template for components imports 
  import { ElementsTableComponent } from './elements-table/elements-table.component'
  import { ElementSortingComponent } from './element-sorting/element-sorting.component'
  import { MarkdownContentsTableComponent } from './markdowncontents-table/markdowncontents-table.component'
  import { MarkdownContentSortingComponent } from './markdowncontent-sorting/markdowncontent-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfElementsComponents: Map<string, any> = new Map([["ElementsTableComponent", ElementsTableComponent],])
  export const MapOfElementSortingComponents: Map<string, any> = new Map([["ElementSortingComponent", ElementSortingComponent],])
  export const MapOfMarkdownContentsComponents: Map<string, any> = new Map([["MarkdownContentsTableComponent", MarkdownContentsTableComponent],])
  export const MapOfMarkdownContentSortingComponents: Map<string, any> = new Map([["MarkdownContentSortingComponent", MarkdownContentSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Element", MapOfElementsComponents],
      ["MarkdownContent", MapOfMarkdownContentsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Element", MapOfElementSortingComponents],
      ["MarkdownContent", MapOfMarkdownContentSortingComponents],
    ]
  )
