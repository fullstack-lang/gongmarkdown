package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"gongmarkdown/go/models"
)

var Default uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Element{}),
			Position: &uml.Position{
				X: 680.000000,
				Y: 110.000000,
			},
			Width:  240.000000,
			Heigth: 93.000000,
			Links: []*uml.Link{
				{
					Field: models.Element{}.SubElements,
					Middlevertice: &uml.Vertice{
						X: 800.000000,
						Y: 276.500000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Element{}.Content,
				},
				{
					Field: models.Element{}.Name,
				},
			},
		},
		{
			Struct: &(models.MarkdownContent{}),
			Position: &uml.Position{
				X: 141.000000,
				Y: 106.000000,
			},
			Width:  240.000000,
			Heigth: 93.000000,
			Links: []*uml.Link{
				{
					Field: models.MarkdownContent{}.Root,
					Middlevertice: &uml.Vertice{
						X: 565.500000,
						Y: 149.500000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.MarkdownContent{}.Content,
				},
				{
					Field: models.MarkdownContent{}.Name,
				},
			},
		},
	},
}
