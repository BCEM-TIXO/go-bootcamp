package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

func main() {
	// Создаем новое изображение размером 300x300 пикселей
	img := image.NewRGBA(image.Rect(0, 0, 300, 300))

	// Заполняем изображение цветом фона
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 255, 255, 255}}, image.ZP, draw.Src)

	// Радиусы кругов мишени
	radii := []int{100, 80, 60, 40, 20}
	// Цвета кругов мишени
	colors := []color.RGBA{
		{255, 0, 0, 255},     // красный
		{255, 255, 255, 255}, // белый
		{255, 0, 0, 255},     // красный
		{255, 255, 255, 255}, // белый
		{255, 0, 0, 255},     // красный
	}

	// Центр мишени
	x, y := 150, 150

	// Рисуем круги мишени
	for i, radius := range radii {
		drawCircle(img, x, y, radius, colors[i])
	}

	// Создаем файл для записи
	file, err := os.Create("amazing_logo.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Записываем изображение в файл в формате PNG
	if err := png.Encode(file, img); err != nil {
		panic(err)
	}
}

// Функция для рисования круга на изображении
func drawCircle(img *image.RGBA, x, y, r int, c color.RGBA) {
	for i := x - r; i <= x+r; i++ {
		for j := y - r; j <= y+r; j++ {
			if math.Pow(float64(i-x), 2)+math.Pow(float64(j-y), 2) <= math.Pow(float64(r), 2) {
				img.Set(i, j, c)
			}
		}
	}
}
