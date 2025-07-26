package main

import (
	"encoding/binary"
)

func murmur3_32(data []byte, seed uint32) uint32 {
	const (
		c1 = 0xcc9e2d51
		c2 = 0x1b873593
	)

	length := len(data)
	h1 := seed

	nblocks := length / 4
	for i := 0; i < nblocks; i++ {
		k1 := binary.LittleEndian.Uint32(data[i*4 : i*4+4])
		k1 *= c1
		k1 = (k1 << 15) | (k1 >> (32 - 15))
		k1 *= c2

		h1 ^= k1
		h1 = (h1 << 13) | (h1 >> (32 - 13))
		h1 = h1*5 + 0xe6546b64
	}

	var k1 uint32
	tail := data[nblocks*4:]

	switch len(tail) {
	case 3:
		k1 ^= uint32(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(tail[0])
		k1 *= c1
		k1 = (k1 << 15) | (k1 >> (32 - 15))
		k1 *= c2
		h1 ^= k1
	}

	h1 ^= uint32(length)

	// Finalization mix (fmix)
	h1 ^= h1 >> 16
	h1 *= 0x85ebca6b
	h1 ^= h1 >> 13
	h1 *= 0xc2b2ae35
	h1 ^= h1 >> 16

	return h1
}
