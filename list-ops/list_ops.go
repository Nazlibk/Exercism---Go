package listops

type IntList []int
type unaryFunc func(x int) int
type predFunc func(n int) bool
type binFunc func(x, y int) int

func (list IntList) Length() int {
	count := 0
	for _ = range list {
		count++
	}
	return count
}

func (list1 IntList) Append(list2 IntList) IntList {
	result := make([]int, len(list1)+len(list2))
	copy(result, list1)
	for i, l := range list2 {
		result[i+len(list1)] = l
	}
	return result
}

func (list IntList) Map(unary func(x int) int) IntList {
	result := make([]int, len(list))
	for i, l := range list {
		result[i] = unary(l)
	}
	return result
}

func (list IntList) Filter(predF func(n int) bool) IntList {
	result := make([]int, len(list))
	index := 0
	for _, l := range list {
		if predF(l) {
			result[index] = l
			index++
		}
	}
	return result[:index]
}

func (list IntList) Reverse() IntList {
	result := make([]int, len(list))
	index := len(list) - 1
	for _, l := range list {
		result[index] = l
		index--
	}
	return result
}

func (list IntList) Concat(lists []IntList) IntList {
	length := len(list)
	for _, l := range lists {
		length += len(l)
	}
	var result IntList
	result = make([]int, 0, length)
	result = list
	for _, l := range lists {
		result = result.Append(l)
	}
	return result
}

func (list IntList) Foldl(binF func(x, y int) int, init int) int {
	result := init
	for _, l := range list {
		result = binF(result, l)
	}
	return result
}

func (list IntList) Foldr(binF func(x, y int) int, init int) int {
	result := init
	for i := len(list) - 1; i >= 0; i-- {
		result = binF(list[i], result)
	}
	return result
}
