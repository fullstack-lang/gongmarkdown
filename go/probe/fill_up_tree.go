package probe

import (
	"fmt"
	"sort"
	"strings"

	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	tree "github.com/fullstack-lang/gongtree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongmarkdown/go/models"
)

func fillUpTree(
	probe *Probe,
) {
	// keep in memory which nodes have been unfolded / folded
	expandedNodesSet := make(map[string]any, 0)
	var _sidebar *tree.Tree
	for __sidebar := range probe.treeStage.Trees {
		_sidebar = __sidebar
	}
	if _sidebar != nil {
		for _, node := range _sidebar.RootNodes {
			if node.IsExpanded {
				expandedNodesSet[strings.Fields(node.Name)[0]] = true
			}
		}
	}

	probe.treeStage.Reset()

	// create tree
	sidebar := (&tree.Tree{Name: "gong"}).Stage(probe.treeStage)

	// collect all gong struct to construe the true
	setOfGongStructs := *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](probe.gongStage)

	sliceOfGongStructsSorted := make([]*gong_models.GongStruct, len(setOfGongStructs))
	i := 0
	for k := range setOfGongStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return sliceOfGongStructsSorted[i].Name < sliceOfGongStructsSorted[j].Name
	})

	for _, gongStruct := range sliceOfGongStructsSorted {

		name := gongStruct.Name + " (" +
			fmt.Sprintf("%d", probe.stageOfInterest.Map_GongStructName_InstancesNb[gongStruct.Name]) + ")"

		nodeGongstruct := (&tree.Node{Name: name}).Stage(probe.treeStage)

		nodeGongstruct.IsExpanded = false
		if _, ok := expandedNodesSet[strings.Fields(name)[0]]; ok {
			nodeGongstruct.IsExpanded = true
		}

		switch gongStruct.Name {
		// insertion point
		case "AnotherDummyData":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.AnotherDummyData](probe.stageOfInterest)
			for _anotherdummydata := range set {
				nodeInstance := (&tree.Node{Name: _anotherdummydata.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_anotherdummydata, "AnotherDummyData", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Cell":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Cell](probe.stageOfInterest)
			for _cell := range set {
				nodeInstance := (&tree.Node{Name: _cell.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_cell, "Cell", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DummyData":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DummyData](probe.stageOfInterest)
			for _dummydata := range set {
				nodeInstance := (&tree.Node{Name: _dummydata.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_dummydata, "DummyData", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Element":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Element](probe.stageOfInterest)
			for _element := range set {
				nodeInstance := (&tree.Node{Name: _element.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_element, "Element", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "MarkdownContent":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.MarkdownContent](probe.stageOfInterest)
			for _markdowncontent := range set {
				nodeInstance := (&tree.Node{Name: _markdowncontent.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_markdowncontent, "MarkdownContent", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Row":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Row](probe.stageOfInterest)
			for _row := range set {
				nodeInstance := (&tree.Node{Name: _row.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_row, "Row", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		}

		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewTreeNodeImplGongstruct(gongStruct, probe)

		// add add button
		addButton := (&tree.Button{
			Name: gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon: string(gongtree_buttons.BUTTON_add)}).Stage(probe.treeStage)
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstruct(
			gongStruct,
			gongtree_buttons.BUTTON_add,
			probe,
		)

		sidebar.RootNodes = append(sidebar.RootNodes, nodeGongstruct)
	}

	// Add a refresh button
	nodeRefreshButton := (&tree.Node{Name: ""}).Stage(probe.treeStage)
	sidebar.RootNodes = append(sidebar.RootNodes, nodeRefreshButton)
	refreshButton := (&tree.Button{
		Name: "RefreshButton" + " " + string(gongtree_buttons.BUTTON_refresh),
		Icon: string(gongtree_buttons.BUTTON_refresh)}).Stage(probe.treeStage)
	nodeRefreshButton.Buttons = append(nodeRefreshButton.Buttons, refreshButton)
	refreshButton.Impl = NewButtonImplRefresh(probe)

	probe.treeStage.Commit()
}

type InstanceNodeCallback[T models.Gongstruct] struct {
	Instance       *T
	gongstructName string
	probe          *Probe
}

func NewInstanceNodeCallback[T models.Gongstruct](
	instance *T,
	gongstructName string,
	probe *Probe) (
	instanceNodeCallback *InstanceNodeCallback[T],
) {
	instanceNodeCallback = new(InstanceNodeCallback[T])

	instanceNodeCallback.probe = probe
	instanceNodeCallback.gongstructName = gongstructName
	instanceNodeCallback.Instance = instance

	return
}

func (instanceNodeCallback *InstanceNodeCallback[T]) OnAfterUpdate(
	gongtreeStage *tree.StageStruct,
	stagedNode, frontNode *tree.Node) {

	FillUpFormFromGongstruct(
		instanceNodeCallback.Instance,
		instanceNodeCallback.probe,
	)
}
