package main

import (
	"errors"
	"math"
	"os"
	"strconv"
	"testing"
)

func TestSolution_getNumbers(t *testing.T) {
	input := "5\n10\n15\n-20\n25\n-20" // Example input data
	expected := 6

	// Save original stdin
	oldStdin := os.Stdin

	// Create a pipe to replace stdin
	r, w, _ := os.Pipe()
	os.Stdin = r

	// Write input to the pipe
	w.WriteString(input)
	w.Close()

	var sol Solution
	err := sol.getNumbers()
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	// Restore original stdin
	os.Stdin = oldStdin
	// sol.printSol(nil, false, false, false, false)
	if len(sol.numbers) != expected {
		t.Errorf("Expected %d numbers, got %d", expected, len(sol.numbers))
	}
	if sol.median != 7.5 {
		t.Errorf("Expected %.2g numbers, got %.2g", 7.5, sol.median)
	}
	if sol.mean != 2.5 {
		t.Errorf("Expected %.2g numbers, got %.2g", 2.5, sol.mean)
	}
	if math.Abs(sol.sd-17) <= 0.01 {
		t.Errorf("Expected %.2g numbers, got %.2g", 17., sol.sd)
	}
	if sol.mode != -20 {
		t.Errorf("Expected %d numbers, got %d", -20, sol.mode)
	}
}

func TestGetNumbers_EmptyInput(t *testing.T) {
	input := ""                           // Пустой ввод
	expectedErr := errors.New("zero len") // Ожидаемая ошибка

	// Создаем канал для стандартного ввода
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r

	// Пишем пустую строку в канал
	w.WriteString(input)
	w.Close()

	var sol Solution
	err := sol.getNumbers()

	// Восстанавливаем оригинальный стандартный ввод
	os.Stdin = oldStdin

	// Проверяем, что ошибка соответствует ожидаемой
	if err.Error() != expectedErr.Error() {
		t.Errorf("Expected error: %v, got: %v", expectedErr, err)
	}
}

func TestGetNumbers_InvalidInput(t *testing.T) {
	input := "5\n10\nabc\n20\n25\n"  // Некорректный ввод
	expectedErr := strconv.ErrSyntax // Ожидаемая ошибка

	// Создаем канал для стандартного ввода
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r

	// Пишем ввод с некорректным числом в канал
	w.WriteString(input)
	w.Close()

	var sol Solution
	err := sol.getNumbers()

	// Восстанавливаем оригинальный стандартный ввод
	os.Stdin = oldStdin

	// Проверяем, что ошибка соответствует ожидаемой
	if numErr, ok := err.(*strconv.NumError); !ok || numErr.Err != expectedErr {
		t.Errorf("Expected error: %v, got: %v", expectedErr, err)
	}
}

func TestGetNumbers_InvalidInput2(t *testing.T) {
	input := "abc\n5\n10\n20\n25\n"  // Некорректный ввод
	expectedErr := strconv.ErrSyntax // Ожидаемая ошибка

	// Создаем канал для стандартного ввода
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r

	// Пишем ввод с некорректным числом в канал
	w.WriteString(input)
	w.Close()

	var sol Solution
	err := sol.getNumbers()

	// Восстанавливаем оригинальный стандартный ввод
	os.Stdin = oldStdin

	// Проверяем, что ошибка соответствует ожидаемой
	if numErr, ok := err.(*strconv.NumError); !ok || numErr.Err != expectedErr {
		t.Errorf("Expected error: %v, got: %v", expectedErr, err)
	}
}

func TestGetNumbers_InvalidInput3(t *testing.T) {
	input := " 5\n10\n20\n25\n"      // Некорректный ввод
	expectedErr := strconv.ErrSyntax // Ожидаемая ошибка

	// Создаем канал для стандартного ввода
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r

	// Пишем ввод с некорректным числом в канал
	w.WriteString(input)
	w.Close()

	var sol Solution
	err := sol.getNumbers()

	// Восстанавливаем оригинальный стандартный ввод
	os.Stdin = oldStdin

	// Проверяем, что ошибка соответствует ожидаемой
	if numErr, ok := err.(*strconv.NumError); !ok || numErr.Err != expectedErr {
		t.Errorf("Expected error: %v, got: %v", expectedErr, err)
	}
}

func TestGetNumbers_InvalidInput4(t *testing.T) {
	input := "5 \n10\n20\n25\n"      // Некорректный ввод
	expectedErr := strconv.ErrSyntax // Ожидаемая ошибка

	// Создаем канал для стандартного ввода
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r

	// Пишем ввод с некорректным числом в канал
	w.WriteString(input)
	w.Close()

	var sol Solution
	err := sol.getNumbers()

	// Восстанавливаем оригинальный стандартный ввод
	os.Stdin = oldStdin

	// Проверяем, что ошибка соответствует ожидаемой
	if numErr, ok := err.(*strconv.NumError); !ok || numErr.Err != expectedErr {
		t.Errorf("Expected error: %v, got: %v", expectedErr, err)
	}
}

func TestSolution_getNumbers2(t *testing.T) {
	input := "5\n" // Example input data
	expected := 1
	expectedMedian := 5.
	expectedMean := 5.
	expectedSD := 0.
	expectedMode := 5

	// Save original stdin
	oldStdin := os.Stdin

	// Create a pipe to replace stdin
	r, w, _ := os.Pipe()
	os.Stdin = r

	// Write input to the pipe
	w.WriteString(input)
	w.Close()

	var sol Solution
	err := sol.getNumbers()
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	// Restore original stdin
	os.Stdin = oldStdin
	// sol.printSol(nil, false, false, false, false)
	if len(sol.numbers) != expected {
		t.Errorf("Expected %d numbers, got %d", expected, len(sol.numbers))
	}
	if sol.median != expectedMedian {
		t.Errorf("Expected median: %.2f, got: %.2f", expectedMedian, sol.median)
	}

	// Проверяем среднее
	if sol.mean != expectedMean {
		t.Errorf("Expected mean: %.2f, got: %.2f", expectedMean, sol.mean)
	}

	// Проверяем стандартное отклонение
	if sol.sd != expectedSD {
		t.Errorf("Expected standard deviation: %.2f, got: %.2f", expectedSD, sol.sd)
	}

	// Проверяем моду
	if sol.mode != expectedMode {
		t.Errorf("Expected mode: %d, got: %d", expectedMode, sol.mode)
	}
}
