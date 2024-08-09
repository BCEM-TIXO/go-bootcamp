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

func CheckTime(old, new Cake) {
	if old.Time != new.Time {
		fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", old.Name, old.Time, new.Time)
	}
}

func CheckIngridients(old, new Cake) {
	var old_map, new_map map[string]Ingredient
	old_map = make(map[string]Ingredient)
	new_map = make(map[string]Ingredient)

	for _, Ingredient := range old.Ingredients {
		old_map[Ingredient.Ingredient_name] = Ingredient
	}

	for _, Ingredient := range new.Ingredients {
		new_map[Ingredient.Ingredient_name] = Ingredient
	}

	for Ingredient_name, old_ing := range old_map {
		new_ing := new_map[Ingredient_name]
		if new_ing.Ingredient_name != old_ing.Ingredient_name {
			fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", old_ing.Ingredient_name, old.Name)
		} else {
			if new_ing.Ingredient_count != old_ing.Ingredient_count {
				fmt.Printf("CHANGED ingredient count \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", new_ing.Ingredient_name, old.Name, old_ing.Ingredient_count, new_ing.Ingredient_count)
			}
			if old_ing.Ingredient_unit != "" && new_ing.Ingredient_unit != "" && new_ing.Ingredient_unit != old_ing.Ingredient_unit {
				fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", old_ing.Ingredient_name, old.Name, old_ing.Ingredient_unit, new_ing.Ingredient_unit)
			} else if new_ing.Ingredient_unit != "" {
				fmt.Printf("ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", old_ing.Ingredient_unit, old_ing.Ingredient_name, old.Name)
			} else if new_ing.Ingredient_unit == "" && old_ing.Ingredient_unit != "" {
				fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", old_ing.Ingredient_unit, old_ing.Ingredient_name, old.Name)
			}
		}
	}

	for Ingredient_name := range new_map {
		if old_map[Ingredient_name].Ingredient_name != Ingredient_name {
			fmt.Printf("ADDED ingridient \"%s\" for cake \"%s\"\n", Ingredient_name, old.Name)
		}
	}
}

func CheckCakes(old, new []Cake) {
	var old_map, new_map map[string]Cake
	old_map = make(map[string]Cake)
	new_map = make(map[string]Cake)

	for _, cake1 := range old {
		old_map[cake1.Name] = cake1
	}

	for _, cake2 := range new {
		new_map[cake2.Name] = cake2
	}

	for cake_name := range old_map {
		if new_map[cake_name].Name != cake_name {
			fmt.Printf("REMOVED cake \"%s\"\n", cake_name)
			continue
		}
		CheckTime(old_map[cake_name], new_map[cake_name])
		CheckIngridients(old_map[cake_name], new_map[cake_name])
	}

	for cake_name := range new_map {
		if old_map[cake_name].Name != cake_name {
			fmt.Printf("ADDED cake \"%s\"\n", cake_name)
		}
	}
}

func CompareDB(old, new []Cake) {
	CheckCakes(old, new)
}

func main() {
	oldDBfilePathPtr := flag.String("old", "", "path to old database file")
	newDBfilePathPtr := flag.String("new", "", "path to new database file")

	flag.Parse()
	old, err := getDB(*oldDBfilePathPtr)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	new, err := getDB(*newDBfilePathPtr)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	CompareDB(old, new)
}
