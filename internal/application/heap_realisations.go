package application

// Dijkstra ------------------------------------------------------
func (pq PQDijkstra) Len() int { return len(pq) }

func (pq PQDijkstra) Less(i, j int) bool { return pq[i].dist < pq[j].dist }

func (pq PQDijkstra) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PQDijkstra) Push(x any) { *pq = append(*pq, x.(pqItem)) }

func (pq *PQDijkstra) Pop() any {
	old := *pq
	i := old[len(*pq)-1]
	*pq = old[:len(old)-1]
	return i
}

// AStar ------------------------------------------------------

func (pq PQAStar) Len() int { return len(pq) }

func (pq PQAStar) Less(i, j int) bool { return pq[i].dist+pq[i].h < pq[j].dist+pq[j].h }

func (pq PQAStar) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PQAStar) Push(x any) { *pq = append(*pq, x.(astarItem)) }

func (pq *PQAStar) Pop() any {
	old := *pq
	i := old[len(*pq)-1]
	*pq = old[:len(old)-1]
	return i
}
