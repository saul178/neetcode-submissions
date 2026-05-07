// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

type Solution struct {

}

func NewSolution() *Solution {
    return &Solution{}
}

func QuickSort(pairs []Pair) []Pair {
    quickSortHelper(pairs, 0, len(pairs)-1)
    return pairs
}
func quickSortHelper(pairs []Pair, start, end int) {
	if end-start+1 <= 1 {
		return
	}

	pivot := pairs[end]
	leftPtr := start

	for i := start; i < end; i++ {
		if pairs[i].Key < pivot.Key {
			pairs[i], pairs[leftPtr] = pairs[leftPtr], pairs[i]
			leftPtr++
		}
	}
	pairs[end] = pairs[leftPtr]
	pairs[leftPtr] = pivot

	quickSortHelper(pairs, start, leftPtr-1)
	quickSortHelper(pairs, leftPtr+1, end)
}
