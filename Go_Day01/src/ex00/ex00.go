package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type DBReader interface {
	ReadDB() ([]Cake, error)
}

type Cakes struct {
	XMLName xml.Name `xml:"recipes"`
	Cakes   []Cake   `json:"cake" xml:"cake"`
}
type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}
type Ingredient struct {
	Ingredient_name  string `json:"ingredient_name" xml:"itemname"`
	Ingredient_count string `json:"ingredient_count" xml:"itemcount"`
	Ingredient_unit  string `json:"ingredient_unit" xml:"itemunit"`
}

type JSONReader struct {
	file_path string
}
type XMLReader struct {
	file_path string
}

func getFile(file_path string) ([]byte, error) {
	file, err := os.Open(file_path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	byteValue, err := io.ReadAll(file)
	return byteValue, err
}

func (r JSONReader) ReadDB() ([]Cake, error) {
	byteValue, err := getFile(r.file_path)
	var cakes Cakes
	json.Unmarshal(byteValue, &cakes)
	return cakes.Cakes, err
}

func (r XMLReader) ReadDB() ([]Cake, error) {
	byteValue, err := getFile(r.file_path)
	var cakes Cakes
	xml.Unmarshal(byteValue, &cakes)
	return cakes.Cakes, err
}

func getDB(file_path string) ([]Cake, error) {
	ext := filepath.Ext(file_path)
	var reader DBReader
	var err error
	if ext == ".json" {
		reader = JSONReader{file_path}
	} else if ext == ".xml" {
		reader = XMLReader{file_path}
	} else {
		err = errors.New("Unsupported file ext: " + file_path)
		return nil, err
	}
	return reader.ReadDB()
}
func main() {
	filePathPtr := flag.String("f", "", "path to database file")

	flag.Parse()
	var err error

	cakes, err := getDB(*filePathPtr)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	var bytes []byte
	if strings.HasSuffix(*filePathPtr, ".json") {
		bytes, _ = xml.MarshalIndent(cakes, "", "    ")
	} else if strings.HasSuffix(*filePathPtr, ".xml") {
		bytes, _ = json.MarshalIndent(cakes, "", "    ")
	}

	fmt.Println(string(bytes))
}
