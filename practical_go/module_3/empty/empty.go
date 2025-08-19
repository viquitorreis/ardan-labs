package main

import "fmt"

func main() {
	var a any // interface{}

	a = 7
	fmt.Println("a:", a)

	a = "Hi"
	fmt.Println("a:", a)
	/* NAO USE ANY, a não ser que (geralmente):
	- Serialização
	- Printar algo
	*/

	s := a.(string) // type assertion
	fmt.Println("s:", s)

	// i := a.(int) // vai dar panic
	i, ok := a.(int)
	if ok {
		fmt.Println("i:", i)
	} else {
		fmt.Printf("not an int (%T)\n", a)
	}

	// type switch
	switch a.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Printf("other type: (%T)\n", a)
	}
}
