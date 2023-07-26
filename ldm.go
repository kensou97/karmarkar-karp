package ldm

import (
	"container/heap"
	"sort"
)

type Member interface {
	Weight() int
}

type Subset struct {
	members []Member
	sum     int64
}

func (s *Subset) Members() []Member {
	return s.members
}

func (s *Subset) Sum() int64 {
	return s.sum
}

func (s *Subset) Merge(other Subset) {
	s.members = append(s.members, other.members...)
	s.sum += other.sum
}

type Partition struct {
	subsets []Subset
}

func (p *Partition) Subsets() []Subset {
	return p.subsets
}

func (p *Partition) Difference() int64 {
	subsets := p.subsets
	return subsets[0].sum - subsets[len(subsets)-1].sum
}

func (p *Partition) Merge(other Partition) {
	var (
		si = p.subsets
		sj = other.subsets
		i  = 0
		j  = len(sj) - 1
	)
	for i < len(si) && j >= 0 {
		si[i].Merge(sj[j])
		i++
		j--
	}
	sort.SliceStable(p.subsets, func(i, j int) bool {
		return p.subsets[i].sum > p.subsets[j].sum
	})
}

func PartitionWithMember(member Member, k int) Partition {
	var subsets []Subset
	subsets = append(subsets, Subset{members: []Member{member}, sum: int64(member.Weight())})
	for i := 1; i < k; i++ {
		subsets = append(subsets, Subset{})
	}
	return Partition{subsets: subsets}
}

func KarmarkarKarp(members []Member, k int) Partition {
	h := &PartitionHeap{}
	for _, member := range members {
		*h = append(*h, PartitionWithMember(member, k))
	}
	heap.Init(h)
	for h.Len() > 1 {
		p1 := heap.Pop(h).(Partition)
		p2 := heap.Pop(h).(Partition)
		p2.Merge(p1)
		heap.Push(h, p2)
	}
	return heap.Pop(h).(Partition)
}
