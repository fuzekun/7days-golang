package leetcode

func quickSort2(arr []int) {
	sort2(arr, 0, len(arr)-1)
}

/*
*
[l, r]
*/
func sort2(arr []int, l int, r int) {
	if l >= r {
		return
	}
	pivoke := arr[l]
	i := l
	j := r
	for i < j {
		for i < j && arr[j] >= pivoke {
			j--
		}
		arr[i] = arr[j]
		for i < j && arr[i] <= pivoke {
			i++
		}
		arr[j] = arr[i]
	}
	arr[i] = pivoke
	sort2(arr, l, i-1)
	sort2(arr, i+1, r)
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)-1] // 选择最后一个元素作为基准

	// 左指针，指向小于的，右指针指向大于的
	left, right := 0, 0

	for right < len(arr)-1 {
		if arr[right] < pivot {
			// 如果右边的大于等于左边的，直接right++。小于的情况下，就会继续left++
			arr[left], arr[right] = arr[right], arr[left]
			left++
		}
		right++
	}

	// 将基准元素放到正确的位置
	arr[left], arr[len(arr)-1] = arr[len(arr)-1], arr[left]

	// 递归排序左右两部分
	quickSort(arr[:left])   // 左半部分: arr[0] 到 arr[left-1]
	quickSort(arr[left+1:]) // 右半部分: arr[left+1] 到 arr[len(arr)-1]

	return arr
}
