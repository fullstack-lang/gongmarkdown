// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	AnotherDummyDataAPIs []*AnotherDummyDataAPI

	CellAPIs []*CellAPI

	DummyDataAPIs []*DummyDataAPI

	ElementAPIs []*ElementAPI

	MarkdownContentAPIs []*MarkdownContentAPI

	RowAPIs []*RowAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, anotherdummydataDB := range backRepo.BackRepoAnotherDummyData.Map_AnotherDummyDataDBID_AnotherDummyDataDB {

		var anotherdummydataAPI AnotherDummyDataAPI
		anotherdummydataAPI.ID = anotherdummydataDB.ID
		anotherdummydataAPI.AnotherDummyDataPointersEncoding = anotherdummydataDB.AnotherDummyDataPointersEncoding
		anotherdummydataDB.CopyBasicFieldsToAnotherDummyData_WOP(&anotherdummydataAPI.AnotherDummyData_WOP)

		backRepoData.AnotherDummyDataAPIs = append(backRepoData.AnotherDummyDataAPIs, &anotherdummydataAPI)
	}

	for _, cellDB := range backRepo.BackRepoCell.Map_CellDBID_CellDB {

		var cellAPI CellAPI
		cellAPI.ID = cellDB.ID
		cellAPI.CellPointersEncoding = cellDB.CellPointersEncoding
		cellDB.CopyBasicFieldsToCell_WOP(&cellAPI.Cell_WOP)

		backRepoData.CellAPIs = append(backRepoData.CellAPIs, &cellAPI)
	}

	for _, dummydataDB := range backRepo.BackRepoDummyData.Map_DummyDataDBID_DummyDataDB {

		var dummydataAPI DummyDataAPI
		dummydataAPI.ID = dummydataDB.ID
		dummydataAPI.DummyDataPointersEncoding = dummydataDB.DummyDataPointersEncoding
		dummydataDB.CopyBasicFieldsToDummyData_WOP(&dummydataAPI.DummyData_WOP)

		backRepoData.DummyDataAPIs = append(backRepoData.DummyDataAPIs, &dummydataAPI)
	}

	for _, elementDB := range backRepo.BackRepoElement.Map_ElementDBID_ElementDB {

		var elementAPI ElementAPI
		elementAPI.ID = elementDB.ID
		elementAPI.ElementPointersEncoding = elementDB.ElementPointersEncoding
		elementDB.CopyBasicFieldsToElement_WOP(&elementAPI.Element_WOP)

		backRepoData.ElementAPIs = append(backRepoData.ElementAPIs, &elementAPI)
	}

	for _, markdowncontentDB := range backRepo.BackRepoMarkdownContent.Map_MarkdownContentDBID_MarkdownContentDB {

		var markdowncontentAPI MarkdownContentAPI
		markdowncontentAPI.ID = markdowncontentDB.ID
		markdowncontentAPI.MarkdownContentPointersEncoding = markdowncontentDB.MarkdownContentPointersEncoding
		markdowncontentDB.CopyBasicFieldsToMarkdownContent_WOP(&markdowncontentAPI.MarkdownContent_WOP)

		backRepoData.MarkdownContentAPIs = append(backRepoData.MarkdownContentAPIs, &markdowncontentAPI)
	}

	for _, rowDB := range backRepo.BackRepoRow.Map_RowDBID_RowDB {

		var rowAPI RowAPI
		rowAPI.ID = rowDB.ID
		rowAPI.RowPointersEncoding = rowDB.RowPointersEncoding
		rowDB.CopyBasicFieldsToRow_WOP(&rowAPI.Row_WOP)

		backRepoData.RowAPIs = append(backRepoData.RowAPIs, &rowAPI)
	}

}
