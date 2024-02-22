package lc380

import (
	"math/rand"
	"time"
)

// 380. Insert Delete GetRandom O(1)

type RandomizedSet struct {
	index map[int]int
	// index  map[int]struct{}
	values []int
	rnd    *rand.Rand
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		// index:  make(map[int]struct{}),
		index:  make(map[int]int),
		values: make([]int, 0),
		rnd:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, exists := this.index[val]
	if !exists {
		// this.index[val] = struct{}{}

		this.values = append(this.values, val)
		this.index[val] = len(this.values) - 1
	}
	return !exists
}

func (this *RandomizedSet) Remove(val int) bool {
	i, exists := this.index[val]
	if exists {

		// idea is to replace deleted item with last item and cut last item from slice

		lastIndex := len(this.values) - 1
		lastElm := this.values[lastIndex]
		this.values[i] = lastElm
		this.index[lastElm] = i

		this.values = this.values[:lastIndex]
		delete(this.index, val)
	}
	return exists
}

func (this *RandomizedSet) GetRandom() int {
	// iteration over map for random item will take ~O(n), but we need ~O(1) ....
	randomIndex := this.rnd.Intn(len(this.index))
	// for k, _ := range this.index {
	// 	if randomIndex == 0 {
	// 		return k
	// 	}
	// 	randomIndex--
	// }

	return this.values[randomIndex]
}
