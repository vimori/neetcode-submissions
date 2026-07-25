func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := l + (r-l)/2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target{
			r = m
			}else {
				return m
			}
		}

		if nums[l] == target{
			return l
		}
	return -1
	}
