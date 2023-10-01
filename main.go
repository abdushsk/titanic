package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

func AddId() {
	f, err := os.Open("assets/survived.csv")
	checkErr(err)
	csvR := csv.NewReader(f)
	fields, err := csvR.ReadAll()
	checkErr(err)
	var bts bytes.Buffer
	wr := csv.NewWriter(&bts)

	for i, v := range fields {
		if i == 0 {
			v = append([]string{"id"}, v...)
		} else {
			v = append([]string{fmt.Sprintf("%d", i)}, v...)
		}
		wr.Write(v)
	}
	wr.Flush()
	ioutil.WriteFile("assets/surv.csv", bts.Bytes(), fs.ModePerm)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	AddId()
}
