package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func main() {
	genDay()
}

var template = `package main

func main() {

}

func partOne() {

}

func partTwo() {

}`

func genDay() {
	day := time.Now().UTC().Day()
	if err := os.Mkdir(strconv.Itoa(day), os.ModePerm); err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(fmt.Sprintf("%d/main.go", day), []byte(template), os.ModePerm); err != nil {
		panic(err)
	}
}
