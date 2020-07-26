package src

func Solution(n, k int, array []int) int {
	sortedArray := MergeSort(array)
	resultMap := make(map[int]int)

	for i := 0; i < len(sortedArray); i++ {
		_, ok := resultMap[int(sortedArray[i]/k)]
		if sortedArray[i]%k != 0 {
			resultMap[sortedArray[i]] = 1
		} else if !ok {
			resultMap[sortedArray[i]] = 1
		}
	}
	return len(resultMap)
}

func merge(arr1, arr2 []int) []int {
	size, i, j := len(arr1)+len(arr2), 0, 0
	result := make([]int, size)
	for k := 0; k < size; k++ {
		switch true {
		case i == len(arr1):
			result[k] = arr2[j]
			j += 1
		case j == len(arr2):
			result[k] = arr1[i]
			i += 1
		case arr1[i] > arr2[j]:
			result[k] = arr2[j]
			j += 1
		case arr1[i] < arr2[j]:
			result[k] = arr1[i]
			i += 1
		case arr1[i] == arr2[j]:
			result[k] = arr2[j]
			j += 1
		}
	}
	return result
}
func MergeSort(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}
	middle := int(len(numbers) / 2)
	return merge(MergeSort(numbers[middle:]), MergeSort(numbers[:middle]))
}
