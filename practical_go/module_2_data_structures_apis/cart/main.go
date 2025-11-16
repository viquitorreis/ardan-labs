package main

import (
	"fmt"
	"sort"
)

func main() {
	cart := []string{"apple", "orange", "banana"}
	fmt.Println("len:", len(cart))
	fmt.Println("cart[1]", cart[1])

	// idx + value
	for i, c := range cart {
		fmt.Println(i, c)
	}

	cart = append(cart, "milk")
	fmt.Println(cart)

	// slicing operator
	fruit := cart[:3]
	fmt.Println("fruit:", fruit)
	fruit = append(fruit, "lemon")
	fmt.Println("fruit:", fruit)
	fmt.Println("cart:", cart) // printa o lemon tambem

	var s []int
	for i := range 10_000 {
		s = appendInt(s, i)
	}
	fmt.Println(s[:10])

	// Exercise: concat, without using a "for" loop
	out := concat([]string{"A", "B"}, []string{"C"})
	fmt.Println(out)

	values := []float64{3, 1, 2}
	fmt.Println(median(values)) // 2
	values = []float64{3, 1, 2, 4}
	fmt.Println(median(values)) // 2.5
	fmt.Println("values: ", values)

	players := []Player{
		{"Rick", 10_000},
		{"Morty", 11},
	}

	// adicionando um bonus
	// Value semantics "for" loop (NAO VAI DAR CERTO pois Go passa por valor)
	for _, p := range players {
		p.Score += 100
	}
	fmt.Println("players:", players)

	for i := range players {
		players[i].Score += 100
	}
	fmt.Println("players:", players)
}

type Player struct {
	Name  string
	Score int
}

func appendInt(slc []int, val int) []int {
	i := len(slc)

	// nao tem mais espaco no array por baixoo dos panos
	// precisa fazer o reallocate e copy
	if len(slc) == cap(slc) {
		size := 2*len(slc) + 1
		fmt.Println(cap(slc), "->", size)
		newSlice := make([]int, size)

		// copiando slice antigo para o novo slice
		copy(newSlice, slc)
		// slc vai apontar para um novo endereco de memoria
		slc = newSlice[:len(slc)]
	}

	// nao vai bugar pois ainda tem mais localizacao na capacidade do array
	slc = slc[:len(slc)+1]
	slc[i] = val // adicionando o novo valor no final do array
	return slc
}

func concat(src, incoming []string) []string {
	s := make([]string, len(src)+len(incoming))
	copy(s, src)                 // copiando src no inicio do array
	copy(s[len(src):], incoming) // copiando o resto no final do array
	return s
}

/*
	Median

- faz o sort dos valores
- se o valor for impar: retorna o middle
- retorna a média dos valores
*/
func median(values []float64) float64 {
	// copy para não fazer mutation no argumento de input
	vals := make([]float64, len(values))
	copy(vals, values)

	sort.Float64s(vals)
	i := len(vals) / 2

	if len(vals)%2 == 1 {
		return vals[i]
	}

	mid := (vals[i-1] + vals[i]) / 2
	return mid
}
