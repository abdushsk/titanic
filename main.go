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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	humans := readHumans("assets/values.csv")

	for _, v := range humans {
		if v.Age < 15 {
			v.Survived = 1
		}
		if v.Sex == "female" && v.Age < 60 {
			v.Survived = 1
		}
		if v.Age > 70 {
			v.Survived = 0
		}
	}

	perc := getSurvPercentage(humans)

	fmt.Println(fmt.Sprintf("Correct : %.2f", perc))
}

func readHumans(name string) []*Human {
	bs, err := os.ReadFile(name)
	checkErr(err)
	var hs []*Human
	err = gocsv.UnmarshalBytes(bs, &hs)
	checkErr(err)
	return hs
}

func getSurvPercentage(humans []*Human) float64 {
	surv := readHumans("assets/surv.csv")

	correct := 0.0

	for _, c := range surv {
		for _, n := range humans {
			if c.Id != n.Id {
				continue
			}
			if c.Survived == n.Survived {
				correct++
				break
			}
		}
	}

	return (correct / float64(len(surv))) * 100
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
