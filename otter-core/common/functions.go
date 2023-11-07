package common

func CalDAGTopologicalSequence(dagGraph map[string][]string) []string {
	dict := make(map[string]int32, 0)
	for _, neighbors := range dagGraph {
		for _, neighbor := range neighbors {
			dict[neighbor] += 1
		}
	}
	queue := make([]string, 0)
	result := make([]string, 0)
	for node, _ := range dagGraph {
		if _, ok := dict[node]; !ok {
			queue = append(queue, node)
		}
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = append(queue[:0], queue[1:]...)
		result = append(result, node)

		neighbors := dagGraph[node]
		for _, neighbor := range neighbors {
			dict[neighbor] -= 1
			if dict[neighbor] <= 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	if len(result) == len(dagGraph) {
		return result
	}
	return nil
}
