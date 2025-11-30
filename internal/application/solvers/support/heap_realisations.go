package support

// Dijkstra ------------------------------------------------------
func (pq PQDijkstra) Len() int { return len(pq) }

func (pq PQDijkstra) Less(i, j int) bool { return pq[i].Dist < pq[j].Dist }

func (pq PQDijkstra) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PQDijkstra) Push(x any) { *pq = append(*pq, x.(PqItem)) }

func (pq *PQDijkstra) Pop() any {
	old := *pq
	i := old[len(*pq)-1]
	*pq = old[:len(old)-1]
	return i
}

// AStar ------------------------------------------------------

func (pq PQAStar) Len() int { return len(pq) }

func (pq PQAStar) Less(i, j int) bool { return pq[i].Dist+pq[i].H < pq[j].Dist+pq[j].H }

func (pq PQAStar) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PQAStar) Push(x any) { *pq = append(*pq, x.(AstarItem)) }

func (pq *PQAStar) Pop() any {
	old := *pq
	i := old[len(*pq)-1]
	*pq = old[:len(old)-1]
	return i
}
