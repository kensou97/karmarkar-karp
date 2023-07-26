package main

import (
	"sort"
	"testing"
)

type IntMember int

func (i IntMember) Weight() int {
	return int(i)
}

func TestKarmarkarKarp(t *testing.T) {
	cases := []struct {
		weights  []int
		k        int
		expected [][]int
	}{
		{
			weights:  []int{8, 7, 6, 5, 4},
			k:        2,
			expected: [][]int{{4, 7, 5}, {8, 6}},
		},
		{
			weights:  []int{8, 7, 6, 5, 4},
			k:        3,
			expected: [][]int{{4, 7}, {5, 6}, {8}},
		},
		{
			weights:  []int{5, 5, 5, 4, 4, 3, 3, 1},
			k:        3,
			expected: [][]int{{5, 3, 3}, {5, 1, 4}, {4, 5}},
		},
	}
	for _, c := range cases {
		var members []Member
		for _, weight := range c.weights {
			members = append(members, IntMember(weight))
		}
		p := KarmarkarKarp(members, c.k)
		if !equals(p.subsets, c.expected) {
			t.Errorf("subsets %v not equals to %v", p.subsets, c.expected)
		}
	}
}

func equals(sets []Subset, expected [][]int) bool {
	if len(sets) != len(expected) {
		return false
	}
	for i, set := range sets {
		if len(set.members) != len(expected[i]) {
			return false
		}
		sort.SliceStable(set.members, func(i, j int) bool {
			return set.members[i].Weight() < set.members[j].Weight()
		})
		sort.SliceStable(expected[i], func(p, q int) bool {
			return expected[i][p] < expected[i][q]
		})
		for j, member := range set.members {
			if member.Weight() != expected[i][j] {
				return false
			}
		}
	}
	return true
}
