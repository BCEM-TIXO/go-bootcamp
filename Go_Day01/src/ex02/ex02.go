package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func compareDumps(fileName1, fileName2 string) {
	data1 := getFileMap(fileName1)
	data2 := getFileMap(fileName2)

	for key := range data2 {
		if _, ok := data1[key]; !ok {
			fmt.Println("ADDED:", key)
		}
	}

	for key := range data1 {
		if _, ok := data2[key]; !ok {
			fmt.Println("REMOVED:", key)
		}
	}

}

func getFileMap(fileName string) map[string]bool {
	var file *os.File
	var err error

	file, err = os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		if errClose := f.Close(); errClose != nil {
			log.Fatal(errClose)
		}
	}(file)

	lines := make(map[string]bool)
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		line := reader.Text()
		lines[line] = true
	}
	return lines
}

type Flags struct {
	oldF, newF *string
}

var fl Flags

func init() {
	fl.oldF = flag.String("old", "", "original db")
	fl.newF = flag.String("new", "", "new db")
}

func main() {
	flag.Parse()
	if *fl.oldF == "" || *fl.newF == "" {
		fmt.Println("ERROR:\n\t--old [file name]\n\t--new [file name]")
		return
	}

	compareDumps(*fl.oldF, *fl.newF)
}
