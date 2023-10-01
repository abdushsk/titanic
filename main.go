package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/gocarina/gocsv"
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

func readHumans(name string) []*Human {
	bs, err := os.ReadFile(name)
	checkErr(err)
	var hs []*Human
	err = gocsv.UnmarshalBytes(bs, &hs)
	checkErr(err)
	return hs
}

type Human struct {
	Id          int     `csv:"id"`
	Survived    int     `csv:"Survived,omitempty"`
	PassengerId int     `csv:"PassengerId"`
	Pclass      int     `csv:"Pclass"`
	Name        string  `csv:"Name"`
	Sex         string  `csv:"Sex"`
	Age         int     `csv:"Age"`
	SibSp       int     `csv:"SibSp"`
	Parch       int     `csv:"Parch"`
	Ticket      string  `csv:"Ticket"`
	Fare        float64 `csv:"Fare"`
	Embarked    string  `csv:"Embarked"`
}
