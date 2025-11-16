package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	fmt.Println(Relu(7))
	fmt.Println(Relu(-1))
	fmt.Println(Relu(1.2))
	fmt.Println(Relu(time.February))

	m, err := NewMatrix[float64](10, 3)
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("m: ", m)

	fmt.Println(m.At(3, 2))

	fmt.Println(Max([]int{3, 1, 2}))     // 3 <nil>
	fmt.Println(Max([]float64{3, 1, 2})) // 3 <nil>
	fmt.Println(Max[int](nil))           // 0 <nil>
}

func Max[T int | float64](n []T) (T, error) {
	if len(n) == 0 {
		return 0, errors.New("Max of empty slice")
	}

	var max T = n[0]
	for _, num := range n {
		if num > max {
			max = num
		}
	}

	return max, nil
}

// Number vai ser um type constraint
// podemos usar uma interface para combinar um conjunto de tipos também
type Number interface {
	~int | ~float64
}

// T é uma CONSTRAINT DE TIPO "type constraint" e não um novo tipo
// essa função pode receber apenas um inteiro ou um float
// vai permitir apenas int e float -> func Relu[T int | float64](i T) T {
// aqui, qualquer tipo mesmo que por baixo dos panos (underlying) seja int ou float vai ser aceito
// func Relu[T ~int | ~float64](i T) T {
func Relu[T Number](i T) T {
	if i < 0 {
		return 0
	}

	return i
}

type Matrix[T Number] struct {
	Rows int
	Cols int

	data []T
}

func (m *Matrix[T]) At(row, col int) T {
	i := (row * m.Cols) + col
	return m.data[i]
}

func NewMatrix[T Number](rows, cols int) (*Matrix[T], error) {
	if rows <= 0 || cols <= 0 {
		return nil, fmt.Errorf("bad dimensions: %d/%d", rows, cols)
	}

	m := Matrix[T]{
		Rows: rows,
		Cols: cols,
		data: make([]T, rows*cols),
	}

	return &m, nil
}
