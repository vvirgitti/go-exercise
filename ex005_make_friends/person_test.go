package main

import (
	"testing"
	"fmt"

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

func TestPeopleCanMakeFriends(*testing.T) {
	personA := person{name: "Rod", age: 50}
	personB := person{name: "Tam", age: 45}

	if personA.makeFriend(personB) {
		assert.Equal(t, personA.friends, ["Tam"], "they should be equal")
	}

}