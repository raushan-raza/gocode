package Bitpackage

func Isset(x, k int) bool {
	res := 1
	for i := 0; i <= k-1; i++ {
		res = res * 2
	}
	if x&res != 0 {
		return true
	} else {
		return false
	}
}

func CountBits(n int) int {
	res := 0
	for n > 0 {
		if n%2 != 0 {
			res = res + 1
		}
		n = n / 2
	}
	return res
}

func IsPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	for n != 1 {
		if n%2 != 0 {
			return false
		}
		n = n / 2
	}
	return true
}

func FindOddElement(slice []int) int {
	for i := 0; i < len(slice); i++ {
		count := 1
		for j := i + 1; j < len(slice); j++ {
			if slice[i] == slice[j] {
				count++
			}
		}
		if count&1 == 1 {
			return slice[i]
		}
	}
	return 0
}
