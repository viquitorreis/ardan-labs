package main

import "fmt"

func main() {
	var i Item
	// # -> mostra o tipo todo e os valores
	fmt.Printf("i: %#v\n", i)
	i = Item{10, 20}
	fmt.Printf("i: %#v\n", i)

	i = Item{
		X: 22,
		Y: 11,
	}

	fmt.Println(NewItem(10, 20))
	fmt.Println(NewItem(10, 200))
}

type Item struct {
	X int
	Y int
}

/*
-- Constructor
func NewItem(x, y int) (Item)
func NewItem(x, y int) (*Item)
func NewItem(x, y int) (Item, error)
func NewItem(x, y int) (*Item, error)

Se:
	- Value Semantics (pass by value): cada um vai ter seu próprio valor
	- POinter Semantics: cada um compartilha a mesma cópia (geralmente na HEAP e precisa de um Lock - geralmente)
*/

const (
	maxX = 600
	maxY = 400
)

func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		// value semantics
		//return Item{}, fmt.Errorf("%d/%d out of bound %d/%d", x, y, maxX, maxY)
		return nil, fmt.Errorf("%d/%d out of bound %d/%d", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}

	// go compiler faz um "escape analysis" e vai alocar i na HEAP (para ver chamar com go - build -gcflags=-m)
	// passar as coisas por valor (!= de pointer semantics) em GO, geralmente vai te dar uma performance melhor
	// pois vai aliviar o garbage collector
	return &i, nil
}
