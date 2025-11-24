package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Matrix struct {
	Rows int
	Cols int
	Data [][]float64
}

func NewMatrix(rows, cols int) Matrix {
	data := make([][]float64, rows)
	for i := range data {
		data[i] = make([]float64, cols)
	}
	return Matrix{Rows: rows, Cols: cols, Data: data}
}

func (m Matrix) PrintMatrix() {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			fmt.Printf("%8.2f ", m.Data[i][j])
		}
		fmt.Println()
	}
}

func Add(m1, m2 Matrix) (Matrix, error) {
	if m1.Rows != m2.Rows || m1.Cols != m2.Cols {
		return NewMatrix(0, 0), errors.New("размеры матриц не совпадают")
	}
	result := NewMatrix(m1.Rows, m1.Cols)
	for i := 0; i < m1.Rows; i++ {
		for j := 0; j < m1.Cols; j++ {
			result.Data[i][j] = m1.Data[i][j] + m2.Data[i][j]
		}
	}
	return result, nil
}

func ScalarMultiply(m Matrix, scalar float64) Matrix {
	result := NewMatrix(m.Rows, m.Cols)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			result.Data[i][j] = m.Data[i][j] * scalar
		}
	}
	return result
}

func Multiply(m1, m2 Matrix) (Matrix, error) {
	if m1.Cols != m2.Rows {
		return NewMatrix(0, 0), errors.New("несовместимые размеры для умножения")
	}
	result := NewMatrix(m1.Rows, m2.Cols)
	for i := 0; i < m1.Rows; i++ {
		for j := 0; j < m2.Cols; j++ {
			sum := 0.0
			for k := 0; k < m1.Cols; k++ {
				sum += m1.Data[i][k] * m2.Data[k][j]
			}
			result.Data[i][j] = sum
		}
	}
	return result, nil
}

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func readLine(prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func parseInt(prompt string) (int, error) {
	s := readLine(prompt)
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("требуется целое число")
	}
	return val, nil
}

func parseFloat(prompt string) (float64, error) {
	s := readLine(prompt)
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, errors.New("требуется число")
	}
	return val, nil
}

func readMatrix(description string) (Matrix, error) {
	fmt.Println(description)
	var size int
	for {
		s, err := parseInt("Размер матрицы (2 или 3): ")
		if err == nil && (s == 2 || s == 3) {
			size = s
			break
		}
		fmt.Println("Ошибка: " + err.Error() + ". Введите 2 или 3.")
	}

	matrix := NewMatrix(size, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for {
				val, err := parseFloat(fmt.Sprintf("Элемент [%d][%d]: ", i, j))
				if err == nil {
					matrix.Data[i][j] = val
					break
				}
				fmt.Println("Ошибка: " + err.Error())
			}
		}
	}
	return matrix, nil
}

func handleAddition() {
	m1, err1 := readMatrix("--- Первая матрица ---")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	m2, err2 := readMatrix("--- Вторая матрица ---")
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	result, err := Add(m1, m2)
	if err != nil {
		fmt.Println("Ошибка сложения:", err)
	} else {
		fmt.Println("\nРезультат сложения:")
		result.PrintMatrix()
	}
}

func handleScalarMultiplication() {
	m1, err := readMatrix("--- Матрица для умножения на скаляр ---")
	if err != nil {
		fmt.Println(err)
		return
	}

	scalar, err := parseFloat("Введите скаляр: ")
	if err != nil {
		fmt.Println("Ошибка ввода скаляра:", err)
		return
	}

	result := ScalarMultiply(m1, scalar)
	fmt.Println("\nРезультат умножения на скаляр:")
	result.PrintMatrix()
}

func handleMatrixMultiplication() {
	m1, err1 := readMatrix("--- Первая матрица ---")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	m2, err2 := readMatrix("--- Вторая матрица ---")
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	if m1.Cols != m2.Rows {
		fmt.Println("Ошибка: размеры матриц несовместимы для умножения.")
		return
	}

	result, err := Multiply(m1, m2)
	if err != nil {
		fmt.Println("Ошибка умножения:", err)
	} else {
		fmt.Println("\nРезультат умножения матриц:")
		result.PrintMatrix()
	}
}

func main() {
	for {
		fmt.Println("\n--- Калькулятор матриц ---")
		fmt.Println("1. Сложение")
		fmt.Println("2. Скалярное умножение")
		fmt.Println("3. Матричное умножение")
		fmt.Println("4. Выход")
		choice, err := parseInt("Ваш выбор: ")

		if err != nil {
			fmt.Println("Ошибка: " + err.Error())
			continue
		}

		switch choice {
		case 1:
			handleAddition()
		case 2:
			handleScalarMultiplication()
		case 3:
			handleMatrixMultiplication()
		case 4:
			fmt.Println("До свидания!")
			os.Exit(0)
		default:
			fmt.Println("Неверный выбор. Введите 1-4.")
		}
	}
}
