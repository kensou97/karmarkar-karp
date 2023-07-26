package main

import (
	"fmt"
	"math/rand"
	"testing"
)

type IntMember int

func (i IntMember) Weight() int {
	return int(i)
}

func TestKarmarkarKarp(t *testing.T) {
	weights := []int{5, 5, 5, 4, 4, 3, 3, 1}
	var members []Member
	for _, weight := range weights {
		members = append(members, IntMember(weight))
	}
	p := KarmarkarKarp(members, 3)
	fmt.Println(p)
}

func TestGenerate(t *testing.T) {
	var is []int
	for i := 0; i < 30; i++ {
		is = append(is, rand.Intn(50))
	}
	fmt.Println(is)
	for _, i := range is {
		fmt.Printf("%d,", i)
	}
	fmt.Println()
}
