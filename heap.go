package main

type PartitionHeap []Partition

func (pq PartitionHeap) Len() int { return len(pq) }

func (pq PartitionHeap) Less(i, j int) bool {
	return pq[i].Difference() > pq[j].Difference()
}

func (pq PartitionHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PartitionHeap) Push(x any) {
	*pq = append(*pq, x.(Partition))
}

func (pq *PartitionHeap) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
