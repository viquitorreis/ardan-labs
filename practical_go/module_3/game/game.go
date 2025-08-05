package main

import (
	"fmt"
	"slices"
)

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

	i.Move(10, 20)

	fmt.Println(i)

	p1 := Player{
		Name: "Victor",
	}

	fmt.Printf("p1: %+v\n", p1)
	fmt.Println("p1.X:", p1.X)
	p1.Move(100, 200)
	fmt.Println("p1 (move) \n", p1)

	fmt.Println(p1.Found("copper"))
	fmt.Println(p1.Found("copper"))
	fmt.Println(p1.Found("gold"))

	fmt.Println("keys:", p1.Keys)
}

type Item struct {
	X int
	Y int
}

type Player struct {
	Name string
	Item // campo implicito (COMPOSITION)
	Keys []string
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

// move i por delta x e delta y
/*
Value vs Pointer Receiver
- Geralmente usa VALUE semantics
- TENTA usar a mesma semantica em TODOS os metodos, se estiver usando value usar value,
	se estiver usando pointer semantics continua usando pointer nos outros métodos

- Quando DEVE usar pointer receiver:
	- Se tiver um campo de LOCK
	- Se precisar mutar a struct (alterar o valor nela)
	- Decoding / unmarshaling
*/
func (i *Item) Move(dx, dy int) {
	i.X += dx
	i.Y += dy
}

// Pointer Semantics - vamos mutar o player
func (p *Player) Found(key string) error {
	switch key {
	case "copper", "jade", "crystal":
		// ok
	default:
		return fmt.Errorf("unknown key: %q", key)
	}

	if !slices.Contains(p.Keys, key) {
		p.Keys = append(p.Keys, key)
	}

	return nil
}
