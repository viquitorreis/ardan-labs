package main

import (
	"bufio"
	"fmt"
	"log"
	"maps"
	"os"
	"regexp"
	"slices"
	"sort"
	"strings"
)

// quais as N palavras mais comuns em sherlock.txt

var wordRe = regexp.MustCompile("[a-zA-z]+")

// codigos que rodam antes da main (nao tem como retornar erro, apenas panic, fatal)
// - var expressoes
// - func init

func main() {
	mapDemo()

	file, err := os.Open("sherlock.txt")
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	defer file.Close()

	freqs := make(map[string]int) // palavra -> contagem
	s := bufio.NewScanner(file)   // vai escanear o arquivo
	for s.Scan() {

		words := wordRe.FindAllString(s.Text(), -1) // queremos todas as ocorrencias sem limitacoes

		for _, word := range words {
			// log.Println(word)
			freqs[strings.ToLower(word)]++
		}
	}

	if err := s.Err(); err != nil {
		log.Fatalf("SCANNER ERROR: %v\n", err)
	}

	top := topN(freqs, 10)
	fmt.Println(top)
}

func mapDemo() {
	heros := map[string]string{ // hero -> name
		"Superman":     "Clark",
		"Wonder Woman": "Diana",
		"Batman":       "Bruce",
	}

	for k := range heros {
		fmt.Println(k)
	}

	for k, v := range heros {
		fmt.Println(k, "is", v)
	}

	n := heros["Batman"]
	fmt.Println(n)

	n = heros["Aquaman"]
	fmt.Printf("%q\n", n) // string vazia

	// jeito correto - checar se existe
	n, ok := heros["Aquaman"]
	if ok {
		fmt.Println(n)
	} else {
		fmt.Println("aquaman nao achado")
	}

	// deletando
	delete(heros, "Batman")
	fmt.Println(heros)
}

func topN(freq map[string]int, n int) []string {
	// pegar as keys
	seq := maps.Keys(freq) // retorna uma sequencia
	words := slices.Collect(seq)
	sort.Slice(words, func(i, j int) bool {
		wi, wj := words[i], words[j]

		// ordena em ordem reversa
		return freq[wi] > freq[wj]
	})

	n = min(n, len(words))
	return words[:n]
}
