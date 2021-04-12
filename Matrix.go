package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	//filePath := "./Resources/matrix_example.csv"
	//ReadCsvFile(filePath)
	//NewMatrix(filePath)
	NewMatrix(5, 5)
}

type IMatrix interface {
	division() float64
	sum() float64
	multiplier(m1 Matrix, m2 Matrix) Matrix
}


type Matrix struct {
	body [][]float64
}

func NewMatrix( columns int, rows int) *Matrix {
	m := &Matrix{}
	fmt.Println("read csv reader return")
	var mtx = [][]int{}
	for i := 0; i < columns; i ++{
		mtx = append(mtx, make([]int, rows))
	}
	fmt.Println(mtx)

	return m
}



func ReadCsvFile(filePath string) [][]string {
	// load csv file
	f, _ := os.Open(filePath)
	// create reader
	r := csv.NewReader(f)
	list, err := r.ReadAll()
	if err == io.EOF || err != nil {
		panic(err) // if error stop here,  i dont know whats this does, but cause me panic
	}
	return list
	//var body [][]string
	//for {
	//	record, err := r.Read()
	//	if err == io.EOF { // if error stop here
	//		break
	//	}
	//	if err != nil {
	//		panic(err) // i dont know whats this does, but cause me panic
	//	}
	//	fmt.Println(record) // print the array line
	//	fmt.Println(record[0]) // take the item on the index
	//	for index, value := range record {
	//		fmt.Printf("record: %v - %v\n", index, value)
	//	}
	//}
}
