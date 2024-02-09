package lc190

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = res<<1 + num&1
		num >>= 1
	}
	// bits.Reverse32(num)
	
	return res
}

// func reverseBits(n uint32) uint32 {
//     n = ((n & 0xFFFF0000) >> 16) | ((n & 0x0000FFFF) << 16)
//     n = ((n & 0xFF00FF00) >> 8) | ((n & 0x00FF00FF) << 8)
//     n = ((n & 0xF0F0F0F0) >> 4) | ((n & 0x0F0F0F0F) << 4)
//     n = ((n & 0xCCCCCCCC) >> 2) | ((n & 0x33333333) << 2)
//     n = ((n & 0xAAAAAAAA) >> 1) | ((n & 0x55555555) << 1)
//     return n
// }
