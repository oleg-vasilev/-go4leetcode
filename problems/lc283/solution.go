package lc283

func moveZeroes(nums []int) {
	moveZeroesV2(nums)
}

// moveZeroesV1 is simplest straightforward solution with O(n^2) complexity
func moveZeroesV1(nums []int) {
a:
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			continue
		}
		// zero found at i pos -
		// find next non-zero value and switch
		for iz := i + 1; iz < len(nums); iz++ {
			if nums[iz] != 0 {
				nums[i], nums[iz] = nums[iz], nums[i]
				continue a
			}
		}
	}
}

// moveZeroesV2 is smart solution with O(n) complexity based on fact that new
// position for non-zero item can be calculated using number of zeros before
func moveZeroesV2(nums []int) {
	var zerosCount int
	for i, v := range nums {
		if v == 0 {
			zerosCount++
		} else {
			if zerosCount > 0 {
				nums[i-zerosCount], nums[i] = v, 0
			}
		}
	}
}
