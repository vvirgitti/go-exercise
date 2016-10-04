package main

import "fmt"

func main() {

}

type person struct {
	name string
	age int
	friends []string
}

func (p person) makeFriend(p2 person) person {
	fmt.Println("PERSON WITHIN THE FUNCTION", p)
	fmt.Println("PERSON WITHIN THE FUNCTION 2", p2)

	p.friends = append(p.friends, p2.name)
	//p2.friends = append(p2.friends, p.name)
	fmt.Println("PERSON END OF FUNCTION", p)
	//fmt.Println("PERSON END OF FUNCTION", p2)
	return p
}
