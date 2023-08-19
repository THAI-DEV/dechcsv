package dechcsv_test

import (
	"fmt"

	"github.com/THAI-DEV/dechcsv"
)

func ExampleSelectDataByRowNo() {
	list := dechcsv.ReadCsvFile("./input/test.csv", false)
	strList := dechcsv.SelectDataByRowNo(list, 1)
	str := strList[:]

	fmt.Println(str)

	// Output:
	// [col1 col2 col3]
}

func ExampleSelectDataByColNo() {
	list := dechcsv.ReadCsvFile("./input/test.csv", false)
	strList := dechcsv.SelectDataByColNo(list, 2)
	str := strList[:]

	fmt.Println(str)

	// Output:
	// [col2 2 B X2]
}

func ExampleSelectDataByRowColNo() {
	list := dechcsv.ReadCsvFile("./input/test.csv", false)
	str := dechcsv.SelectDataByRowColNo(list, 3, 3)

	fmt.Println(str)

	// Output:
	// c
}

func ExampleRemoveDataByRowNo() {
	list := dechcsv.ReadCsvFile("./input/test.csv", false)
	newList := dechcsv.RemoveDataByRowNo(list, 1)
	strList := dechcsv.SelectDataByRowNo(newList, 1)

	str := strList[:]

	fmt.Println(str)

	// Output:
	// [1 2 3]
}

func ExampleRemoveDataByColNo() {
	list := dechcsv.ReadCsvFile("./input/test.csv", false)
	newList := dechcsv.RemoveDataByColNo(list, 1)
	strList := dechcsv.SelectDataByRowNo(newList, 1)

	str := strList[:]

	fmt.Println(str)

	// Output:
	// [col2 col3]
}

func ExampleReplaceDataByRowColNo() {
	list := dechcsv.ReadCsvFile("./input/test.csv", false)
	newList := dechcsv.ReplaceDataByRowColNo(list, 1, 2, "Z", false)
	strList := dechcsv.SelectDataByRowNo(newList, 1)

	str := strList[:]

	fmt.Println(str)

	// Output:
	// [col1 Z col3]
}

func ExampleSelectRowNoByData() {
	list1 := dechcsv.ReadCsvFile("./input/test.csv", false)
	nList1 := dechcsv.SelectRowNoByData(list1, "col2", false)

	list2 := dechcsv.ReadCsvFile("./input/test.csv", false)
	nList2 := dechcsv.SelectRowNoByData(list2, "2", true)

	fmt.Println(nList1)
	fmt.Println(nList2)

	// Output:
	// [1]
	// [1 2 4]
}

func ExampleSelectColNoByData() {
	list := dechcsv.ReadCsvFile("./input/test.csv", false)
	row, col := dechcsv.DataInfo(list)

	fmt.Println(row, col)

	// Output:
	// 4 3

}
