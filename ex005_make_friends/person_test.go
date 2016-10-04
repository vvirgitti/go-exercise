package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestPersonHasNameAgeAndFriends(*testing.T) {
	p := person{name: "Bob", age: 20}
	p.friends = append(p.friends, "Pat")
	//Output: {Bob, 20, [Pat]}
}

func TestPersonCanBeMultiple(t *testing.T) {
people := []person{
	{name: "Rob", age: 30},
	{name: "Tom", age: 31},
	{name: "Jon", age: 32},
}

	people[0].friends = append(people[0].friends, "Pat")
	people[1].friends = append(people[1].friends, "Mat")
	people[2].friends = append(people[2].friends, "Kat")

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
	//assert.Equal(t, personB.friends, "Rod")

}

func TestPersonListFriends(t *testing.T) {
	p := person{name: "Bob", age: 20}
	p.friends = append(p.friends, "John")
	p.friends = append(p.friends, "Mark")
	p.friends = append(p.friends, "Luke")

	list := listFriends(p)

	assert.Equal(t, list, []string{"John", "Mark", "Luke"})

}