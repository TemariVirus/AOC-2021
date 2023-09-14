package main

type BitArray struct {
	Length int
	words  []uint64
}

func makeBitArray(length, capacity int) BitArray {
	length = max(0, length)
	capacity = max(capacity, length)
	capacity = (capacity + 63) / 64
	return BitArray{length, make([]uint64, capacity)}
}

func (b *BitArray) get(index int) bool {
	return (b.words[index/64]>>(index%64))&1 == 1
}

func (b *BitArray) getRange(start, end int) uint64 {
	start_word := start / 64
	end_word := (end - 1) / 64
	start_bit := start % 64
	end_bit := (end-1)%64 + 1
	end_bit = 64 - end_bit

	var word uint64
	if start_word == end_word {
		word = b.words[start_word] << end_bit >> (end_bit + start_bit)
	} else {
		word = (b.words[start_word] >> start_bit) | (b.words[end_word] << end_bit >> end_bit << (64 - start_bit))
	}

	reversed := uint64(0)
	for i := 0; i < end-start; i++ {
		reversed <<= 1
		reversed |= word & 1
		word >>= 1
	}
	return reversed
}

func (b *BitArray) set(index int) {
	b.words[index/64] |= uint64(1) << (index % 64)
}

func (b *BitArray) unset(index int) {
	b.words[index/64] &= ^(uint64(1) << (index % 64))
}

func (b *BitArray) append(set bool) {
	index := b.Length
	if b.Length/64 >= len(b.words) {
		b.words = append(b.words, make([]uint64, max(1, len(b.words)))...)
	}
	if set {
		b.set(index)
	} else {
		b.unset(index)
	}
	b.Length++
}

func (b *BitArray) toSlice() []bool {
	result := make([]bool, b.Length)
	for i := 0; i < b.Length; i++ {
		result[i] = b.get(i)
	}
	return result
}

func (b *BitArray) setCount() int {
	return aggregate(b.words, 0, func(agg int, value uint64, index int) int {
		return agg + popcount(value)
	})
}
