package qsort

func quickSort(values []int, left, right int) {
	key := values[left]
	p := left
	i, j := left, right

	for i < j {
		for j > p && values[j] >= key {
			j--
		}
		if j > p {
			values[p] = values[j]
			p = j
		}

		for i < p && values[i] <= key {
			i++
		}
		if i < p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = key

	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}
