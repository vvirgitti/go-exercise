package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestPersonHasNameAgeAndFriends(*testing.T) {
	Bob := person{name: "Bob", age: 20, friends: []string{"Raf"}}
	fmt.Print(Bob)
	//Output: {Bob, 20, [Raf]}
}

func TestPersonCanBeMultiple(t *testing.T) {
people := []person{
	{name: "Rob", age: 30, friends: []string{"Pat"}},
	{name: "Tom", age: 31, friends: []string{"Mat"}},
	{name: "Jon", age: 32, friends: []string{"Kat"}},
}

	fmt.Print(people)
	//Output: [{Rob, 30, [Pat]}, {Ton, 31, [Mat]}, {Jon, 31, [Kat]}]
}


func TestPeopleCanMakeFriends(t *testing.T) {
	personA := person{name: "Rod", age: 50}
	personB := person{name: "Tam", age: 45}

	makeFriends(personA, personB)

	fmt.Print("TEST", personA)
	fmt.Print("TEST", personB)

	assert.Equal(t, personA.friends, []string{"Tam"})
	assert.Equal(t, personB.friends, []string{"Rod"})

}

func TestPersonListFriends(t *testing.T) {
	p := person{name: "Bob", age: 20}
	p.friends = append(p.friends, "John")
	p.friends = append(p.friends, "Mark")
	p.friends = append(p.friends, "Luke")

	list := listFriends(p)

	assert.Equal(t, list, []string{"John", "Mark", "Luke"})

}