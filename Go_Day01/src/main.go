package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	// "sort"
)

// Cake структура для представления информации о куличе
type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

// Ingredient структура для представления информации об ингредиенте
type Ingredient struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit" xml:"itemunit"`
}

// DBReader интерфейс для чтения базы данных
type DBReader interface {
	Read(filename string) ([]Cake, error)
}

// XMLReader реализация интерфейса DBReader для чтения XML файлов
type XMLReader struct{}

// Read читает XML файл и возвращает список куличей
func (r XMLReader) Read(filename string) ([]Cake, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var recipes struct {
		Cakes []Cake `xml:"cake"`
	}
	if err := xml.Unmarshal(data, &recipes); err != nil {
		return nil, err
	}

	return recipes.Cakes, nil
}

// JSONReader реализация интерфейса DBReader для чтения JSON файлов
type JSONReader struct{}

// Read читает JSON файл и возвращает список куличей
func (r JSONReader) Read(filename string) ([]Cake, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cakes []Cake
	if err := json.NewDecoder(file).Decode(&cakes); err != nil {
		return nil, err
	}

	return cakes, nil
}

// Read читает XML файл и возвращает список куличей
func CompareDB(oldDB, newDB []Cake) {
	oldMap := make(map[string]Cake)
	for _, cake := range oldDB {
		oldMap[cake.Name] = cake
	}

	for _, newCake := range newDB {
		oldCake, exists := oldMap[newCake.Name]
		if !exists {
			fmt.Printf("ADDED cake \"%s\"\n", newCake.Name)
			continue
		}

		if newCake.Time != oldCake.Time {
			fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", newCake.Name, newCake.Time, oldCake.Time)
		}

		oldIngredients := make(map[string]Ingredient)
		for _, ingredient := range oldCake.Ingredients {
			oldIngredients[ingredient.Name] = ingredient
		}

		for _, newIngredient := range newCake.Ingredients {
			oldIngredient, exists := oldIngredients[newIngredient.Name]
			if !exists {
				fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", newIngredient.Name, newCake.Name)
				continue
			}

			if newIngredient.Count != oldIngredient.Count {
				fmt.Printf("CHANGED count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", newIngredient.Name, newCake.Name, newIngredient.Count, oldIngredient.Count)
			}

			if newIngredient.Unit != oldIngredient.Unit {
				fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", newIngredient.Name, newCake.Name, newIngredient.Unit, oldIngredient.Unit)
				// } else if newIngredient.Unit == nil && oldIngredient.Unit != nil {
				// 	fmt.Printf("REMOVED unit for ingredient \"%s\" for cake \"%s\"\n", newIngredient.Name, newCake.Name)
				// } else if newIngredient.Unit != nil && oldIngredient.Unit == nil {
				// 	fmt.Printf("ADDED unit for ingredient \"%s\" for cake \"%s\"\n", newIngredient.Name, newCake.Name)
			}

			delete(oldIngredients, newIngredient.Name)
		}

		for name := range oldIngredients {
			fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", name, newCake.Name)
		}
	}

	for _, oldCake := range oldDB {
		_, exists := oldMap[oldCake.Name]
		if !exists {
			fmt.Printf("REMOVED cake \"%s\"\n", oldCake.Name)
		}
	}
}

func main() {
	xmlReader := XMLReader{}
	jsonReader := JSONReader{}

	oldDB, err := xmlReader.Read("xml.xml")
	if err != nil {
		fmt.Println("Error reading old database:", err)
		return
	}

	fmt.Println(oldDB)
	newDB, err := jsonReader.Read("json.json")
	if err != nil {
		fmt.Println("Error reading new database:", err)
		return
	}

	CompareDB(oldDB, newDB)
}
