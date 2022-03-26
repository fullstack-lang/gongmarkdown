// insertion point sub template for components imports 
  import { ParagraphsTableComponent } from './paragraphs-table/paragraphs-table.component'
  import { ParagraphSortingComponent } from './paragraph-sorting/paragraph-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfParagraphsComponents: Map<string, any> = new Map([["ParagraphsTableComponent", ParagraphsTableComponent],])
  export const MapOfParagraphSortingComponents: Map<string, any> = new Map([["ParagraphSortingComponent", ParagraphSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Paragraph", MapOfParagraphsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Paragraph", MapOfParagraphSortingComponents],
    ]
  )
