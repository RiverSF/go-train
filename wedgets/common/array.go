package common

import "strconv"

func IsContain(items []int, item int) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func IsContainStr(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func StrArrToIntArr(strArr []string) []int {
	var intArr []int
	for _, item := range strArr {
		if i, err := strconv.Atoi(item); err == nil {
			intArr = append(intArr, i)
		}
	}
	return intArr
}

// 字符数组：交集
func IntersectionStrSet(a []string, b []string) (inter []string) {
	m := make(map[string]string)
	nn := make([]string, 0)
	for _, v := range a {
		m[v] = v
	}
	for _, v := range b {
		times, _ := m[v]
		if len(times) > 0 {
			nn = append(nn, v)
		}
	}
	return nn
}

// 字符数组：是否有交集
func IsIntersectionStr(slice1 []string, slice2 []string) bool {
	m := make(map[string]string)
	for _, v := range slice1 {
		m[v] = v
	}
	for _, v := range slice2 {
		_, ok := m[v]
		if ok {
			return true
		}
	}
	return false
}

// 字符数组：并集
func SliceUnionStrSet(slice1 []string, slice2 []string) []string {
	result := make([]string, 0)
	flagMap := make(map[string]bool, 0)
	slice1 = append(slice1, slice2...)
	for _, v := range slice1 {
		if _, ok := flagMap[v]; ok {
			continue
		}
		flagMap[v] = true
		result = append(result, v)
	}
	return result
}

// 字符数组：差集
func SupplementaryStrSet(slice1, slice2 []string) []string {
	m := make(map[string]string)
	for _, v := range slice1 {
		m[v] = v
	}
	for _, v := range slice2 {
		if m[v] != "" {
			delete(m, v)
		}
	}
	var str []string
	for _, s2 := range m {
		str = append(str, s2)
	}
	return str
}

// 整形数组：差集
func SupplementaryIntSet(slice1, slice2 []int) []int {
	m := make(map[int]int)
	for _, v := range slice1 {
		m[v] = v
	}
	for _, v := range slice2 {
		if _, ok := m[v]; ok {
			delete(m, v)
		}
	}

	var n []int
	for _, v := range m {
		n = append(n, v)
	}
	return n
}

// 整形数组：是否有交集
func IsIntersectionInt(slice1 []int, slice2 []int) bool {
	m := make(map[int]int)
	for _, v := range slice1 {
		m[v] = v
	}
	for _, v := range slice2 {
		_, ok := m[v]
		if ok {
			return true
		}
	}
	return false
}

// 整形数组：并集
func SliceUnionIntSet(slice1, slice2 []int) []int {
	result := make([]int, 0)
	flagMap := make(map[int]bool, 0)
	slice1 = append(slice1, slice2...)
	for _, v := range slice1 {
		if _, ok := flagMap[v]; ok {
			continue
		}
		flagMap[v] = true
		result = append(result, v)
	}
	return result
}

func IntArrToInt32Arr(intArr []int) []int32 {
	var int32Arr []int32
	for _, i := range intArr {
		int32Arr = append(int32Arr, int32(i))
	}
	return int32Arr
}

func Int32ArrToIntArr(int32Arr []int32) []int {
	var intArr []int
	for _, i := range int32Arr {
		intArr = append(intArr, int(i))
	}
	return intArr
}

func Int32ArrToStringArr(intArr []int32) []string {
	var stringArr []string
	for _, i := range intArr {
		stringArr = append(stringArr, strconv.Itoa(int(i)))
	}
	return stringArr
}
