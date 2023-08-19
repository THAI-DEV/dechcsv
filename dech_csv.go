package dechcsv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func WriteNewCsvFile(data [][]string, outputFile string) {
	f, err := os.Create(outputFile)

	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	err = w.WriteAll(data) // calls Flush internally

	if err != nil {
		log.Fatal(err)
	}
}

func WriteAppendCsvFile(data [][]string, outputFile string) {
	f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	err = w.WriteAll(data) // calls Flush internally

	if err != nil {
		log.Fatal(err)
	}
}

func ReadCsvFile(inputFile string, isSkipFirstRow bool) [][]string {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if isSkipFirstRow {
		row1, err := bufio.NewReader(f).ReadSlice('\n')
		if err != nil {
			fmt.Println(err)
		}
		_, err = f.Seek(int64(len(row1)), io.SeekStart)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Read remaining rows
	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		log.Println(err)
	}

	return rows
}

func ShowData(data [][]string) {
	for _, rowVal := range data {
		lenData := len(rowVal)
		for i := 0; i < lenData; i++ {
			fmt.Print(rowVal[i])
			if i < lenData-1 {
				fmt.Print(",")
			}

		}

		fmt.Println()
	}
}

func RemoveDataByRowNo(data [][]string, rowNo int) [][]string {
	result := [][]string{}
	for i, rowVal := range data {
		if i+1 != rowNo {
			result = append(result, rowVal)
		}
	}

	return result
}

func RemoveDataByColNo(data [][]string, colNo int) [][]string {
	result := [][]string{}

	for _, rowVal := range data {
		colData := []string{}
		for i := 0; i < len(rowVal); i++ {

			if i+1 != colNo {
				colData = append(colData, rowVal[i])
			}

		}

		result = append(result, colData)
	}

	return result
}

func SelectDataByRowNo(data [][]string, rowNo int) []string {
	for i, rowVal := range data {
		if i+1 == rowNo {
			return rowVal
		}
	}

	return []string{}
}

func SelectDataByColNo(data [][]string, colNo int) []string {
	result := []string{}

	for _, rowVal := range data {
		for i := 0; i < len(rowVal); i++ {
			if i+1 == colNo {
				result = append(result, rowVal[i])
			}
		}
	}

	return result
}

func SelectDataByRowColNo(data [][]string, rowNo int, colNo int) string {
	r, c := DataInfo(data)

	if (rowNo < 1) || (colNo < 1) || (rowNo > r) || (colNo > c) {
		log.Println("Data is not in range")
		return ""
	}

	return data[rowNo-1][colNo-1]
}

func ReplaceDataByRowColNo(data [][]string, rowNo int, colNo int, replaceData string, isReplaceOldData bool) [][]string {
	r, c := DataInfo(data)

	if (rowNo < 1) || (colNo < 1) || (rowNo > r) || (colNo > c) {
		log.Println("Data is not in range")
		return nil
	}

	result := [][]string{}

	if isReplaceOldData {
		result = data
	} else {
		result = CloneNewData(data)
	}

	result[rowNo-1][colNo-1] = replaceData

	return result
}

func CloneNewData(data [][]string) [][]string {
	result := [][]string{}

	for i := 0; i < len(data); i++ {
		rowData := data[i]
		r := []string{}
		for j := 0; j < len(rowData); j++ {
			r = append(r, rowData[j])
		}
		result = append(result, r)
	}

	return result
}

func DataInfo(data [][]string) (row int, col int) {
	row = 0
	col = 0

	if data == nil {
		return row, col
	}

	row = len(data)
	if row > 0 {
		col = len(data[0])
	}

	return row, col
}

func SelectRowNoByData(data [][]string, chkData string, isContain bool) []int {
	result := []int{}

	row, _ := DataInfo(data)
	for i := 0; i < row; i++ {
		d := SelectDataByRowNo(data, i+1)
		if isFoundData(d, chkData, isContain) {
			result = append(result, i+1)
		}

	}

	return result
}

func SelectColNoByData(data [][]string, chkData string, isContain bool) []int {
	result := []int{}

	_, col := DataInfo(data)
	for i := 0; i < col; i++ {
		d := SelectDataByColNo(data, i+1)
		if isFoundData(d, chkData, isContain) {
			result = append(result, i+1)
		}

	}

	return result
}

func isFoundData(data []string, chkData string, isContain bool) bool {
	for _, v := range data {
		condition := v == chkData
		if isContain {
			condition = strings.Contains(v, chkData)
		}

		if condition {
			return true
		}
	}

	return false
}
