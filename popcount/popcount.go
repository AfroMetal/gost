package popcount

// pc[i] is population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

//PopCount returns population count of x (count of bits set to 1)
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	result := 0
	var i uint = 0
	for ; i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}

func PopCount3(x uint64) int {
	result := uint64(0)
	for ; x > 0; x >>= 1 {
		result += x & 1
	}
	return int(result)
}

func PopCount4(x uint64) int {
	result := uint64(0)
	for ; x > 0; x &= x - 1 {
		result++
	}
	return int(result)
}
