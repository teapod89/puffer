package report

import (
	"log"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func Out(f string, infos []map[string]string, fnCount, dFnCount int) error {
	fh := excelize.NewFile()
	// シートを新規作成
	index := fh.NewSheet("Sheet1")

	//透明の罫線は非表示に設定する
	fh.SetSheetViewOptions("Sheet1", 0, excelize.ShowGridLines(false))

	// シートに値を書き込む
	fh.SetCellValue("Sheet1", "A1", "Report")

	//先頭行の見出し設定
	fh.SetCellStyle("Sheet1", "A2", "G2", makeTitleStyle((fh)))

	//ヘッダー行の出力
	fh.SetCellValue("Sheet1", "A2", "no")
	fh.SetCellValue("Sheet1", "B2", "refer file")
	fh.SetCellValue("Sheet1", "C2", "refer hash")
	fh.SetCellValue("Sheet1", "D2", "duplicate file")
	fh.SetCellValue("Sheet1", "E2", "duplicate hash")
	fh.SetCellValue("Sheet1", "F2", "stat state")
	fh.SetCellValue("Sheet1", "G2", "rm state")
	//列幅の設定
	widthOffset := 4
	fh.SetColWidth("Sheet1", "C", "C", float64(70+widthOffset))
	fh.SetColWidth("Sheet1", "E", "E", float64(70+widthOffset))
	fh.SetColWidth("Sheet1", "B", "B", float64(fnCount+widthOffset))
	fh.SetColWidth("Sheet1", "D", "D", float64(dFnCount+widthOffset))

	offset := 3
	for c, info := range infos {
		c = c + offset
		axisIndex := strconv.Itoa(c)
		fh.SetCellValue("Sheet1", "A"+axisIndex, strconv.Itoa(c-2))
		fh.SetCellValue("Sheet1", "B"+axisIndex, info["filename"])
		fh.SetCellValue("Sheet1", "C"+axisIndex, info["hash"])
		fh.SetCellValue("Sheet1", "D"+axisIndex, info["duplicate_filename"])
		fh.SetCellValue("Sheet1", "E"+axisIndex, info["duplicate_hash"])
		fh.AddComment("Sheet1", "B"+axisIndex, `{"author":"system: ","text":"dir:`+info["directory"]+`"}`)
		fh.AddComment("Sheet1", "D"+axisIndex, `{"author":"system: ","text":"dir:`+info["duplicate_directory"]+`"}`)
		_, isStat := info["filestat"]
		if isStat {
			fh.SetCellValue("Sheet1", "F"+axisIndex, info["filestat"])
		} else {
			fh.SetCellValue("Sheet1", "F"+axisIndex, "-")
		}
		_, isRemove := info["fileremove"]
		if isRemove {
			fh.SetCellValue("Sheet1", "G"+axisIndex, info["fileremove"])
		} else {
			fh.SetCellValue("Sheet1", "G"+axisIndex, "-")
		}
	}
	fh.SetCellStyle("Sheet1", "A"+strconv.Itoa(offset), "G"+strconv.Itoa(len(infos)+2), makeBodyStyle(fh))

	// シート1をアクティブにする。
	fh.SetActiveSheet(index)

	if err := fh.SaveAs(f); err != nil {
		log.Fatalln("failed to output report.")
	}

	return nil
}

func makeTitleStyle(fh *excelize.File) int {
	template := `{
		"fill": {
			"type": "pattern",
			"color": [
				"#D3D3D3"
			],
			"pattern": 1
		},
		"border": [
			{
				"type": "bottom",
				"color": "000000",
				"style": 1
			},
			{
				"type": "top",
				"color": "000000",
				"style": 1
			},
			{
				"type": "left",
				"color": "000000",
				"style": 1
			},
			{
				"type": "right",
				"color": "000000",
				"style": 1
			}
		]
	}`

	id, err := fh.NewStyle(template)
	if err != nil {
		log.Fatalln("failed to cell style setting.")
	}
	return id
}

func makeBodyStyle(fh *excelize.File) int {
	template := `{
		"border": [
			{
				"type": "bottom",
				"color": "000000",
				"style": 1
			},
			{
				"type": "top",
				"color": "000000",
				"style": 1
			},
			{
				"type": "left",
				"color": "000000",
				"style": 1
			},
			{
				"type": "right",
				"color": "000000",
				"style": 1
			}
		]
	}`

	id, err := fh.NewStyle(template)
	if err != nil {
		log.Fatalln("failed to cell style setting.")
	}
	return id
}
