package main

type person struct {
	name string
	age int
	friends []string
}

func main() {


}


func makeFriends(p *person, p2 *person) {

	p.friends = append(p.friends, p2.name)
	p2.friends = append(p2.friends, p.name)

}

func listFriends(p person) []string {

	return p.friends

}
